package producer

import (
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

	"github.com/bygui86/go-traces/kafka-producer/commons"
	"github.com/bygui86/go-traces/kafka-producer/logging"
	"github.com/bygui86/go-traces/kafka-producer/monitoring"
	"github.com/bygui86/go-traces/kafka-producer/tracing"
)

// Produce messages to topic (asynchronously)
func (p *KafkaProducer) startProducer() {
	p.ticker = time.NewTicker(1 * time.Second)
	topicPartition := kafka.TopicPartition{Topic: &p.config.kafkaTopic, Partition: kafka.PartitionAny}
	counter := 0

	for {
		select {
		case <-p.ticker.C:
			span := opentracing.StartSpan(p.name)

			span.SetTag("app", commons.ServiceName)

			msg := p.messages[counter]

			carrier := tracing.KafkaHeadersCarrier([]kafka.Header{
				{"example", []byte("example-value")},
				{"app", []byte(commons.ServiceName)},
			})
			traceErr := tracing.Inject(span, &carrier)
			if traceErr != nil {
				logging.SugaredLog.Errorf("Producer failed to inject tracing span: %s", traceErr.Error())
				monitoring.IncreaseFailTracingInjCounter(commons.ServiceName, p.config.kafkaTopic)
			}

			kafkaMsg := &kafka.Message{
				TopicPartition: topicPartition,
				Value:          []byte(msg),
				Headers:        carrier,
			}

			err := p.producer.Produce(kafkaMsg, nil)
			if err != nil {
				logging.SugaredLog.Errorf("Producer failed to publish message %s: %s", msg, err.Error())
				monitoring.IncreaseFailPublishedMsgCounter(commons.ServiceName, p.config.kafkaTopic)
				continue
			}

			if counter == len(p.messages)-1 {
				counter = 0
			} else {
				counter++
			}

			span.Finish()

			monitoring.IncreaseOpsCounter(commons.ServiceName)

		case <-p.stop:
			return
		}
	}
}

// Delivery report handler for produced messages
func (p *KafkaProducer) startEventListener() {
	logging.Log.Info("Start kafka event listener")

	// WARN: headers are not yet supported by kafka.Producer.Events channel
	for e := range p.producer.Events() {

		switch event := e.(type) {

		case *kafka.Message:
			if event.TopicPartition.Error != nil { // FAIL
				topicInfo := fmt.Sprintf("name[%s], partition[%d], offset[%d]",
					*event.TopicPartition.Topic, event.TopicPartition.Partition, event.TopicPartition.Offset)
				msgInfo := fmt.Sprintf("key[%s], value[%s], timestamp[%v]",
					string(event.Key), string(event.Value), event.Timestamp)
				headersInfo := fmt.Sprintf("%+v", event.Headers)
				logging.SugaredLog.Errorf("Message delivery FAILED: topic[%s], msg[%s], headers[%s]",
					topicInfo, msgInfo, headersInfo)

				monitoring.IncreaseFailPublishedMsgCounter(commons.ServiceName, *event.TopicPartition.Topic)

			} else { // SUCCESS
				topicInfo := fmt.Sprintf("name[%s], partition[%d], offset[%d]",
					*event.TopicPartition.Topic, event.TopicPartition.Partition, event.TopicPartition.Offset)
				msgInfo := fmt.Sprintf("key[%s], value[%s], timestamp[%v]",
					string(event.Key), string(event.Value), event.Timestamp)
				headersInfo := fmt.Sprintf("%+v", event.Headers)
				logging.SugaredLog.Infof("Message delivery: topic[%s], msg[%s], headers[%s]",
					topicInfo, msgInfo, headersInfo)

				monitoring.IncreaseSuccPublishedMsgCounter(commons.ServiceName, *event.TopicPartition.Topic)
			}

		default:
			logging.SugaredLog.Debugf("Kafka event not supported: %+v", event)
		}
	}
}

package consumer

import (
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

	"github.com/bygui86/go-traces/kafka-consumer/commons"
	"github.com/bygui86/go-traces/kafka-consumer/logging"
	"github.com/bygui86/go-traces/kafka-consumer/monitoring"
	"github.com/bygui86/go-traces/kafka-consumer/tracing"
)

func (c *KafkaConsumer) subscribeToTopics() error {
	subErr := c.consumer.SubscribeTopics(c.config.kafkaTopics, nil)
	if subErr != nil {
		return subErr
	}
	return nil
}

func (c *KafkaConsumer) startConsumer() {
	for {
		msg, err := c.consumer.ReadMessage(-1)

		monitoring.IncreaseOpsCounter(commons.ServiceName)

		if err == nil { // SUCCESS
			carrier := tracing.KafkaHeadersCarrier(msg.Headers)
			spanCtx, extErr := tracing.Extract(&carrier)
			if extErr != nil {
				logging.SugaredLog.Errorf("Error extracting span context from message: %s", extErr.Error())
			}

			var span opentracing.Span
			if spanCtx != nil {
				span = opentracing.StartSpan(c.name, ext.RPCServerOption(spanCtx))
			}

			topicInfo, msgInfo, headersInfo := c.getMessageInfo(msg)
			logging.SugaredLog.Infof("Message received: topic[%s], msg[%s], headers[%s]",
				topicInfo, msgInfo, headersInfo)

			if span != nil {
				span.SetTag("app", commons.ServiceName)
				span.Finish()
			}

			monitoring.IncreaseSuccConsumedMsgCounter(commons.ServiceName, *msg.TopicPartition.Topic)

		} else { // FAIL
			// INFO: The client will automatically try to recover from all errors.
			if msg != nil {
				topicInfo, msgInfo, headersInfo := c.getMessageInfo(msg)
				logging.SugaredLog.Errorf("Consumer error on message: topic[%s], msg[%s], headers[%s], error[%s]",
					topicInfo, msgInfo, headersInfo, err.Error())
				monitoring.IncreaseFailConsumedMsgCounter(commons.ServiceName, *msg.TopicPartition.Topic)
			} else {
				logging.SugaredLog.Errorf("Consumer error: %s", err.Error())
				monitoring.IncreaseConsumerErrCounter(commons.ServiceName)
			}
		}
	}
}

func (c *KafkaConsumer) getMessageInfo(msg *kafka.Message) (string, string, string) {
	topicInfo := fmt.Sprintf("name[%s], partition[%d], offset[%d]",
		*msg.TopicPartition.Topic, msg.TopicPartition.Partition, msg.TopicPartition.Offset)
	msgInfo := fmt.Sprintf("key[%s], value[%s], timestamp[%v]",
		string(msg.Key), string(msg.Value), msg.Timestamp)
	headersInfo := fmt.Sprintf("%+v", msg.Headers)
	return topicInfo, msgInfo, headersInfo
}

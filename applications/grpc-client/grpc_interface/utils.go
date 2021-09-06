package grpc_interface

import (
	"context"
	"strings"
	"time"

	"github.com/bygui86/go-traces/grpc-client/commons"
	"github.com/bygui86/go-traces/grpc-client/logging"
	"github.com/bygui86/go-traces/grpc-client/monitoring"
)

func (c *Client) startGreeting() {
	logging.SugaredLog.Infof("Start greeting [ %s ] every %s sec",
		strings.Join(c.config.greetingNames, commons.ListSeparator), c.config.greetingInterval.String())
	c.ticker = time.NewTicker(c.config.greetingInterval)

	index := 0
	for {
		select {
		case <-c.ticker.C:
			if index > len(c.config.greetingNames)-1 {
				index = 0
			}

			c.greet(index)

			index++
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Client) greet(index int) {
	// WARN: the connection context is one-shot, it must be refreshed before every request
	ctx, cancel := context.WithTimeout(context.Background(), c.config.grpcConnectionTimeout)
	defer cancel()

	name := c.config.greetingNames[index]

	response, err := c.helloServiceClient.SayHello(ctx, &HelloRequest{Name: name})
	if err != nil {
		logging.SugaredLog.Errorf("Greeting to %s failed: %v", name, err.Error())
		return
	}

	logging.SugaredLog.Infof(response.Greeting)

	monitoring.IncreaseOpsCounter(commons.ServiceName)
	monitoring.IncreaseGreetingsCounter(commons.ServiceName, name)
}

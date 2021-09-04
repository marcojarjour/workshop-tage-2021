package grpc_interface

import (
	"context"
	"time"

	"github.com/bygui86/go-traces/grpc-client/commons"
	"github.com/bygui86/go-traces/grpc-client/logging"
	"github.com/bygui86/go-traces/grpc-client/monitoring"
)

func (c *Client) startGreeting() {
	logging.SugaredLog.Infof("Start greeting %s every %s sec",
		c.config.greetingName, c.config.greetingInterval.String())
	c.ticker = time.NewTicker(c.config.greetingInterval)

	for {
		select {
		case <-c.ticker.C:
			// WARN: the connection context is one-shot, it must be refreshed before every request
			ctx, cancel := context.WithTimeout(context.Background(), c.config.grpcConnectionTimeout)
			defer cancel()

			response, err := c.helloServiceClient.SayHello(ctx, &HelloRequest{Name: c.config.greetingName})
			if err != nil {
				logging.SugaredLog.Errorf("Greeting %s failed: %v", c.config.greetingName, err.Error())
				continue
			}

			logging.SugaredLog.Infof(response.Greeting)

			monitoring.IncreaseOpsCounter(commons.ServiceName)
			monitoring.IncreaseGreetingsCounter(commons.ServiceName, c.config.greetingName)

		case <-c.ctx.Done():
			return
		}
	}
}

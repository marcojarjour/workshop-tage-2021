package grpc_interface

import (
	"time"

	"github.com/bygui86/go-traces/grpc-client/logging"
	"github.com/bygui86/go-traces/grpc-client/utils"
)

const (
	serverAddressEnvVar     = "GRPC_SERVER_ADDRESS"
	connectionTimeoutEnvVar = "GRPC_CONNECTION_TIMEOUT" // in seconds
	greetingNamesEnvVar     = "GRPC_GREETING_NAMES"     // comma-separated list
	greetingIntervalEnvVar  = "GRPC_GREETING_INTERVAL"  // in seconds

	serverAddressDefault     = "localhost:50051"
	connectionTimeoutDefault = 2
	greetingIntervalDefault  = 1
)

var (
	greetingNamesDefault = []string{"anonymous"}
)

func loadConfig() *config {
	logging.Log.Info("Load gRPC configurations")

	return &config{
		grpcServerAddress:     utils.GetStringEnv(serverAddressEnvVar, serverAddressDefault),
		grpcConnectionTimeout: time.Duration(utils.GetIntEnv(connectionTimeoutEnvVar, connectionTimeoutDefault)) * time.Second,
		greetingNames:         utils.GetStringArrayEnv(greetingNamesEnvVar, greetingNamesDefault),
		greetingInterval:      time.Duration(utils.GetIntEnv(greetingIntervalEnvVar, greetingIntervalDefault)) * time.Second,
	}
}

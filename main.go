package main

import (
	"context"
	"time"

	hz "github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/logger"
)

const clusterName = "PUT-YOUR-CLUSTER-NAME-HERE"
const cloudToken = "PUT-YOUR-HAZELCAST-CLOUD-TOKEN-HERE"

func main() {
	// We create two loggers below, one for the application and one for the client.
	// That's to make line/function logging work OK.
	// Create the zap Logger for the application and make sure all log lines are handled.
	lg := MakeZapLogger(0)
	defer lg.Sync()
	// This logger is to be used for the client.
	zl := MakeZapLogger(2)
	defer zl.Sync()
	lg.Info("The application started running...")
	// Create the default configuration.
	config := hz.Config{}
	// Set the zap Log Adaptor as the custom logger for the client.
	config.Logger.CustomLogger = NewZapLogAdaptor(logger.WeightTrace, zl)
	// Configure the client to work with Hazelcast Cloud.
	// You can delete the configuration lines below to use an on-prem Hazelcast instance.
	cc := &config.Cluster
	cc.Name = clusterName
	cc.Cloud.Enabled = true
	cc.Cloud.Token = cloudToken
	// End of Hazelcast Cloud settings.
	// Crate and start the client.
	client, err := hz.StartNewClientWithConfig(context.TODO(), config)
	if err != nil {
		panic(err)
	}
	// Just wait a bit so some log is displayed.
	time.Sleep(5 * time.Second)
	// Time to shutdown.
	if err := client.Shutdown(context.TODO()); err != nil {
		panic(err)
	}
	lg.Info("That's all folks!")
}

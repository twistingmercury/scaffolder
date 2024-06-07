package main

import (
	"twistingmercury/test/conf"
	"context"

	"github.com/spf13/viper"
	"github.com/twistingmercury/telemetry/logging"

	"twistingmercury/test/server"

	"github.com/twistingmercury/utils"
)

func main() {
	conf.Initialize()

	ctx, cancel := context.WithCancel(context.Background())
	utils.ListenForInterrupt(cancel)

	if err := server.Bootstrap(ctx,
		viper.GetString(conf.ViperServiceNameKey),
		viper.GetString(conf.ViperBuildVersionKey),
		viper.GetString(conf.ViperNamespaceKey),
		viper.GetString(conf.ViperEnviormentKey),
	); err != nil {
		logging.Fatal(err, "failed to bootstrap the server")
	}

	server.Start()
}

package main

import (
	"log"

	"github.com/Be1chenok/messenger/config"
	"github.com/Be1chenok/messenger/internal/delivery/http/handler"
	"github.com/Be1chenok/messenger/internal/delivery/http/server"
	"github.com/Be1chenok/messenger/pkg/logger"
	"github.com/spf13/cobra"
)

const configFlagName = "config"

func main() {
	var conf *config.Config

	rootCmd := &cobra.Command{
		Use:   "server",
		Short: "management server",
		Run: func(cmd *cobra.Command, args []string) {
			confPath, err := cmd.Flags().GetString(configFlagName)
			if err != nil {
				log.Fatalf("get config path flag value: %v", err)
			}

			conf, err = config.New(confPath)
			if err != nil {
				log.Fatalf("init config: %v", err)
			}
		},
	}

	rootCmd.Flags().StringP(configFlagName, "c", "", "config file path")

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("execute root command: %v", err)
	}

	run(conf)
}

func run(conf *config.Config) {
	logger, err := logger.New(conf.Logger.DebugMode, conf.Logger.LogDir)
	if err != nil {
		log.Fatalf("init logger: %v", err)
	}

	handler := handler.New(&conf.Handler, logger)

	srv := server.New(&conf.HTTPServer, logger, server.InitHandler(handler.Init))

	srv.Run()
}

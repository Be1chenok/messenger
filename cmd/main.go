package main

import (
	"log"

	"github.com/Be1chenok/messenger/config"
	"github.com/Be1chenok/messenger/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

const configFlagName = "config"

func main() {
	logger, err := logger.New()
	if err != nil {
		log.Fatal(err)
	}

	var conf *config.Config

	rootCmd := &cobra.Command{
		Use:   "server",
		Short: "management server",
		Run: func(cmd *cobra.Command, args []string) {
			confPath, err := cmd.Flags().GetString(configFlagName)
			if err != nil {
				logger.Fatal().Err(err).Msg("failed to get config path flag value")
			}

			conf, err = config.New(confPath)
			if err != nil {
				logger.Fatal().Err(err).Msg("failed to init config")
			}
		},
	}

	rootCmd.Flags().StringP(configFlagName, "c", "", "config file path")

	if err := rootCmd.Execute(); err != nil {
		logger.Fatal().Err(err).Msg("failed to execute root command")
	}

	run(conf, logger)
}

func run(conf *config.Config, logger *zerolog.Logger) {

}

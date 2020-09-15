package cmd

import (
	"github.com/neuvector/neuvector-nexus-iq/integration"
	"github.com/neuvector/neuvector-nexus-iq/internal/logger"
	"github.com/spf13/viper"
)

var (
	rootConfig integration.RootConfig
	configFile string
)

var v = viper.New()

func initConfig() {
	v.SetConfigType("yaml")

	if configFile != "" {
		// Explicitly set config file path
		v.SetConfigFile(configFile)
	} else {
		// Implicitly Search config in current directory
		v.AddConfigPath(".")
		v.SetConfigName("config")
	}

	// Load configurations from environment variables
	// For example: NV_NX_ADDRESS, NV_NX_PORT
	// Environment variables are bound in arguments.go
	v.SetEnvPrefix("nv_nx")
	v.AutomaticEnv()

	// Read from config file
	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
			logger.Errorf("config file %s", err)
		}
	}

	configFileUsed := viper.ConfigFileUsed()
	if configFileUsed != "" {
		logger.Infof("using config file %s", configFileUsed)
	}

	// Parse config into struct
	// https://github.com/spf13/viper#unmarshaling
	err = v.Unmarshal(&rootConfig)
	if err != nil {
		logger.Fatalf("unable to parse config file: %v", err)
	}
}

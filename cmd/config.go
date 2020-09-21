package cmd

import (
	"bytes"
	"github.com/neuvector/neuvector-nexus-iq/integration"
	"github.com/neuvector/neuvector-nexus-iq/internal/logger"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"strings"
)

var (
	rootConfig integration.RootConfig
	configFile string
)

var v = viper.New()

func initConfig() {
	v.SetConfigType("yaml")

	// Load default configuration
	// Viper needs to know the configuration keys in order to load configuration from environment variables
	b, err := yaml.Marshal(integration.DefaultConfig)
	if err != nil {
		logger.Fatalf("unable to load default config: %v", err)
	}

	defaultConfig := bytes.NewReader(b)
	if err := v.MergeConfig(defaultConfig); err != nil {
		logger.Fatalf("unable to load default config: %v", err)
	}

	// Load configuration from environment variables
	// For example: NV_NX_ADDRESS, NV_NX_PORT, NV_NX_NEUVECTOR_USERNAME, NV_NX_NEUVECTOR_PASSWORD
	v.AutomaticEnv()
	v.SetEnvPrefix("NV_NX")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if configFile != "" {
		// Explicitly set config file path
		v.SetConfigFile(configFile)
	} else {
		// Implicitly search config in current directory
		v.AddConfigPath(".")
		v.SetConfigName("config")
	}

	// Read from config file
	err = v.MergeInConfig()
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
		logger.Fatalf("unable to load config: %v", err)
	}
}

package integration

// RootConfig contains the configuration of the integration
// "mapstructure" annotation is required for Viper to unmarshal from YAML
// "yaml" annotation is required to marshal the (default) config
type RootConfig struct {
	NeuVectorConfig NeuVectorConfig `mapstructure:"neuvector" yaml:"neuvector"`
	NexusIqConfig   NexusIqConfig   `mapstructure:"nexusiq" yaml:"nexusiq"`

	Address string `mapstructure:"address" yaml:"address"`
	Port    int16  `mapstructure:"port" yaml:"port"`
	Verbose bool   `mapstructure:"verbose" yaml:"verbose"`
}

type NeuVectorConfig struct {
	Endpoint string `mapstructure:"endpoint" yaml:"endpoint"`
	Username string `mapstructure:"username" yaml:"username"`
	Password string `mapstructure:"password" yaml:"password"`
	Insecure bool   `mapstructure:"insecure" yaml:"insecure"`
}

type NexusIqConfig struct {
	Endpoint         string `mapstructure:"endpoint" yaml:"endpoint"`
	Username         string `mapstructure:"username" yaml:"username"`
	Password         string `mapstructure:"password" yaml:"password"`
	Insecure         bool   `mapstructure:"insecure" yaml:"insecure"`
	OrganizationId   string `mapstructure:"organization_id" yaml:"organization_id"`
	OrganizationName string `mapstructure:"organization_name" yaml:"organization_name"`
	Source           string `mapstructure:"source" yaml:"source"`
	AppNameLabel     string `mapstructure:"app_name_label" yaml:"app_name_label"`
}

var DefaultConfig = RootConfig{
	Address: "127.0.0.1",
	Port:    5080,

	NeuVectorConfig: NeuVectorConfig{
		Endpoint: "",
		Username: "",
		Password: "",
		Insecure: false,
	},

	NexusIqConfig: NexusIqConfig{
		Endpoint:         "",
		Username:         "",
		Password:         "",
		Insecure:         false,
		OrganizationId:   "",
		OrganizationName: "",
		Source:           "NeuVector",
		AppNameLabel:     "",
	},
}

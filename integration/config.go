package integration

type RootConfig struct {
	NeuVectorConfig NeuVectorConfig `mapstructure:"neuvector"`
	NexusIqConfig   NexusIqConfig   `mapstructure:"nexusiq"`

	Address string `mapstructure:"address"`
	Port    int16  `mapstructure:"port"`
	Verbose bool   `mapstructure:"verbose"`
}

type NeuVectorConfig struct {
	Endpoint string `mapstructure:"endpoint"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Insecure bool   `mapstructure:"insecure"`
}

type NexusIqConfig struct {
	Endpoint         string `mapstructure:"endpoint"`
	Username         string `mapstructure:"username"`
	Password         string `mapstructure:"password"`
	Insecure         bool   `mapstructure:"insecure"`
	OrganizationId   string `mapstructure:"organization_id"`
	OrganizationName string `mapstructure:"organization_name"`
	Source           string `mapstructure:"source"`
	AppNameLabel     string `mapstructure:"app_name_label"`
}

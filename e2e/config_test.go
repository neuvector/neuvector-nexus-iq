// +build e2e

package e2e

import "github.com/neuvector/neuvector-nexus-iq/integration"

func newTestConfig() *integration.RootConfig {
	c := &integration.RootConfig{
		// Loopback interface of the Docker host is reachable from Docker containers on the same host via host.docker.internal
		Address: "127.0.0.1",
		Port:    12080,
		NeuVectorConfig: integration.NeuVectorConfig{
			Endpoint: "https://127.0.0.1:10443",
			Insecure: true,
			Username: "admin",
			Password: "admin",
		},
		NexusIqConfig: integration.NexusIqConfig{
			Endpoint:         "http://127.0.0.1:8070",
			Insecure:         true,
			Username:         "admin",
			Password:         "admin123",
			OrganizationName: "Sandbox Organization",
			Source:           "neuvector",
			AppNameLabel:     "com.sonatype.nexus.iq.applicationName",
		},
	}

	return c
}

func newIntegrationServer() *integration.Server {
	config := newTestConfig()
	return integration.NewServer(config)
}

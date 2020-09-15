package integration

import (
	"fmt"
	"github.com/neuvector/neuvector-nexus-iq/nexusiq"
	"net/url"
)

func NewNexusIqClientFromConfig(c NexusIqConfig) (*nexusiq.Client, error) {
	// Parse endpoint
	nxUrl, err := url.Parse(c.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid Nexus IQ endpoint %s", c.Endpoint)
	}

	// Create client instance
	client := nexusiq.NewClient(&nexusiq.Client{
		Schemes:  []string{nxUrl.Scheme},
		Host:     nxUrl.Host,
		Insecure: c.Insecure,
		Username: c.Username,
		Password: c.Password,
		//Debug:    true,
	})

	return client, nil
}

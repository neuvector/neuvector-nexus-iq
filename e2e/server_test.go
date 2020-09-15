// +build e2e

package e2e

import (
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/neuvector/neuvector-nexus-iq/integration"
	"net/http"
	"testing"
	"time"
)

func TestIntegrationServerStart(t *testing.T) {
	// Configuration
	config := newTestConfig()

	// Start integration server
	integrationServer := newIntegrationServer()

	go func() {
		integrationServer.Start()
	}()

	// Wait for server to be healthy
	err := waitForIntegrationServer(config)
	if err != nil {
		t.Error(err)
	}
}

func waitForIntegrationServer(config *integration.RootConfig) error {
	healthEndpoint := fmt.Sprintf("http://%[1]s:%[2]d/health", config.Address, config.Port)
	healthOperation := func() error {
		_, err := http.Get(healthEndpoint)
		if err != nil {
			return err
		}
		return nil
	}

	err := backoff.Retry(healthOperation, backoff.WithMaxRetries(backoff.NewConstantBackOff(1*time.Second), 10))
	if err != nil {
		return err
	}

	return nil
}

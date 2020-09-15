// +build e2e

package e2e

import (
	"github.com/neuvector/neuvector-nexus-iq/integration"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNeuVectorClient_Authenticate(t *testing.T) {
	t.Skip()

	// Configuration
	config := newTestConfig()

	// NeuVector client
	c, err := integration.NewNeuVectorClientFromConfig(config.NeuVectorConfig)
	if err != nil {
		t.Error(err)
	}

	// Authenticate with NeuVector controller
	token, err := c.GetAuthToken()
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, token)
}

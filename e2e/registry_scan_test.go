// +build e2e

package e2e

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-openapi/swag"
	"github.com/neuvector/neuvector-nexus-iq/integration"
	"github.com/neuvector/neuvector-nexus-iq/internal/logger"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/client/scan"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/models"
	"github.com/sirupsen/logrus"
)

func TestRegistryScan(t *testing.T) {
	// Test id
	testId := newTestId()
	t.Logf("Start registry scan integration test (id = %s)", testId)

	// Test logging
	testLogger, logRecorder := logger.NewChannelLogger(128)

	l, err := logger.NewCustomLogrusLogger(testLogger)
	if err != nil {
		t.Fatal(err)
	}
	integration.SetLogger(l)

	// Configuration
	config := newTestConfig()

	// NeuVector client
	c, err := integration.NewNeuVectorClientFromConfig(config.NeuVectorConfig)
	if err != nil {
		t.Fatal(err)
	}

	// # Setup

	// ## Create test registry
	testRegistryName := fmt.Sprintf("test-%s", testId)
	testRegistry := &models.RESTRegistryConfig{
		RegistryType: swag.String("Docker Registry"),
		Name:         swag.String(testRegistryName),
		Registry:     "http://127.0.0.1:5000/",
		Filters: []string{
			//"debian:jessie-20200607",
			//"localhost:5000/debian:jessie-20200607",
			//"*/debian:jessie-20200607",
			"*",
		},
		RescanAfterDbUpdate: false,
		ScanLayers:          false,
		Schedule:            nil,
	}
	err = c.CreateRegistry(testRegistry)
	if err != nil {
		t.Fatal(err)
	}

	// Wait for registry to be available
	getRegistryOperation := func() error {
		_, err := c.GetRegistryByName(testRegistryName)
		return err
	}
	err = retryWithConstantBackoff(getRegistryOperation, 1*time.Second, 10)
	if err != nil {
		t.Fatal(err)
	}

	// # Teardown

	// ## Delete test registry
	defer func() {
		_ = c.DeleteRegistry(testRegistryName)
		// FOLLOW UP delete registry fails sometimes if executed shortly after registry has been created; this must not fail the test
		//if err != nil {
		//	t.Error(err)
		//}
	}()

	// Start integration server
	intServer := newIntegrationServer()
	intServerStop := make(chan interface{})

	go func() {
		intServer.Start()
		close(intServerStop)
	}()

	// Graceful shutdown integration server after test
	defer func() {
		_ = intServer.ShutdownWithTimeout(5 * time.Second)
	}()

	// Wait for integration server health check
	err = waitForIntegrationServer(&intServer.Config)
	if err != nil {
		t.Fatal(err)
	}

	// Authenticate with NeuVector controller
	token, err := c.GetAuthToken()
	if err != nil {
		t.Error(err)
	}

	// Trigger registry scan
	// - POST /scan/registry/{name}/scan
	t.Logf("Start registry scan")
	_, err = c.Root.Scan.PostScanRegistryNameScan((&scan.PostScanRegistryNameScanParams{
		XAuthToken: token,
		Name:       testRegistryName,
	}).WithTimeout(c.Timeout))
	if err != nil {
		t.Error(err)
	}

	// Handle test events
loop:
	for {
		select {
		case entry := <-logRecorder.Entries:
			// Succeed if log with event type container report appears
			if entry.Data != nil {
				if eventType, hasEventType := entry.Data["event_type"]; hasEventType && eventType == "registry_report" {
					break loop
				}
			}

			// Fail if an error log appears
			if entry.Level <= logrus.ErrorLevel {
				t.Error(entry)
				break loop
			}
		case <-intServerStop:
			t.Fatal("integration server has stopped unexpectedly")
		case <-time.After(30 * time.Second):
			// Test has timed out
			t.Fatal("test timed out")
		}
	}

	// Assert
	// FOLLOWUP assert vulnerability scan report in NexusIq using the Get scan report history api
	// FOLLOWUP register a webhook with Nexus IQ and assert inbound call from Nexus IQ
	// https://help.sonatype.com/iqserver/automating/rest-apis/report-related-rest-apis---v2#Report-relatedRESTAPIs-v2-GettheScanReportHistory
}

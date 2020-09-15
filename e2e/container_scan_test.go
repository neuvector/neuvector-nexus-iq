// +build e2e

package e2e

import (
	"bytes"
	"fmt"
	"github.com/neuvector/neuvector-nexus-iq/integration"
	"github.com/neuvector/neuvector-nexus-iq/internal/logger"
	"github.com/sirupsen/logrus"
	"os/exec"
	"testing"
	"time"
)

func TestContainerScan(t *testing.T) {
	// Test id
	testId := newTestId()
	t.Logf("Start container scan integration test (id = %s)", testId)

	// Test logging
	testLogger, logRecorder := logger.NewChannelLogger(128)

	l, err := logger.NewCustomLogrusLogger(testLogger)
	if err != nil {
		t.Fatal(err)
	}
	integration.SetLogger(l)

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

	// Start container workload
	// FOLLOWUP consider replacing system exec calls by the Docker go SDK
	// - https://docs.docker.com/engine/api/sdk/examples/
	// - detached
	// --label "com.sonatype.nexus.iq.applicationName=Test Application"
	// docker run -d --name "test-workload" --rm debian:jessie-20200607-slim /bin/bash -c "tail -f /dev/null"
	testContainerName := fmt.Sprintf("test-%s", testId)
	dockerRunCmd := exec.Command("docker", "run", "-d", "--rm", "--name", testContainerName, "--label", "com.sonatype.nexus.iq.applicationName=Test Application", "debian:jessie-20200607-slim", "/bin/bash", "-c", "tail -f /dev/null")
	var dockerRunOut bytes.Buffer
	dockerRunCmd.Stdout = &dockerRunOut
	err = dockerRunCmd.Run()
	if err != nil {
		t.Fatal(newTestError("failed to start container workload", err))
	}

	// Stop container workload
	// docker stop "test-workload"
	defer func() {
		dockerStopCmd := exec.Command("docker", "stop", testContainerName)
		err := dockerStopCmd.Run()
		if err != nil {
			t.Error(newTestError("failed to stop container workload", err))
		}
	}()

	// Handle test events
loop:
	for {
		select {
		case entry := <-logRecorder.Entries:
			// Succeed if log with event type container report appears
			if entry.Data != nil {
				if eventType, hasEventType := entry.Data["event_type"]; hasEventType && eventType == "container_report" {
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
}

// +build e2e

package e2e

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Setup
	log.Println("Setup integration tests")

	// FOLLOWUP waitForNeuVectorController with retry; call isAvailable in client
	// FOLLOWUP waitForNexusIq with retry; call isAvailable in client

	// Setup NeuVector Controller
	err := setupNeuVectorController(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Setup NexusIQ
	err = setupNexusIq(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Run integration tests
	exitVal := m.Run()

	// Teardown
	log.Println("Teardown integration tests")
	// nothing

	os.Exit(exitVal)
}

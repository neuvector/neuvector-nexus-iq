// +build e2e

package e2e

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"
	"testing"

	"github.com/go-openapi/swag"
	"github.com/neuvector/neuvector-nexus-iq/integration"
	"github.com/neuvector/neuvector-nexus-iq/nexusiq"
	"github.com/neuvector/neuvector-nexus-iq/nexusiq/client/license"
)

const (
	// Path to the Nexus IQ license relative to the current package path
	nexusIqLicensePath = "./../test/licenses/nexusiq.lic"
)

var setupNexusIqOnce sync.Once

func setupNexusIq(config *integration.RootConfig) error {
	var err error

	setupNexusIqOnce.Do(func() {
		if config == nil {
			config = newTestConfig()
		}

		err = setupNexusIqWithConfig(*config)
	})

	return err
}

func setupNexusIqWithConfig(config integration.RootConfig) error {
	// Configure client
	nxUrl, err := url.Parse(config.NexusIqConfig.Endpoint)
	if err != nil {
		return fmt.Errorf("invalid Nexus IQ endpoint %s", config.NexusIqConfig.Endpoint)
	}

	// Create client instance
	c := nexusiq.NewClient(&nexusiq.Client{
		Schemes: []string{nxUrl.Scheme},
		Host:    nxUrl.Host,
		// License endpoints are not below the default base path
		BasePath: swag.String(""),
		Insecure: config.NexusIqConfig.Insecure,
		Username: config.NexusIqConfig.Username,
		Password: config.NexusIqConfig.Password,
	})

	// Get license status
	hasLicense := false

	getLicenseResponse, err := c.Root.License.GetLicense((&license.GetLicenseParams{}).WithTimeout(c.Timeout), c.BasicAuth())
	if err != nil {
		if _, paymentRequired := err.(*license.GetLicensePaymentRequired); paymentRequired {
			// License is not configured
			hasLicense = false
		} else {
			return err
		}
	}

	if getLicenseResponse != nil && getLicenseResponse.Payload != nil {
		hasLicense = true
	}

	// Apply license
	if hasLicense {
		log.Println("Nexus Iq license is applied")
	} else {
		log.Printf("Apply Nexus Iq license from file %s\n", nexusIqLicensePath)

		licenseReader, err := os.Open(nexusIqLicensePath)
		if err != nil {
			return err
		}
		defer licenseReader.Close()
		// TODO fix error 406 not acceptable when configuring the Nexus IQ license
		_, err = c.Root.License.SetLicense((&license.SetLicenseParams{
			File: licenseReader,
		}).WithTimeout(c.Timeout), c.BasicAuth())
		if err != nil {
			return err
		}
	}

	return nil
}

func TestSetupNexusIq(t *testing.T) {
	t.Skip()

}

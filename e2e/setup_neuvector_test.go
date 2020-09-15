// +build e2e

package e2e

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"testing"

	"github.com/go-openapi/swag"

	"github.com/neuvector/neuvector-nexus-iq/integration"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/client/e_u_l_a"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/client/response_rule"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/client/scan"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/client/system"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/models"
	"github.com/stretchr/testify/assert"
)

const (
	// Path to the NeuVector license relative to the current package path
	neuVectorLicensePath = "./../test/licenses/neuvector.txt"
)

var setupNeuVectorControllerOnce sync.Once

func setupNeuVectorController(config *integration.RootConfig) error {
	var err error

	setupNeuVectorControllerOnce.Do(func() {
		if config == nil {
			config = newTestConfig()
		}

		err = setupNeuVectorControllerWithConfig(*config)
	})

	return err
}

func setupNeuVectorControllerWithConfig(config integration.RootConfig) error {
	// Configure client
	c, err := integration.NewNeuVectorClientFromConfig(config.NeuVectorConfig)
	if err != nil {
		return err
	}

	// Authenticate with NeuVector controller
	token, err := c.GetAuthToken()
	if err != nil {
		return err
	}

	// Assert version
	// FOLLOWUP require the minimum version
	// using GET /controller
	// requires Controller tag

	// Get EULA status
	eulaStatusResponse, err := c.Root.Eula.GetEula((&e_u_l_a.GetEulaParams{
		XAuthToken: token,
	}).WithTimeout(c.Timeout))
	if err != nil {
		return err
	}

	// EULA
	eulaAccepted := false

	if p := eulaStatusResponse.Payload; p != nil {
		if e := p.Eula; e != nil {
			if e.Accepted != nil {
				eulaAccepted = *e.Accepted
			}
		}
	}

	if eulaAccepted {
		log.Println("EULA is accepted")
	} else {
		// Accept EULA
		log.Println("Accept EULA")
		err = c.AcceptEula()
		if err != nil {
			return err
		}
	}

	// Get license status
	getLicenseResp, err := c.Root.System.GetSystemLicense((&system.GetSystemLicenseParams{
		XAuthToken: token,
	}).WithTimeout(c.Timeout))
	if err != nil {
		return err
	}

	hasLicense := false

	if p := getLicenseResp.Payload; p != nil {
		hasLicense = p.License != nil && p.License.Info != nil
	}

	// Apply license
	if hasLicense {
		log.Println("License is applied")
	} else {
		log.Printf("Apply license from file %s\n", neuVectorLicensePath)

		licenseBytes, err := ioutil.ReadFile(neuVectorLicensePath)
		if err != nil {
			return err
		}

		license := string(licenseBytes)
		err = c.ApplyLicense(license)
		if err != nil {
			return err
		}
	}

	// System configuration struct
	systemConfig := newNeuVectorSystemConfig(&config)

	// System configuration
	log.Println("Apply NeuVector system configuration")
	_, err = c.Root.System.PatchSystemConfig((&system.PatchSystemConfigParams{
		XAuthToken: token,
		Body: &models.RESTSystemConfigConfigData{
			Config: systemConfig,
		},
	}).WithTimeout(c.Timeout))
	if err != nil {
		return err
	}

	// Required response rules
	registryScanEventResponseRule := &models.RESTResponseRule{
		Disable: false,
		Event:   swag.String("cve-report"),
		Conditions: []*models.RESTCLUSEventCondition{
			{
				Type:  "name",
				Value: "Registry.Scan.Report",
			},
		},
		Actions: []string{"webhook"},
		Group:   "",
		Comment: "e2e.Registry.Scan.Report",
	}

	containerScanEventResponseRule := &models.RESTResponseRule{
		Disable: false,
		Event:   swag.String("cve-report"),
		Conditions: []*models.RESTCLUSEventCondition{
			{
				Type:  "name",
				Value: "Container.Scan.Report",
			},
		},
		Actions: []string{"webhook"},
		Group:   "",
		Comment: "e2e.Container.Scan.Report",
	}

	// Get list of all response rules
	// GET /response/rule
	responseRulesResponse, err := c.Root.ResponseRule.GetResponseRule((&response_rule.GetResponseRuleParams{
		XAuthToken: token,
	}).WithTimeout(c.Timeout))
	if err != nil {
		return err
	}

	rules := responseRulesResponse.Payload.Rules
	hasRegistryScanEventResponseRule := false
	hasContainerScanEventResponseRule := false

	for _, rule := range rules {
		if swag.StringValue(rule.CfgType) != "user_created" {
			continue
		}

		// Identify response rule by comment
		switch rule.Comment {
		case registryScanEventResponseRule.Comment:
			hasRegistryScanEventResponseRule = true
		case containerScanEventResponseRule.Comment:
			hasContainerScanEventResponseRule = true
		}
	}

	// Configure response rules to log to webhook
	// - PATCH /response/rule/{id}

	if !hasRegistryScanEventResponseRule {
		// Create registry scan event response rule
		err := c.CreateResponseRule(registryScanEventResponseRule)
		if err != nil {
			return err
		}
	}

	if !hasContainerScanEventResponseRule {
		// Create container scan event response rule
		err := c.CreateResponseRule(containerScanEventResponseRule)
		if err != nil {
			return err
		}
	}

	// Configure public Dockerhub registry
	// -/scan/registry/{name}
	publicDockerHub := &models.RESTRegistryConfig{
		RegistryType: swag.String("Docker Registry"),
		Name:         swag.String("dockerhub-public"),
		Registry:     "https://registry.hub.docker.com/",
		Filters: []string{
			"debian:jessie-20200607",
		},
		RescanAfterDbUpdate: false,
		ScanLayers:          false,
		Schedule:            nil,
	}

	err = c.CreateOrUpdateRegistry(publicDockerHub)
	if err != nil {
		return err
	}

	// Enable container auto scan
	// - required to trigger container scan reports from starting containers
	// - /scan/config
	log.Println("Enable auto scan")
	_, err = c.Root.Scan.PatchScanConfig((&scan.PatchScanConfigParams{
		XAuthToken: token,
		Body: &models.RESTScanConfigData{
			Config: &models.RESTScanConfig{
				AutoScan: swag.Bool(true),
			},
		},
	}).WithTimeout(c.Timeout))
	if err != nil {
		return err
	}

	return nil
}

func newNeuVectorSystemConfig(config *integration.RootConfig) *models.RESTSystemConfigConfig {
	systemConfig := &models.RESTSystemConfigConfig{}

	// Configure registry proxy
	// - only relevant for Docker Desktop environment
	systemConfig.RegistryHTTPProxyStatus = true
	systemConfig.RegistryHTTPProxy = &models.RESTProxy{
		URL: swag.String("http://gateway.docker.internal:3128"),
	}
	systemConfig.RegistryHTTPSProxyStatus = true
	systemConfig.RegistryHTTPSProxy = &models.RESTProxy{
		URL: swag.String("http://gateway.docker.internal:3129"),
	}

	// Configure webhook
	systemConfig.WebhookStatus = true
	systemConfig.WebhookURL = fmt.Sprintf("http://host.docker.internal:%d/webhook", config.Port)

	return systemConfig
}

func TestSetupNeuVectorController(t *testing.T) {
	t.Skip()

	// Configure client
	config := newTestConfig()
	c, err := integration.NewNeuVectorClientFromConfig(config.NeuVectorConfig)
	if err != nil {
		t.Error(err)
	}

	// Authenticate with NeuVector controller
	token, err := c.GetAuthToken()
	if err != nil {
		t.Error(err)
	}

	// Assert the NeuVector Controller configuration
	expectedSystemConfig := newNeuVectorSystemConfig(config)
	systemConfigResponse, err := c.Root.System.GetSystemConfig((&system.GetSystemConfigParams{
		XAuthToken: token,
	}).WithTimeout(c.Timeout))
	if err != nil {
		t.Error(err)
	}

	actualSystemConfig := systemConfigResponse.Payload.Config

	assert.NotNil(t, actualSystemConfig)
	assert.Equal(t, expectedSystemConfig.RegistryHTTPProxyStatus, swag.BoolValue(actualSystemConfig.RegistryHTTPProxyStatus))
	assert.NotNil(t, actualSystemConfig.RegistryHTTPProxy)
	assert.Equal(t, expectedSystemConfig.RegistryHTTPProxy.URL, actualSystemConfig.RegistryHTTPProxy.URL)

	assert.Equal(t, expectedSystemConfig.RegistryHTTPSProxyStatus, swag.BoolValue(actualSystemConfig.RegistryHTTPSProxyStatus))
	assert.NotNil(t, actualSystemConfig.RegistryHTTPSProxy)
	assert.Equal(t, expectedSystemConfig.RegistryHTTPSProxy.URL, actualSystemConfig.RegistryHTTPSProxy.URL)
}

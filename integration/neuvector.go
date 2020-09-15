package integration

import (
	"fmt"
	"github.com/go-openapi/swag"
	"github.com/neuvector/neuvector-nexus-iq/internal/logger"
	"github.com/neuvector/neuvector-nexus-iq/neuvector"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/client/scan"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/models"
	"github.com/neuvector/neuvector-nexus-iq/nexusiq"
	nx_models "github.com/neuvector/neuvector-nexus-iq/nexusiq/models"
	"net/url"
)

func NewNeuVectorClientFromConfig(c NeuVectorConfig) (*neuvector.Client, error) {
	// Parse endpoint
	nvUrl, err := url.Parse(c.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid NeuVector Controller endpoint %s", c.Endpoint)
	}

	// Create client instance
	client := neuvector.NewClient(&neuvector.Client{
		Schemes:  []string{nvUrl.Scheme},
		Host:     nvUrl.Host,
		Insecure: c.Insecure,
		Username: c.Username,
		Password: c.Password,
		//Debug:    true,
	})

	return client, nil
}

type NeuVectorNexusIq struct {
	Config RootConfig

	NeuVectorClient *neuvector.Client
	NexusIqClient   *nexusiq.Client
}

type neuVectorNexusIq interface {
	ProcessWebhookRequest(*neuvector.WebhookRequest) error
}

// ProcessWebhookRequest processes a webhook request from the NeuVector controller
func (i *NeuVectorNexusIq) ProcessWebhookRequest(wr *neuvector.WebhookRequest) error {
	// Parse web hook data
	wd, err := wr.ToWebhookData()
	if err != nil {
		return err
	}

	// Process by event
	switch wd.Event {
	case neuvector.EventRegistryScanReport:
		return i.processRegistryScanEvent(wd)
	case neuvector.EventContainerScanReport:
		return i.processContainerScanEvent(wd)
	default:
		return fmt.Errorf("event %s is not supported", wd.Event)
	}

}

func (i *NeuVectorNexusIq) processRegistryScanEvent(wd *neuvector.WebhookData) error {
	nvc := i.Config.NeuVectorConfig
	nxc := i.Config.NexusIqConfig

	rsrd, err := wd.ToWebhookRegistryScanReportData()
	if err != nil {
		return err
	}

	// NeuVector Client
	nv, err := NewNeuVectorClientFromConfig(nvc)
	if err != nil {
		return err
	}

	// NexusIQ Client
	nx, err := NewNexusIqClientFromConfig(nxc)
	if err != nil {
		return err
	}

	// Authenticate with NeuVector
	nvToken, err := nv.GetAuthToken()
	if err != nil {
		return err
	}

	// Get registry scan report
	// GET /scan/registry/{name}/image/{id}
	log.Infof("get registry scan report for image id %s in registry %s", rsrd.ImageId, rsrd.RegistryName)
	nvReportResponse, err := nv.Root.Scan.GetScanRegistryNameImageID((&scan.GetScanRegistryNameImageIDParams{
		XAuthToken: nvToken,
		Name:       rsrd.RegistryName,
		ID:         rsrd.ImageId,
	}).WithTimeout(nv.Timeout))
	if err != nil {
		return err
	}

	// Get NexusIQ organization
	nexusOrganizationId, err := nexusGetOrganizationByIdOrName(nx, nxc.OrganizationId, nxc.OrganizationName)
	if err != nil {
		return err
	}

	// Get NexusIQ application
	var nxApp *nx_models.Application

	// Get NexusIQ application by report labels
	if nxc.AppNameLabel != "" {
		if appNameLabelValue, hasLabel := lookupReportLabelValue([]*models.RESTScanReportLabels{nvReportResponse.Payload.Report.Labels}, nxc.AppNameLabel); hasLabel && appNameLabelValue != "" {
			nxApp, err = nexusGetOrCreateApplication(nx, appNameLabelValue, appNameLabelValue, nexusOrganizationId)
			if err != nil {
				return err
			}
		}
	}

	// Get NexusIQ application inferred from repository and tag
	if nxApp == nil {
		// assume application public id corresponds to "[repository]-[tag]"
		nexusAppPublicId := fmt.Sprintf("%s-%s", rsrd.Repository, rsrd.Tag)
		nexusAppName := nexusAppPublicId

		nxApp, err = nexusGetOrCreateApplication(nx, nexusAppPublicId, nexusAppName, nexusOrganizationId)
		if err != nil {
			return err
		}
	}

	if nxApp == nil {
		return fmt.Errorf("failed to infer or create a Neuxs IQ application")
	}

	nexusAppId := swag.StringValue(nxApp.ID)

	// Map NeuVector report to CycloneDX
	if nvReportResponse.Payload == nil || nvReportResponse.Payload.Report == nil {
		return fmt.Errorf("GET /scan/registry/{name}/image/{id} did not provide a scan report")
	}

	err = neuVectorScanReportToNexusIq(nvReportResponse.Payload.Report, nx, nexusAppId, i.Config.NexusIqConfig.Source)
	if err != nil {
		return err
	}

	// Success log
	log.WithFields(logger.Fields{
		"event_type":    "registry_report",
		"image_id":      rsrd.ImageId,
		"registry_name": rsrd.RegistryName,
	}).Infof("sent NeuVector registry scan report for image id %[1]s in registry %[2]s to Nexus IQ application %[3]s", rsrd.ImageId, rsrd.RegistryName, nexusAppId)

	return nil
}

func (i *NeuVectorNexusIq) processContainerScanEvent(wd *neuvector.WebhookData) error {
	nvc := i.Config.NeuVectorConfig
	nxc := i.Config.NexusIqConfig

	csrd, err := wd.ToWebhookContainerScanReportData()
	if err != nil {
		return err
	}

	// NeuVector Client
	nv, err := NewNeuVectorClientFromConfig(nvc)
	if err != nil {
		return err
	}

	// NexusIQ Client
	nx, err := NewNexusIqClientFromConfig(nxc)
	if err != nil {
		return err
	}

	// Authenticate with NeuVector
	nvToken, err := nv.GetAuthToken()
	if err != nil {
		return err
	}

	// Get container scan report
	// /scan/workload/{id}
	log.Infof("get container scan report for workload id %s", csrd.WorkloadId)
	nvReportResponse, err := nv.Root.Scan.GetScanWorkloadID((&scan.GetScanWorkloadIDParams{
		XAuthToken: nvToken,
		ID:         csrd.WorkloadId,
	}).WithTimeout(nv.Timeout))
	if err != nil {
		return err
	}

	// Get NexusIQ organization
	nexusOrganizationId, err := nexusGetOrganizationByIdOrName(nx, nxc.OrganizationId, nxc.OrganizationName)

	// Get NexusIQ application
	var nxApp *nx_models.Application

	// Get NexusIQ application by report labels
	if nxc.AppNameLabel != "" {
		if appNameLabelValue, hasLabel := lookupReportLabelValue([]*models.RESTScanReportLabels{nvReportResponse.Payload.Report.Labels}, nxc.AppNameLabel); hasLabel && appNameLabelValue != "" {
			nxApp, err = nexusGetOrCreateApplication(nx, appNameLabelValue, appNameLabelValue, nexusOrganizationId)
			if err != nil {
				return err
			}
		}
	}

	// Get NexusIQ application inferred from workload image
	if nxApp == nil {
		nexusAppName := nexusiq.SanitizeApplicationName(csrd.WorkloadImage)
		nexusAppPublicId := nexusiq.SanitizeApplicationPublicId(csrd.WorkloadImage)

		nxApp, err = nexusGetOrCreateApplication(nx, nexusAppPublicId, nexusAppName, nexusOrganizationId)
		if err != nil {
			return err
		}
	}

	if nxApp == nil {
		return fmt.Errorf("failed to infer or create a Neuxs IQ application")
	}

	nexusAppId := swag.StringValue(nxApp.ID)

	// Map NeuVector report to CycloneDX
	if nvReportResponse.Payload == nil || nvReportResponse.Payload.Report == nil {
		return fmt.Errorf("GET /scan/workload/{id} did not provide a scan report")
	}

	err = neuVectorScanReportToNexusIq(nvReportResponse.Payload.Report, nx, nexusAppId, i.Config.NexusIqConfig.Source)
	if err != nil {
		return err
	}

	// Success log
	log.WithFields(logger.Fields{
		"event_type": "container_report",
	}).Infof("sent NeuVector container scan report for workload id %[1]s to Nexus IQ application %[2]s", csrd.WorkloadId, nexusAppId)

	return nil
}

func neuVectorScanReportToNexusIq(report *models.RESTScanReport, nx *nexusiq.Client, appId string, source string) error {
	bom, err := NeuVectorScanReportToCycloneDx(*report)
	if err != nil {
		return err
	}
	bomXml, err := bom.ToXml()
	if err != nil {
		return err
	}

	// Submit CycloneDX report to NexusIQ
	err = nx.SendThirdPartyScanReport(appId, source, bomXml, false)
	if err != nil {
		return err
	}

	return nil
}

// nexusGetOrCreateApplication creates a new Nexus IQ application if an there an application with the given public id does not exists.
// Returns the application or an error.
func nexusGetOrCreateApplication(nx *nexusiq.Client, appPublicId string, appName string, organizationId string) (*nx_models.Application, error) {
	nxApp, err := nx.GetApplicationByPublicId(appPublicId)
	if err != nil {
		switch err.(type) {
		case *nexusiq.ApplicationNotFoundError:
			// Create application
			nxApp, err = nx.CreateApplication(appName, appPublicId, organizationId)
			if err != nil {
				return nil, err
			}
		default:
			return nil, err
		}
	}

	return nxApp, nil
}

// nexusGetOrganizationByIdOrName returns the Nexus Organization by the provided id with first preference or by the provided name by second preference
func nexusGetOrganizationByIdOrName(nx *nexusiq.Client, orgId string, orgName string) (string, error) {
	if len(orgId) > 0 {
		// Return organization id from config
		return orgId, nil
	} else if len(orgName) > 0 {
		// Retrieve organization id by organization name
		organization, err := nx.GetOrganizationByName(orgName)
		if err != nil {
			return "", err
		}

		return swag.StringValue(organization.ID), nil
	}

	return "", fmt.Errorf("either OrganizationId or OrganizationName must be provided")
}

// lookupReportLabelValue returns the value of a label of a NeuVector RESTScanReportLabels or an empty string if the key is not contained
func lookupReportLabelValue(labels []*models.RESTScanReportLabels, key string) (string, bool) {
	for _, label := range labels {
		if label == nil || len(label.Key) != 1 || label.Key[0] != key {
			continue
		}

		return label.Value, true
	}

	return "", false
}

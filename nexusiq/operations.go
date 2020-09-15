package nexusiq

import (
	"fmt"
	"github.com/go-openapi/swag"
	"github.com/neuvector/neuvector-nexus-iq/nexusiq/client/application"
	"github.com/neuvector/neuvector-nexus-iq/nexusiq/client/organization"
	"github.com/neuvector/neuvector-nexus-iq/nexusiq/client/third_party_scan"
	"github.com/neuvector/neuvector-nexus-iq/nexusiq/models"
)

type ApplicationNotFoundError struct {
	PublicId string
}

func (e *ApplicationNotFoundError) Error() string {
	return fmt.Sprintf("application with public id %s not found", e.PublicId)
}

type AmbiguousApplicationsError struct {
	PublicId string
}

func (e *AmbiguousApplicationsError) Error() string {
	return fmt.Sprintf("ambiguous applications for public id %s", e.PublicId)
}

// GetApplication returns a single application by the public id
func (c *Client) GetApplicationByPublicId(publicId string) (*models.Application, error) {
	r, err := c.Root.Application.GetApplication((&application.GetApplicationParams{
		PublicID: &publicId,
	}).WithTimeout(c.Timeout), c.BasicAuth())
	if err != nil {
		return nil, err
	}

	apps := r.Payload.Applications

	if len(apps) == 0 {
		return nil, &ApplicationNotFoundError{PublicId: publicId}
	} else if len(apps) > 1 {
		return nil, &AmbiguousApplicationsError{PublicId: publicId}
	}

	return apps[0], nil
}

// CreateApplication creates an application and returns the created application
// name: valid characters are alphanumeric, "_", ".", "-", or spaces
// publicId: valid characters are alphanumeric, "_", "." or "-"
func (c *Client) CreateApplication(name string, publicId string, organizationId string) (*models.Application, error) {
	r, err := c.Root.Application.CreateApplication((&application.CreateApplicationParams{
		Application: &models.NewApplication{
			Name:           &name,
			PublicID:       &publicId,
			OrganizationID: &organizationId,
		},
	}).WithTimeout(c.Timeout), c.BasicAuth())
	if err != nil {
		return nil, err
	}

	app := r.Payload
	if app == nil {
		return nil, fmt.Errorf("failed to create application")
	}

	return app, nil
}

// SendThirdPartyScan submits a third party scan report to Nexus IQ
func (c *Client) SendThirdPartyScanReport(applicationInternalId string, source string, bom string, wait bool) error {
	if wait == true {
		return fmt.Errorf("wait for third party scan report processing is currently not implemented")
	}

	r, err := c.Root.ThirdPartyScan.PostScan((&third_party_scan.PostScanParams{
		ApplicationInternalID: applicationInternalId,
		Source:                source,
		Bom: &Bom{
			Xml: bom,
		},
	}).WithTimeout(c.Timeout), c.BasicAuth())
	if err != nil {
		return err
	}

	if r.Payload == nil || r.Payload.StatusURL == nil {
		return fmt.Errorf("send third party scan report did not yield a status url")
	}

	// FOLLOWUP if wait == true, poll status url for completion

	return nil
}

func (c *Client) GetOrganizationByName(name string) (*models.Organization, error) {
	r, err := c.Root.Organization.GetOrganizations((&organization.GetOrganizationsParams{
		OrganizationName: swag.String(name),
	}).WithTimeout(c.Timeout), c.BasicAuth())
	if err != nil {
		return nil, err
	}

	orgs := r.Payload.Organizations

	if len(orgs) == 0 {
		return nil, fmt.Errorf("organization not found")
	} else if len(orgs) > 1 {
		return nil, fmt.Errorf("ambiguous organizations found")
	}

	return orgs[0], nil
}

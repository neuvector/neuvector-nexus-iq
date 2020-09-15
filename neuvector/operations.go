package neuvector

import (
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/client/authentication"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/client/e_u_l_a"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/client/response_rule"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/client/scan"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/client/system"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/models"
	"time"
)

func (c *Client) GetAuthToken() (string, error) {

	if c.token == nil || time.Now().UTC().Add(10*time.Second).After(c.tokenExpiry) {
		// Get a token if no token is available or an existing is almost expired
		r, err := c.Root.Authentication.PostAuth((&authentication.PostAuthParams{
			Body: &models.RESTAuthData{
				Password: &models.RESTAuthPassword{
					Username: &c.Username,
					Password: &c.Password,
				},
			},
		}).WithTimeout(c.Timeout))

		if err != nil {
			return "", err
		}

		token := r.Payload.Token

		// Remember time when token was issued
		// - the RESTToken only contains a tokenTimeout
		c.tokenIssued = time.Now().UTC()

		var tokenTimeout uint32 = 0

		if token.Timeout != nil {
			tokenTimeout = *token.Timeout
		}

		// Remember token expiry time
		c.tokenExpiry = c.tokenIssued.Add(time.Duration(tokenTimeout) * time.Second)

		// Remember token
		c.token = token
	}

	if c.token.Token == nil {
		return "", ClientError("auth token not available")
	}

	token := *c.token.Token

	return token, nil
}

func (c *Client) AcceptEula() error {
	token, err := c.GetAuthToken()
	if err != nil {
		return err
	}

	_, err = c.Root.Eula.PostEula((&e_u_l_a.PostEulaParams{
		XAuthToken: token,
		Body: &models.RESTEULAData{
			Eula: &models.RESTEULA{
				Accepted: swag.Bool(true),
			},
		},
	}).WithTimeout(c.Timeout))

	return err
}

func (c *Client) ApplyLicense(licenseKey string) error {
	token, err := c.GetAuthToken()
	if err != nil {
		return err
	}

	_, err = c.Root.System.PostSystemLicenseUpdate((&system.PostSystemLicenseUpdateParams{
		XAuthToken: token,
		Body: &models.RESTLicenseKey{
			LicenseKey: &licenseKey,
		},
	}).WithTimeout(c.Timeout))

	return err
}

func (c *Client) GetRegistryByName(registryName string) (*models.RESTRegistrySummary, error) {
	token, err := c.GetAuthToken()
	if err != nil {
		return nil, err
	}

	r, err := c.Root.Scan.GetScanRegistryName((&scan.GetScanRegistryNameParams{
		XAuthToken: token,
		Name:       registryName,
	}).WithTimeout(c.Timeout))
	if err != nil {
		return nil, err
	}

	if r.Payload != nil && r.Payload.Summary != nil {
		return r.Payload.Summary, nil
	}

	return nil, fmt.Errorf("registry not found")
}

func (c *Client) CreateRegistry(r *models.RESTRegistryConfig) error {
	token, err := c.GetAuthToken()
	if err != nil {
		return err
	}

	_, err = c.Root.Scan.PostScanRegistry((&scan.PostScanRegistryParams{
		XAuthToken: token,
		Body: &models.RESTRegistryConfigData{
			Config: r,
		},
	}).WithTimeout(c.Timeout))

	return err
}

func (c *Client) UpdateRegistry(r *models.RESTRegistryConfig) error {
	token, err := c.GetAuthToken()
	if err != nil {
		return err
	}

	_, err = c.Root.Scan.PatchScanRegistryName((&scan.PatchScanRegistryNameParams{
		XAuthToken: token,
		Name:       swag.StringValue(r.Name),
		Body: &models.RESTRegistryConfigData{
			Config: r,
		},
	}).WithTimeout(c.Timeout))

	return err
}

func (c *Client) CreateOrUpdateRegistry(r *models.RESTRegistryConfig) error {
	token, err := c.GetAuthToken()
	if err != nil {
		return err
	}

	// Get registry
	_, err = c.Root.Scan.GetScanRegistryName((&scan.GetScanRegistryNameParams{
		XAuthToken: token,
		Name:       swag.StringValue(r.Name),
	}).WithTimeout(c.Timeout))

	if err != nil {
		if apiErr, isApiErr := err.(*runtime.APIError); isApiErr {
			if apiErr.Code == 404 {
				// Create registry
				err = c.CreateRegistry(r)
				return err
			}
		} else {
			return err
		}
	}

	// Update registry
	err = c.UpdateRegistry(r)

	return err
}

func (c *Client) DeleteRegistry(registryName string) error {
	token, err := c.GetAuthToken()
	if err != nil {
		return err
	}

	_, err = c.Root.Scan.DeleteScanRegistryName((&scan.DeleteScanRegistryNameParams{
		XAuthToken: token,
		Name:       registryName,
	}).WithTimeout(c.Timeout))

	return err
}

func (c *Client) CreateResponseRule(rule *models.RESTResponseRule) error {
	token, err := c.GetAuthToken()
	if err != nil {
		return err
	}

	_, err = c.Root.ResponseRule.PatchResponseRule((&response_rule.PatchResponseRuleParams{
		XAuthToken: token,
		Body: &models.RESTResponseRuleActionData{
			Insert: &models.RESTResponseRuleInsert{
				Rules: []*models.RESTResponseRule{
					rule,
				},
			},
		},
	}).WithTimeout(c.Timeout))
	if err != nil {
		return err
	}

	return err
}

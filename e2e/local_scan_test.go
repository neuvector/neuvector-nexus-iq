// +build e2e

package e2e

import (
	"github.com/go-openapi/swag"
	"github.com/neuvector/neuvector-nexus-iq/neuvector"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/client/scan"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocalScan(t *testing.T) {
	t.Skip()

	// FOLLOWUP implement TestLocalScan scenario

	c := neuvector.NewClient(&neuvector.Client{
		Username: "admin",
		Password: "admin",
	})

	token, err := c.GetAuthToken()
	if err != nil {
		t.Fatal(err)
	}

	r, err := c.Root.Scan.PostScanRepository((&scan.PostScanRepositoryParams{
		XAuthToken: token,
		Body: &models.RESTScanRepoReqData{
			Request: &models.RESTScanRepoReq{
				Repository: swag.String("debian"),
				Tag:        swag.String("jessie-20200607-slim"),
				ScanLayers: swag.Bool(true),
			},
		},
	}).WithTimeout(c.Timeout))
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, *r.Payload.Report.Repository)
}

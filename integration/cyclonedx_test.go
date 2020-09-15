package integration

import (
	"encoding/json"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNeuVectorScanRepoReportToCycloneDx_Marshal(t *testing.T) {
	// Load NeuVector scan report
	var nvScanReport models.RESTScanRepoReportData
	err := json.Unmarshal([]byte(nvRepositoryScanReportDebianImageJson), &nvScanReport)
	if err != nil {
		t.Error(err)
	}

	// Convert NeuVector report to CycloneDX
	bom, err := NeuVectorScanRepoReportToCycloneDx(nvScanReport)
	if err != nil {
		t.Error(err)
	}

	// Fix bom serial number
	bom.SerialNumberAttr = "urn:uuid:53763c4a-cb05-42b9-bd5f-8b733dba5682"

	// Marshal CycloneDX struct to XML
	bomXml, err := bom.ToXml()
	if err != nil {
		t.Fatal(err)
	}

	// Assert CycloneDX XML
	assert.Equal(t, nvRepositoryScanReportDebianImageBomXml, bomXml)
}

func TestNeuVectorScanReportToCycloneDx_Marshal(t *testing.T) {
	// Load NeuVector scan report
	var nvScanReport models.RESTScanReportData
	err := json.Unmarshal([]byte(nvRegistryScanReportDebianJson), &nvScanReport)
	if err != nil {
		t.Error(err)
	}

	// Convert NeuVector report to CycloneDX
	bom, err := NeuVectorScanReportToCycloneDx(*nvScanReport.Report)
	if err != nil {
		t.Error(err)
	}

	// Fix bom serial number to allow comparison with test fixture
	bom.SerialNumberAttr = "urn:uuid:53763c4a-cb05-42b9-bd5f-8b733dba5682"

	// Marshal CycloneDX struct to XML
	bomXml, err := bom.ToXml()
	if err != nil {
		t.Fatal(err)
	}

	// Assert marshalled CycloneDX XML
	assert.Equal(t, nvRegistryScanReportDebianImageBomXml, bomXml)
}

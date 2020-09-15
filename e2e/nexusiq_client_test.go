// +build e2e

package e2e

import (
	"github.com/go-openapi/swag"
	"github.com/neuvector/neuvector-nexus-iq/integration"
	"github.com/neuvector/neuvector-nexus-iq/nexusiq"
	"github.com/neuvector/neuvector-nexus-iq/nexusiq/client/application"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testBom = `<?xml version="1.0"?>
<bom serialNumber="urn:uuid:3e671687-395b-41f5-a30f-a58921a69b79" version="1"
     xmlns="http://cyclonedx.org/schema/bom/1.1"
     xmlns:v="http://cyclonedx.org/schema/ext/vulnerability/1.0">
  <components>
    <component type="library" bom-ref="pkg:maven/com.fasterxml.jackson.core/jackson-databind@2.9.9?type=jar">
      <publisher>FasterXML</publisher>
      <group>com.fasterxml.jackson.core</group>
      <name>jackson-databind</name>
      <version>2.9.9</version>
      <purl>pkg:maven/com.fasterxml.jackson.core/jackson-databind@2.9.9?type=jar</purl>
      <v:vulnerabilities>
        <v:vulnerability ref="pkg:maven/com.fasterxml.jackson.core/jackson-databind@2.9.9?type=jar">
          <v:id>CVE-2018-7489</v:id>
          <v:source name="NVD">
            <v:url>https://nvd.nist.gov/vuln/detail/CVE-2018-7489</v:url>
          </v:source>
          <v:ratings>
            <v:rating>
              <v:score>
                <v:base>9.8</v:base>
                <v:impact>5.9</v:impact>
                <v:exploitability>3.0</v:exploitability>
              </v:score>
              <v:severity>Critical</v:severity>
              <v:method>CVSSv3</v:method>
              <v:vector>AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H</v:vector>
            </v:rating>
          </v:ratings>
          <v:cwes>
            <v:cwe>184</v:cwe>
            <v:cwe>502</v:cwe>
          </v:cwes>
          <v:description>FasterXML jackson-databind before 2.7.9.3, 2.8.x before 2.8.11.1 and 2.9.x before 2.9.5 allows unauthenticated remote code execution because of an incomplete fix for the CVE-2017-7525 deserialization flaw. This is exploitable by sending maliciously crafted JSON input to the readValue method of the ObjectMapper, bypassing a blacklist that is ineffective if the c3p0 libraries are available in the classpath.</v:description>
          <v:recommendations>
            <v:recommendation>Upgrade</v:recommendation>
          </v:recommendations>
          <v:advisories>
            <v:advisory>https://github.com/FasterXML/jackson-databind/issues/1931</v:advisory>
            <v:advisory>http://www.securityfocus.com/bid/103203</v:advisory>
            <v:advisory>http://www.securitytracker.com/id/1040693</v:advisory>
            <v:advisory>http://www.securitytracker.com/id/1041890</v:advisory>
          </v:advisories>
        </v:vulnerability>
      </v:vulnerabilities>
    </component>
  </components>
</bom>
`
)

func NewNexusIqClient() (*nexusiq.Client, error) {
	// Configuration
	config := newTestConfig()

	// NeuVector client
	c, err := integration.NewNexusIqClientFromConfig(config.NexusIqConfig)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func TestNexusIqClient_GetSandboxApplication(t *testing.T) {
	t.Skip()

	c, err := NewNexusIqClient()
	if err != nil {
		t.Fatal(err)
	}

	publicId := "sandbox-application"

	r, err := c.Root.Application.GetApplication((&application.GetApplicationParams{
		PublicID: &publicId,
	}).WithTimeout(c.Timeout), c.BasicAuth())
	if err != nil {
		t.Fatal(err)
	}

	assert.NotEmpty(t, r.Payload.Applications)
}

func TestNexusIqClient_GetOrganizationByName(t *testing.T) {
	t.Skip()

	c, err := NewNexusIqClient()
	if err != nil {
		t.Fatal(err)
	}

	organizationName := "Sandbox Organization"

	organization, err := c.GetOrganizationByName(organizationName)
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, organization)
	assert.Equal(t, organizationName, swag.StringValue(organization.Name))
	assert.NotEmpty(t, organization.ID)
}

func TestNexusIqClient_SendThirdPartyScanReport(t *testing.T) {
	t.Skip()

	c, err := NewNexusIqClient()
	if err != nil {
		t.Fatal(err)
	}

	publicId := "sandbox-application"

	app, err := c.GetApplicationByPublicId(publicId)
	if err != nil {
		t.Fatal(err)
	}

	appId := *app.ID

	err = c.SendThirdPartyScanReport(appId, "neuvector", testBom, false)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, err)
}

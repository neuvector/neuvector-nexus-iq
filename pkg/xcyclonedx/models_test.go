package xcyclonedx

import (
	"encoding/xml"
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

func TestCycloneDx_Marshal(t *testing.T) {
	// Bom
	bom := Bom{
		XMLNs:            "http://cyclonedx.org/schema/bom/1.1",
		XMLNsV:           "http://cyclonedx.org/schema/ext/vulnerability/1.0",
		SerialNumberAttr: "urn:uuid:3e671687-395b-41f5-a30f-a58921a69b79",
		VersionAttr:      1,
		Components: &Components{
			Component: []*Component{},
		},
	}

	// Components
	component := &Component{
		TypeAttr:   "library",
		BomrefAttr: "pkg:maven/com.fasterxml.jackson.core/jackson-databind@2.9.9?type=jar",
		Publisher:  []string{"FasterXML"},
		Group:      []string{"com.fasterxml.jackson.core"},
		Name:       []string{"jackson-databind"},
		Version:    []string{"2.9.9"},
		Purl:       []string{"pkg:maven/com.fasterxml.jackson.core/jackson-databind@2.9.9?type=jar"},
		Vulnerabilities: &Vulnerabilities{
			Vulnerability: []*Vulnerability{},
		},
	}

	bom.Components.Component = append(bom.Components.Component, component)

	// Vulnerabilities
	vulnerability := &Vulnerability{
		RefAttr: "pkg:maven/com.fasterxml.jackson.core/jackson-databind@2.9.9?type=jar",
		Id:      []string{"CVE-2018-7489"},
		Source: []*Source{&Source{
			NameAttr: "NVD",
			Url:      []string{"https://nvd.nist.gov/vuln/detail/CVE-2018-7489"},
		}},
		Ratings: []*Ratings{
			&Ratings{
				Rating: []*ScoreType{
					&ScoreType{
						Score: []*Score{
							&Score{
								Base:           []float64{9.8},
								Impact:         []float64{5.9},
								Exploitability: []float64{3.0},
							},
						},
						Severity: []string{"Critical"},
						Method:   []string{"CVSSv3"},
						Vector:   []string{"AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H"},
					},
				},
			},
		},
	}

	component.Vulnerabilities.Vulnerability = append(component.Vulnerabilities.Vulnerability, vulnerability)

	// Marshal to xml string
	marshalledBom, err := xml.MarshalIndent(bom, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	strBom := string(marshalledBom)
	t.Log(strBom)
}

func TestCycloneDx_Unmarshal(t *testing.T) {

}

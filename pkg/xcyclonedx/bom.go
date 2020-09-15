package xcyclonedx

import (
	"encoding/xml"
	"github.com/google/uuid"
)

func NewBom(serialNumber string) *Bom {
	return &Bom{
		XMLNs:            "http://cyclonedx.org/schema/bom/1.1",
		SerialNumberAttr: serialNumber,
		VersionAttr:      1,
		Components: &Components{
			Component: []*Component{},
		},
	}
}

func (b *Bom) WithVulnerabilities() *Bom {
	b.XMLNsV = "http://cyclonedx.org/schema/ext/vulnerability/1.0"
	b.Vulnerabilities = &Vulnerabilities{
		Vulnerability: []*Vulnerability{},
	}

	return b
}

func (b *Bom) ToXml() (string, error) {
	xmlBytes, err := xml.Marshal(b)
	if err != nil {
		return "", err
	}
	return string(xmlBytes), nil
}

// NewSerialNumber generates a serial number
// Example: "urn:uuid:3e671687-395b-41f5-a30f-a58921a69b79"
func NewSerialNumber() string {
	return uuid.New().URN()
}

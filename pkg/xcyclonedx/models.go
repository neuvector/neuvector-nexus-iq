package xcyclonedx

import "encoding/xml"

// Bom ...
type Bom struct {
	XMLNs            string      `xml:"xmlns,attr"`
	XMLNsV           string      `xml:"xmlns:v,attr"`
	XMLName          xml.Name    `xml:"bom"`
	VersionAttr      int         `xml:"version,attr,omitempty"`
	SerialNumberAttr string      `xml:"serialNumber,attr,omitempty"`
	Components       *Components `xml:"components"`
	//ExternalReferences []*ExternalReferences `xml:"externalReferences"`
	Vulnerabilities *Vulnerabilities `xml:"v:vulnerabilities"`
}

// Components ...
type Components struct {
	XMLName   xml.Name     `xml:"components"`
	Component []*Component `xml:"component"`
}

// Component ...
type Component struct {
	XMLName     xml.Name `xml:"component"`
	TypeAttr    string   `xml:"type,attr"`
	BomrefAttr  string   `xml:"bom-ref,attr,omitempty"`
	Publisher   []string `xml:"publisher"`
	Group       []string `xml:"group"`
	Name        []string `xml:"name"`
	Version     []string `xml:"version"`
	Description []string `xml:"description"`
	Scope       []string `xml:"scope"`
	//Hashes             []*Hashes             `xml:"hashes"`
	//Licenses           []*Licenses           `xml:"licenses"`
	Copyright []string `xml:"copyright"`
	Cpe       []string `xml:"cpe"`
	Purl      []string `xml:"purl"`
	Modified  []bool   `xml:"modified"`
	//Pedigree           []*PedigreeType       `xml:"pedigree"`
	//ExternalReferences []*ExternalReferences `xml:"externalReferences"`
	//Components         []*Components         `xml:"components"`
	Vulnerabilities *Vulnerabilities `xml:"v:vulnerabilities"`
}

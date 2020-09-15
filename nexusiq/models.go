package nexusiq

type XmlContent interface {
	ToXml() string
}

// Bom contains the CycloneDX formatted XML BOM
type Bom struct {
	Xml string //`xml:",innerxml"`
}

func (b *Bom) ToXml() string {
	return b.Xml
}

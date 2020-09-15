package integration

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/neuvector/neuvector-nexus-iq/neuvector/models"
	cyclonedx "github.com/neuvector/neuvector-nexus-iq/pkg/xcyclonedx"
	"github.com/package-url/packageurl-go"
	"sort"
)

func NeuVectorScanRepoReportToCycloneDx(nvScanRepoReport models.RESTScanRepoReportData) (*cyclonedx.Bom, error) {
	// Generate serial number
	//serialNumber := "urn:uuid:3e671687-395b-41f5-a30f-a58921a69b79"
	serialNumber := cyclonedx.NewSerialNumber()

	// Create a CycloneDX BOM structure
	bom := cyclonedx.NewBom(serialNumber).WithVulnerabilities()

	// Map image meta data to CycloneDX component
	nvReport := nvScanRepoReport.Report

	// Image repository
	imageRepository := ""
	if nvReport.Repository != nil {
		imageRepository = *nvReport.Repository
	} else {
		err := fmt.Errorf("report does not contain an image repository")
		return nil, err
	}

	// Image digest
	imageDigest := ""
	if nvReport.Digest != nil {
		imageDigest = *nvReport.Digest
	} else {
		return nil, fmt.Errorf("report does not contain an image digest")
	}

	// Image registry
	imageRegistry := ""
	if nvReport.Registry != nil {
		imageRegistry = *nvReport.Registry
	}

	imagePurl := cyclonedx.NewDockerPackageUrl(imageRepository, imageDigest, imageRegistry)
	imagePurlStr := imagePurl.String()

	// CycloneDX component
	component := &cyclonedx.Component{
		TypeAttr:   cyclonedx.ApplicationComponent,
		BomrefAttr: imagePurlStr,
		Name:       []string{imageRepository},
		Version:    []string{imageDigest},
		Purl:       []string{imagePurlStr},
		Vulnerabilities: &cyclonedx.Vulnerabilities{
			Vulnerability: []*cyclonedx.Vulnerability{},
		},
	}

	bom.Components.Component = append(bom.Components.Component, component)

	// Use image package url as reference from vulnerability to component
	vulComponentRefFunc := func(v *models.RESTVulnerability) *packageurl.PackageURL {
		return imagePurl
	}

	// Vulnerabilities
	vulnerabilities, err := NeuVectorVulnerabilitiesToCycloneDxVulnerabilities(nvReport.Vulnerabilities, vulComponentRefFunc)
	if err != nil {
		return nil, err
	}

	component.Vulnerabilities = vulnerabilities

	return bom, nil
}

// Create a CycloneDX BOM with vulnerabilities extension from a NeuVector RESTScanReport
// In the resulting BOM, components are separate from the vulnerabilities
// https://cyclonedx.org/ext/vulnerability/#example-usage-bom-node
func NeuVectorScanReportToCycloneDx(nvScanReport models.RESTScanReport) (*cyclonedx.Bom, error) {
	// If true, the BOM will contain the vulnerabilties separated from the components
	var separateVulnerabilities = false

	// Dissect NeuVector report
	nvVuls := nvScanReport.Vulnerabilities

	// Generate serial number
	//serialNumber := "urn:uuid:3e671687-395b-41f5-a30f-a58921a69b79"
	serialNumber := uuid.New().URN()

	// Create a CycloneDX BOM structure
	bom := cyclonedx.NewBom(serialNumber).WithVulnerabilities()

	// Map CycloneDX components from NeuVector vulnerabilities
	components, err := NeuVectorVulnerabilitiesToCycloneDxComponents(nvVuls, NeuVectorVulnerabilityToGenericPurl, !separateVulnerabilities)
	if err != nil {
		return nil, err
	}
	bom.Components = components

	// Map CycloneDX vulnerabilities from NeuVector vulnerabilities
	if separateVulnerabilities {
		vulnerabilities, err := NeuVectorVulnerabilitiesToCycloneDxVulnerabilities(nvVuls, NeuVectorVulnerabilityToGenericPurl)
		if err != nil {
			return nil, err
		}

		// Set vulnerabilities as top level on the bom
		bom.Vulnerabilities = vulnerabilities
	}

	return bom, nil
}

// NeuVectorVulnerabilitiesToCycloneDxComponents maps a list of NeuVector RESTVulnerability to a list of CycloneDX components
// In the resulting list, components are unique with regard to package name and version
func NeuVectorVulnerabilitiesToCycloneDxComponents(nvVuls []*models.RESTVulnerability, vulComponentRefFunc VulnerabilityToComponentPurlFunc, includeVulnerabilities bool) (*cyclonedx.Components, error) {
	componentsMap := map[string]*cyclonedx.Component{}

	for _, nvVul := range nvVuls {
		packageName := nvVul.PackageName
		if packageName == nil || *packageName == "" {
			continue
		}

		packageVersion := nvVul.PackageVersion
		if packageVersion == nil || *packageVersion == "" {
			continue
		}

		// Create package url based on package name and version
		purl := cyclonedx.NewGenericOsPackageUrl(*packageName, *packageVersion)
		purlStr := purl.ToString()

		// Find component
		component, exists := componentsMap[purlStr]

		if !exists {
			// Create new component
			component = &cyclonedx.Component{
				TypeAttr:   cyclonedx.ApplicationComponent,
				BomrefAttr: purlStr,
				Name:       []string{*packageName},
				Version:    []string{*packageVersion},
				Purl:       []string{purlStr},
			}

			if includeVulnerabilities {
				component.Vulnerabilities = &cyclonedx.Vulnerabilities{
					Vulnerability: []*cyclonedx.Vulnerability{},
				}
			}

			componentsMap[purlStr] = component
		}

		// Optionally add vulnerability
		if includeVulnerabilities {
			vul, err := NeuVectorVulnerabilityToCycloneDxVulnerability(nvVul, vulComponentRefFunc)
			if err != nil {
				return nil, err
			}
			component.Vulnerabilities.Vulnerability = append(component.Vulnerabilities.Vulnerability, vul)
		}
	}

	// Create list from component map values
	var components []*cyclonedx.Component
	for _, component := range componentsMap {
		components = append(components, component)
	}

	// Sort component list by purl (to allow deterministic tests)
	sort.SliceStable(components, func(i, j int) bool {
		return components[i].BomrefAttr < components[j].BomrefAttr
	})

	// Components container in BOM
	componentsContainer := &cyclonedx.Components{
		Component: components,
	}

	return componentsContainer, nil
}

// VulnerabilityToComponentPurlFunc maps a NeuVector vulnerability to the package url of a component
type VulnerabilityToComponentPurlFunc = func(*models.RESTVulnerability) *packageurl.PackageURL

func NeuVectorVulnerabilityToGenericPurl(vul *models.RESTVulnerability) *packageurl.PackageURL {
	packageName := vul.PackageName
	if packageName == nil || *packageName == "" {
		return nil
	}

	packageVersion := vul.PackageVersion
	if packageVersion == nil || *packageVersion == "" {
		return nil
	}

	return cyclonedx.NewGenericOsPackageUrl(*packageName, *packageVersion)
}

// NeuVectorVulnerabilitiesCycloneDx maps from a list NeuVector RESTVulnerability to a list of CycloneDX vulnerabilities
func NeuVectorVulnerabilitiesToCycloneDxVulnerabilities(nvVuls []*models.RESTVulnerability, vulComponentRefFunc VulnerabilityToComponentPurlFunc) (*cyclonedx.Vulnerabilities, error) {
	var vuls []*cyclonedx.Vulnerability

	for _, nvVul := range nvVuls {
		vul, err := NeuVectorVulnerabilityToCycloneDxVulnerability(nvVul, vulComponentRefFunc)
		if err != nil {
			return nil, err
		}

		vuls = append(vuls, vul)
	}

	vulsContainer := &cyclonedx.Vulnerabilities{
		Vulnerability: vuls,
	}

	return vulsContainer, nil
}

// NeuVectorVulnerabilityToCycloneDxVulnerability maps a single NeuVector RESTVulnerability to a CycloneDX vulnerability
func NeuVectorVulnerabilityToCycloneDxVulnerability(nvVul *models.RESTVulnerability, componentRef VulnerabilityToComponentPurlFunc) (*cyclonedx.Vulnerability, error) {
	// ID/CVE
	// NV Vulnerability Name => BOM Vulnerability ID
	id, hasId := requireString(nvVul.Name)
	if !hasId {
		return nil, fmt.Errorf("vulnerability does not contain an id")
	}

	// Description
	// NV Vulnerability Description => BOM Vulnerability Description
	description := optionalString(nvVul.Description)

	// Advisories
	// NV Vulnerability Link => BOM Vulnerability Advisory
	var advisories []*cyclonedx.Advisories = nil

	link, hasLink := requireString(nvVul.Link)
	if hasLink {
		advisories = []*cyclonedx.Advisories{
			&cyclonedx.Advisories{
				Advisory: link,
			},
		}
	}

	// Ratings
	score, hasScore := requireFloat64(nvVul.ScoreV3)
	if !hasScore {
		return nil, fmt.Errorf("vulnerability does not contain a score")
	}

	var severity []string
	if s, hasSeverity := requireString(nvVul.Severity); hasSeverity {
		severity = []string{s}
	}

	var vector []string
	if v, hasVector := requireString(nvVul.VectorsV3); hasVector {
		vector = []string{v}
	}

	ratings := []*cyclonedx.Ratings{
		&cyclonedx.Ratings{
			Rating: []*cyclonedx.ScoreType{
				&cyclonedx.ScoreType{
					Score: []*cyclonedx.Score{
						&cyclonedx.Score{
							Base: []float64{score},
						},
					},
					Severity: severity,
					Method:   []string{"CVSSv3"},
					Vector:   vector,
				},
			},
		},
	}

	refAttr := componentRef(nvVul).ToString()
	vul := &cyclonedx.Vulnerability{
		RefAttr:     refAttr,
		Id:          []string{id},
		Description: []string{description},
		Advisories:  advisories,
		Ratings:     ratings,
	}

	return vul, nil
}

func requireString(v *string) (string, bool) {
	if v == nil {
		return "", false
	}

	return *v, true
}

func requireFloat64(v *float64) (float64, bool) {
	if v == nil {
		return 0, false
	}

	return *v, true
}

func optionalString(v *string) string {
	if v == nil {
		return ""
	}

	return *v
}

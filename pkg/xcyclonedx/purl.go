package xcyclonedx

import (
	"github.com/package-url/packageurl-go"
	"strings"
)

// NewDockerPackageUrl creates a PackageURL for a Docker image
// Namespace is the image registry
// Version should be the image sha256 or a tag. A sha256 is preferred
// Qualifier "repository_url" is mapped from an optional image registry
// https://github.com/package-url/purl-spec#rules-for-each-purl-component
func NewDockerPackageUrl(repository string, version string, registry string) *packageurl.PackageURL {
	var qualifiers []packageurl.Qualifier

	if registry != "" {
		qualifiers = append(qualifiers, packageurl.Qualifier{
			Key:   "repository_url",
			Value: registry,
		})
	}

	// Split repository into namespace and name
	// Example: repository = glibc/libc-bin => dockerNamespace = glibc; dockerName = libc-bin
	dockerRepositorySplitted := strings.Split(repository, "/")
	dockerNamespace := strings.Join(dockerRepositorySplitted[:len(dockerRepositorySplitted)-1], "")
	dockerName := strings.Join(dockerRepositorySplitted[len(dockerRepositorySplitted)-1:], "")

	return packageurl.NewPackageURL(packageurl.TypeDocker, dockerNamespace, dockerName, version, qualifiers, "")
}

func NewGenericOsPackageUrl(packageName string, packageVersion string) *packageurl.PackageURL {
	// Split packageName into namespace and name
	// Example: packageName = glibc/libc-bin => pkgNamespace = glibc; pkgName = libc-bin
	pkgNameSplitted := strings.Split(packageName, "/")
	pkgNamespace := strings.Join(pkgNameSplitted[:len(pkgNameSplitted)-1], "")
	pkgName := strings.Join(pkgNameSplitted[len(pkgNameSplitted)-1:], "")

	return packageurl.NewPackageURL(packageurl.TypeGeneric, pkgNamespace, pkgName, packageVersion, packageurl.Qualifiers{}, "")
}

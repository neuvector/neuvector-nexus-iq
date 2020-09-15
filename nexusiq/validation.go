package nexusiq

import "regexp"

// Valid characters for a Nexus IQ application name are alphanumeric, "_", ".", "-", or spaces
var invalidAppNameRe = regexp.MustCompile(`[^a-zA-z0-9_\.\- ]`)

// Validates an application name
func ValidateApplicationName(name string) bool {
	return invalidAppNameRe.MatchString(name)
}

// Cleans a potential application name from invalid characters
// Invalid characters are removed
func SanitizeApplicationName(name string) string {
	return invalidAppNameRe.ReplaceAllString(name, "")
}

// Valid characters for a Nexus IQ application public id are alphanumeric, "_", "." or "-"
var invalidAppPublicIdRe = regexp.MustCompile(`[^a-zA-z0-9_\.\-]`)

// Validates an application name
func ValidateApplicationPublicId(publicId string) bool {
	return invalidAppPublicIdRe.MatchString(publicId)
}

// Cleans a potential application public id from invalid characters
// Invalid characters are replaced by an underscore
func SanitizeApplicationPublicId(publicId string) string {
	return invalidAppNameRe.ReplaceAllString(publicId, "_")
}

package build

// These variables are intended to be set during the build process
var (
	// Version
	// - For a release version: "v1.2.1"
	// - For a non-release version: "261bc7f"
	// Output of "git describe --tags --always"
	Version = "dev"

	// Long commit hash
	// Example: "3ad37b5fdbc73dc29b067aedf34a6962d3001a6d"
	// Output of "git rev-parse HEAD"
	Commit = ""

	// Human-readable time of the build in UTC time zone
	// Example: "Sun Sep 20 13:44:23 UTC 2020"
	// Output of "date -u"
	Time = ""
)

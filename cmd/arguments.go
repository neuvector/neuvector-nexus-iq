package cmd

import "github.com/spf13/pflag"

const addressConfigKey = "address"

const addressFlagName = "address"

func addressFlag(fs *pflag.FlagSet) {
	fs.StringVarP(&rootConfig.Address, addressFlagName, "a", "", "Address of the webhook server")
	_ = v.BindPFlag(addressConfigKey, fs.Lookup(addressFlagName))
}

const portConfigKey = "port"

const portFlagName = "port"

func portFlag(fs *pflag.FlagSet) {
	fs.Int16VarP(&rootConfig.Port, portFlagName, "p", 5080, "Port of the webhook server")
	_ = v.BindPFlag(portConfigKey, fs.Lookup(portFlagName))
}

const neuvectorEndpointConfigKey = "neuvector.endpoint"

const neuvectorEndpointFlagName = "nv-endpoint"

func neuvectorEndpointFlag(fs *pflag.FlagSet) {
	fs.StringVar(&rootConfig.NeuVectorConfig.Endpoint, neuvectorEndpointFlagName, "", "Endpoint of the NeuVector Controller REST API (example: https://127.0.0.1:10443)")
	_ = v.BindPFlag(neuvectorEndpointConfigKey, fs.Lookup(neuvectorEndpointFlagName))
}

const neuvectorUsernameConfigKey = "neuvector.username"

const neuvectorUsernameFlagName = "nv-username"

func neuvectorUsernameFlag(fs *pflag.FlagSet) {
	fs.StringVar(&rootConfig.NeuVectorConfig.Username, neuvectorUsernameFlagName, "", "Username of the NeuVector Controller")
	_ = v.BindPFlag(neuvectorUsernameConfigKey, fs.Lookup(neuvectorUsernameFlagName))
}

const neuvectorPasswordConfigKey = "neuvector.password"

const neuvectorPasswordFlagName = "nv-password"

func neuvectorPasswordFlag(fs *pflag.FlagSet) {
	fs.StringVar(&rootConfig.NeuVectorConfig.Password, neuvectorPasswordFlagName, "", "Password of the NeuVector Controller")
	_ = v.BindPFlag(neuvectorPasswordConfigKey, fs.Lookup(neuvectorPasswordFlagName))
}

const neuvectorInsecureConfigKey = "neuvector.insecure"

const neuvectorInsecureFlagName = "nv-insecure"

func neuvectorInsecureFlag(fs *pflag.FlagSet) {
	fs.BoolVar(&rootConfig.NeuVectorConfig.Insecure, neuvectorInsecureFlagName, false, "If set, TLS certificate verification is skipped for the NeuVector controller. This should be used in testing scenarios only.")
	_ = v.BindPFlag(neuvectorInsecureConfigKey, fs.Lookup(neuvectorInsecureFlagName))
}

const nexusEndpointConfigKey = "nexusiq.endpoint"

const nexusEndpointFlagName = "nx-endpoint"

func nexusEndpointFlag(fs *pflag.FlagSet) {
	fs.StringVar(&rootConfig.NexusIqConfig.Endpoint, nexusEndpointFlagName, "", "Endpoint of the Nexus IQ REST API (example: http://127.0.0.1:8070)")
	_ = v.BindPFlag(nexusEndpointConfigKey, fs.Lookup(nexusEndpointFlagName))
}

const nexusUsernameConfigKey = "nexusiq.username"

const nexusUsernameFlagName = "nx-username"

func nexusUsernameFlag(fs *pflag.FlagSet) {
	fs.StringVar(&rootConfig.NexusIqConfig.Username, nexusUsernameFlagName, "", "Username of Nexus IQ")
	_ = v.BindPFlag(nexusUsernameConfigKey, fs.Lookup(nexusUsernameFlagName))
}

const nexusPasswordConfigKey = "nexusiq.password"

const nexusPasswordFlagName = "nx-password"

func nexusPasswordFlag(fs *pflag.FlagSet) {
	fs.StringVar(&rootConfig.NexusIqConfig.Password, nexusPasswordFlagName, "", "Password of Nexus IQ")
	_ = v.BindPFlag(nexusPasswordConfigKey, fs.Lookup(nexusPasswordFlagName))
}

const nexusInsecureConfigKey = "nexusiq.insecure"

const nexusInsecureFlagName = "nx-insecure"

func nexusInsecureFlag(fs *pflag.FlagSet) {
	fs.BoolVar(&rootConfig.NexusIqConfig.Insecure, nexusInsecureFlagName, false, "If set, TLS certificate verification is skipped for NexusIQ. This should be used for testing scenarios only.")
	_ = v.BindPFlag(nexusInsecureConfigKey, fs.Lookup(nexusInsecureFlagName))
}

const nexusSourceConfigKey = "nexusiq.source"

const nexusSourceFlagName = "nx-source"

func nexusSourceFlag(fs *pflag.FlagSet) {
	fs.StringVar(&rootConfig.NexusIqConfig.Source, nexusSourceFlagName, "NeuVector", "Source of vulnerabilities in Nexus IQ")
	_ = v.BindPFlag(nexusSourceConfigKey, fs.Lookup(nexusSourceFlagName))
}

const nexusOrganizationNameConfigKey = "nexusiq.organization_name"

const nexusOrganizationNameFlagName = "nx-org"

func nexusOrganizationNameFlag(fs *pflag.FlagSet) {
	fs.StringVar(&rootConfig.NexusIqConfig.OrganizationName, nexusOrganizationNameFlagName, "", "Name of the Nexus IQ organization to which vulnerabilities are reported")
	_ = v.BindPFlag(nexusOrganizationNameConfigKey, fs.Lookup(nexusOrganizationNameFlagName))
}

const applicationNameLabelConfigKey = "nexusiq.app_name_label"

const applicationNameLabelFlagName = "nx-app-name-label"

func applicationNameLabelFlag(fs *pflag.FlagSet) {
	fs.StringVar(&rootConfig.NexusIqConfig.AppNameLabel, applicationNameLabelFlagName, "com.sonatype.nexus.iq.applicationName", "Key of the label from which the name of the Nexus IQ application is inferred. If not provided, the application name will be derived from the name of the image.")
	_ = v.BindPFlag(applicationNameLabelConfigKey, fs.Lookup(applicationNameLabelFlagName))
}

package cmd

import "github.com/spf13/pflag"

const addressFlagName = "address"

func addressFlag(fs *pflag.FlagSet) {
	fs.StringVarP(&rootConfig.Address, addressFlagName, "a", "", "Address of the webhook server")
	_ = v.BindPFlag("address", fs.Lookup(addressFlagName))
}

const portFlagName = "port"

func portFlag(fs *pflag.FlagSet) {
	fs.Int16VarP(&rootConfig.Port, portFlagName, "p", 5080, "Port of the webhook server")
	_ = v.BindPFlag("port", fs.Lookup(portFlagName))
}

const neuvectorEndpointFlagName = "nv-endpoint"

func neuvectorEndpointFlag(fs *pflag.FlagSet) {
	fs.StringVar(&rootConfig.NeuVectorConfig.Endpoint, neuvectorEndpointFlagName, "", "Endpoint of the NeuVector Controller REST API (example: https://127.0.0.1:10443)")
}

const neuvectorUsernameFlagName = "nv-username"

func neuvectorUsernameFlag(fs *pflag.FlagSet) {
	fs.StringVar(&rootConfig.NeuVectorConfig.Username, neuvectorUsernameFlagName, "", "Username of the NeuVector Controller")
}

const neuvectorPasswordFlagName = "nv-password"

func neuvectorPasswordFlag(fs *pflag.FlagSet) {
	fs.StringVar(&rootConfig.NeuVectorConfig.Password, neuvectorPasswordFlagName, "", "Password of the NeuVector Controller")
}

const neuvectorInsecureFlagName = "nv-insecure"

func neuvectorInsecureFlag(fs *pflag.FlagSet) {
	fs.BoolVar(&rootConfig.NeuVectorConfig.Insecure, neuvectorInsecureFlagName, false, "If set, TLS certificate verification is skipped for the NeuVector controller. This should be used in testing scenarios only.")
}

const nexusEndpointFlagName = "nx-endpoint"

func nexusEndpointFlag(fs *pflag.FlagSet) {
	fs.StringVar(&rootConfig.NexusIqConfig.Endpoint, nexusEndpointFlagName, "", "Endpoint of the Nexus IQ REST API (example: http://127.0.0.1:8070)")
}

const nexusUsernameFlagName = "nx-username"

func nexusUsernameFlag(fs *pflag.FlagSet) {
	fs.StringVar(&rootConfig.NexusIqConfig.Username, nexusUsernameFlagName, "", "Username of Nexus IQ")
}

const nexusPasswordFlagName = "nx-password"

func nexusPasswordFlag(fs *pflag.FlagSet) {
	fs.StringVar(&rootConfig.NexusIqConfig.Password, nexusPasswordFlagName, "", "Password of Nexus IQ")
}

const nexusInsecureFlagName = "nx-insecure"

func nexusInsecureFlag(fs *pflag.FlagSet) {
	fs.BoolVar(&rootConfig.NexusIqConfig.Insecure, nexusInsecureFlagName, false, "If set, TLS certificate verification is skipped for NexusIQ. This should be used for testing scenarios only.")
}

const nexusSourceFlagName = "nx-source"

func nexusSourceFlag(fs *pflag.FlagSet) {
	fs.StringVar(&rootConfig.NexusIqConfig.Source, nexusSourceFlagName, "NeuVector", "Source of vulnerabilities in Nexus IQ")
}

const nexusOrganizationNameFlagName = "nx-org"

func nexusOrganizationNameFlag(fs *pflag.FlagSet) {
	fs.StringVar(&rootConfig.NexusIqConfig.OrganizationName, nexusOrganizationNameFlagName, "", "Name of the Nexus IQ organization to which vulnerabilities are reported")
}

const applicationNameLabelFlagName = "nx-app-name-label"

func applicationNameLabelFlag(fs *pflag.FlagSet) {
	fs.StringVar(&rootConfig.NexusIqConfig.AppNameLabel, applicationNameLabelFlagName, "com.sonatype.nexus.iq.applicationName", "Key of the label from which the name of the Nexus IQ application is inferred. If not provided, the application name will be derived from the name of the image.")
}

package cmd

import (
	"github.com/neuvector/neuvector-nexus-iq/integration"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start server",
	Run:   serveCmdExecute,
}

func init() {
	rootCmd.AddCommand(serveCmd)

	fs := serveCmd.Flags()

	// Integration arguments
	addressFlag(fs)
	portFlag(fs)

	// NeuVector arguments
	neuvectorEndpointFlag(fs)
	neuvectorUsernameFlag(fs)
	neuvectorPasswordFlag(fs)
	neuvectorInsecureFlag(fs)

	// Nexus arguments
	nexusEndpointFlag(fs)
	nexusUsernameFlag(fs)
	nexusPasswordFlag(fs)
	nexusInsecureFlag(fs)
	nexusSourceFlag(fs)
	nexusOrganizationNameFlag(fs)
	applicationNameLabelFlag(fs)
}

func serveCmdExecute(cmd *cobra.Command, args []string) {
	integrationServer := integration.NewServer(&rootConfig)

	integrationServer.Start()
}

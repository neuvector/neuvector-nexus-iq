package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version info",
	Run:   versionCmdExecute,
}

func versionCmdExecute(cmd *cobra.Command, args []string) {
	// FOLLOWUP integration version info from build process
	fmt.Println("NeuVector Nexus IQ integration v0.0.1")
}

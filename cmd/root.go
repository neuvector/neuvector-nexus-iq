package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nv-nx-iq",
	Short: "NeuVector Nexus IQ integration",
	Long:  "Integrate the NeuVector Container Security platform with Nexus IQ",
	Run:   rootCmdExecute,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file (default is ./config.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&rootConfig.Verbose, "verbose", "v", false, "verbose output")
}

func rootCmdExecute(cmd *cobra.Command, args []string) {

}

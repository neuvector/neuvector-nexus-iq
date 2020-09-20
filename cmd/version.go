package cmd

import (
	"fmt"
	"github.com/neuvector/neuvector-nexus-iq/build"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)

	versionCmd.Flags().BoolVar(&versionArgs.Short, "short", false, "Print version only")
}

type versionArgsType struct {
	Short bool
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version info",
	Run:   versionCmdExecute,
}

var versionArgs = &versionArgsType{}

func versionCmdExecute(cmd *cobra.Command, args []string) {
	if versionArgs.Short {
		fmt.Println(build.Version)

		return
	}

	fmt.Printf("NeuVector Nexus IQ integration %s\n", build.Version)

	if build.Commit != "" {
		fmt.Printf("Commit: %s\n", build.Commit)
	}

	if build.Time != "" {
		fmt.Printf("Build time: %s\n", build.Time)
	}
}

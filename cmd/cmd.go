package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	// Initialize Viper
	cobra.OnInitialize(initConfig)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

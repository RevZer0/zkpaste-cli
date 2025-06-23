package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the zkpaste tool version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("zkpaste ver. 0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

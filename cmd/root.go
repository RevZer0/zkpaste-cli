package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "zkpaste",
	Short: "ZK.paste CLI tool",
	Long: `zkpaste tool is a CLI interface to the zkpaste.com Zero-Knowlege pastebin service.
This tool allow to create, read and delete paste on zkpaste.com without using browser and javascript.
More at https://zkpaste.com`,
}

func Execute() error {
	return rootCmd.Execute()
}

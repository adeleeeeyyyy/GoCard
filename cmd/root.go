package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "repocard",
	Short: "CLI tool for generating GitHub repo card",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
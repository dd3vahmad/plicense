package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "plicense",
	Short: "Easily add licenses to your projects from the CLI",
	Long: `
    plicense (project license) is a tiny tool for quickly adding licenses to you project from the terminal (Command Line)
    without having to leave your codebase. It provides various other license related useful options and features too.
  `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to plicense")
	},
}

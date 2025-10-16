package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "plicense",
	Short: "Add open-source licenses to your project interactively",
	Long:  "plicense lets you preview and add licenses to your project using an interactive terminal UI.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'plicense add' to start the interactive license picker")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

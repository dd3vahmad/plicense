package cli

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "plicense",
	Short:   "Add open-source licenses to your project interactively",
	Long:    "plicense lets you preview and add licenses to your project using an interactive terminal UI.",
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s", color.HiBlueString("\nName: pLICENSE\n"))
		fmt.Printf("Version: v%s\n", color.HiYellowString(Version))
		fmt.Printf("Commit: %s\n", Commit)
		fmt.Printf("Built at: %s\n", Date)
		fmt.Printf("%s", color.HiCyanString("Run 'plicense --help' to see the list of commands\n"))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

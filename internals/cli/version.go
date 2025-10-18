package cli

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var shortOutput bool

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show pLICENSE version information",
	Run: func(cmd *cobra.Command, args []string) {
		if shortOutput {
			fmt.Printf("v%s\n", Version)
			return
		}
		fmt.Printf("v%s\n", color.HiYellowString(Version))
		fmt.Printf("Commit: %s\n", Commit)
		fmt.Printf("Built at: %s\n", Date)
	},
}

func init() {
	versionCmd.Flags().BoolVarP(&shortOutput, "short", "s", false, "print only the version number")
	rootCmd.AddCommand(versionCmd)
}

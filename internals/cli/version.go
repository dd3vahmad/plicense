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
	Short: "Show plicense version information",
	Run: func(cmd *cobra.Command, args []string) {
		if shortOutput {
			fmt.Println(Version)
			return
		}
		fmt.Printf("%s %s\n", color.HiCyanString("plicense"), color.HiYellowString(Version))
		fmt.Printf("commit: %s\n", Commit)
		fmt.Printf("built at: %s\n", Date)
	},
}

func init() {
	versionCmd.Flags().BoolVarP(&shortOutput, "short", "s", false, "print only the version number")
	rootCmd.AddCommand(versionCmd)
}

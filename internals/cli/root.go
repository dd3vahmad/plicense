package cli

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var showHelp bool

var rootCmd = &cobra.Command{
	Use:     "plicense",
	Short:   "Add open-source licenses to your project interactively",
	Long:    "plicense lets you preview, add, and manage licenses for your projects using an interactive terminal UI.",
	Version: Version,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\nName: %s", color.HiBlueString("pLICENSE\n"))
		fmt.Printf("Version: v%s\n", color.HiYellowString(Version))
		fmt.Printf("Commit: %s\n", color.HiMagentaString(Commit))
		fmt.Printf("Built at: %s\n", color.HiCyanString(Date))
		fmt.Printf("%s", color.HiCyanString("\nRun 'plicense --help' to see the list of commands.\n"))
	},
}

func init() {
	rootCmd.SetHelpTemplate(customHelpTemplate)
}

func Execute() {
	rootCmd.Flags().BoolVarP(&showHelp, "help", "h", false, "Show help details for plicense")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var customHelpTemplate = fmt.Sprintf(`
%s
Usage: {{.UseLine}}

{{.Long}}

%s
   add         Add open-source licenses interactively or by key
   delete      Delete the LICENSE file from the current directory
   view        View the contents of your project’s license

%s
   version     Show the installed version of plicense
   help        Show this help message

Examples:
   plicense add
   plicense add mit
   plicense delete
   plicense view
   plicense version

Run 'plicense help <command>' for details on a specific command.
`,
	color.HiCyanString("These are common pLICENSE commands used in various situations:"),
	color.HiYellowString("Manage licenses for your project"),
	color.HiYellowString("Get information about plicense"),
)

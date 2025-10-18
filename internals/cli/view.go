package cli

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dd3vahmad/plicense/internals/ui"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the license of your current project",
	RunE: func(cmd *cobra.Command, args []string) error {
		m, err := ui.NewLicenseModel()
		if err != nil {
			return fmt.Errorf("error loading licenses: %w", err)
		}

		p := tea.NewProgram(m)
		if _, err := p.Run(); err != nil {
			fmt.Println("Error running TUI:", err)
			os.Exit(1)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}

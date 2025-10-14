package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	ui "github.com/dd3vahmad/plicense/ui"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Interactively choose and add a license to your project",
	RunE: func(cmd *cobra.Command, args []string) error {
		m, err := ui.NewModel("licenses")
		if err != nil {
			return fmt.Errorf("Error loading licenses: %w", err)
		}

		p := tea.NewProgram(m)
		if _, err := p.Run(); err != nil {
			fmt.Println("Error running TUI:", err)
			os.Exit(1)
		}
		return nil
	},
}

func cmd() {
	rootCmd.AddCommand(addCmd)
}

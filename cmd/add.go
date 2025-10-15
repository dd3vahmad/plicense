package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dd3vahmad/plicense/fetch"
	ui "github.com/dd3vahmad/plicense/ui"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Interactively choose and add a license to your project",
	RunE: func(cmd *cobra.Command, args []string) error {
		m, err := ui.NewModel("licenses")
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
	// licenses, err := fetch.LicenseList()
	license, err := fetch.LicenseDetails("mit")
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Printf("LicenseList: %v\n", licenses)
	fmt.Printf("License: %v\n", license)

	rootCmd.AddCommand(addCmd)
}

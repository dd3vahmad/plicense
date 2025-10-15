package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dd3vahmad/plicense/fetch"
	"github.com/dd3vahmad/plicense/ui"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Interactively choose and add a license to your project",
	RunE: func(cmd *cobra.Command, args []string) error {
		// files, err := os.ReadDir(dir)
		// if err != nil {
		// 	return model{}, err
		// }

		// var licenses []License
		// for _, f := range files {
		// 	path := filepath.Join(dir, f.Name())
		// 	data, _ := os.ReadFile(path)
		// 	name := f.Name()
		// 	body := string(data)
		// 	licenses = append(licenses, License{Name: name, path: path, Body: body})
		// }

		licenses, err := fetch.LicenseList("licenses")
		if err != nil {
			fmt.Println(err)
			return nil
		}

		m, err := ui.NewModel(licenses)
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
	rootCmd.AddCommand(addCmd)
}

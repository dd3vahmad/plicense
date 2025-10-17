package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dd3vahmad/plicense/internals/entity"
	"github.com/dd3vahmad/plicense/internals/fetch"
	"github.com/dd3vahmad/plicense/internals/ui"
	"github.com/spf13/cobra"
)

func isPingReachable() bool {
	cmd := exec.Command("ping", "-c", "1", "www.google.com")
	err := cmd.Run()
	return err == nil
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Interactively choose and add a license to your project",
	RunE: func(cmd *cobra.Command, args []string) error {
		licensesFilePath, _ := fetch.LicensePath("")
		dir := filepath.Dir(licensesFilePath)

		var licenses []entity.License
		if isPingReachable() {
			lcs, err := fetch.LicenseList()
			if err != nil {
				fmt.Println(err)
				return nil
			}
			licenses = lcs
		} else {
			files, err := os.ReadDir(dir)
			if err != nil {
				return err
			}

			for _, f := range files {
				if f.Name() == "licenses.json" {
					continue
				}

				path := filepath.Join(dir, f.Name())
				data, _ := os.ReadFile(path)

				var license entity.License
				json.Unmarshal(data, &license)
				licenses = append(licenses, license)
			}
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

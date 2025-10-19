package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	name = ""
	dflt = ""
	show = false
)

type Config struct {
	Name    string `json:"name"`
	Default string `json:"default"`
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure plicense license details",
	Run: func(cmd *cobra.Command, args []string) {
		var config Config
		configDir, _ := os.UserConfigDir()
		path := filepath.Join(configDir, "plicense", "config.json")

		if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
			fmt.Printf("%s", color.HiRedString("Error creating pLICENSE config folder"))
			return
		}

		data, err := os.ReadFile(path)
		if err == nil && len(data) > 0 {
			if errr := json.Unmarshal(data, &config); errr != nil {
				fmt.Println("Error decoding config:", errr)
			}
		}

		if name != "" {
			config.Name = name
		}
		if dflt != "" {
			config.Default = dflt
		}

		file, err := os.Create(path)
		if err != nil {
			fmt.Printf("%s", color.HiRedString("Error creating pLICENSE config file"))
			return
		}
		defer file.Close()

		if err := json.NewEncoder(file).Encode(config); err != nil {
			fmt.Println("Error saving config:", err)
		}

		if show {
			fmt.Printf("\n%s\n\n", color.HiCyanString("Your pLICENSE config details"))
			fmt.Printf("Name: %s\n", config.Name)
			fmt.Printf("Default License: %s\n", config.Default)
		}
	},
}

func init() {
	configCmd.Flags().BoolVarP(&show, "show", "s", false, "To view your pLICENSE config details")
	configCmd.Flags().StringVar(&name, "name", "", "To configure user's default name")
	configCmd.Flags().StringVar(&dflt, "default", "", "To configure user's default license")
	rootCmd.AddCommand(configCmd)
}

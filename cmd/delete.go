package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete your project license directly from the CLI",
	RunE: func(cmd *cobra.Command, args []string) error {
		var file string
		if len(args) == 1 {
			file = args[0]
		} else if len(args) == 0 {
			file = "LICENSE"
		} else {
			return fmt.Errorf("too many arguments as license file name")
		}

		err := os.Remove(file)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

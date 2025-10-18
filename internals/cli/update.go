package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update pLICENSE to the latest version",
	Long: `Update pLICENSE to the latest version released on GitHub.

pLICENSE automatically detects your installation type and updates accordingly.

If installed via:
- Homebrew → updates via 'brew upgrade plicense'
- install.sh script → runs the installer again
- manually (GoReleaser binary) → downloads latest release from GitHub
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(color.HiCyanString("🔍 Detecting installation method..."))

		if isHomebrewInstall() {
			fmt.Println(color.HiYellowString("Detected Homebrew installation."))
			runCommand("brew", "upgrade", "plicense")
			return
		}

		if isSnapInstall() {
			fmt.Println(color.HiYellowString("Detected Snap installation."))
			runCommand("sudo", "snap", "refresh", "plicense")
			return
		}

		fmt.Println(color.HiYellowString("No package manager detected — updating from GitHub releases..."))
		updateFromGitHub()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func isHomebrewInstall() bool {
	path, _ := exec.LookPath("plicense")
	return strings.Contains(path, "Cellar")
}

func isSnapInstall() bool {
	path, _ := exec.LookPath("plicense")
	return strings.Contains(path, "/snap/")
}

func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Println(color.HiRedString("Failed to run command: %v", err))
	}
}

func updateFromGitHub() {
	apiURL := "https://api.github.com/repos/dd3vahmad/plicense/releases/latest"
	out, err := exec.Command("curl", "-s", apiURL).Output()
	if err != nil {
		fmt.Println(color.HiRedString("Failed to check GitHub releases: %v", err))
		return
	}

	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := json.Unmarshal(out, &release); err != nil || release.TagName == "" {
		fmt.Println(color.HiRedString("Failed to parse latest release."))
		return
	}

	fmt.Printf("📦 Latest version: %s\n", color.HiYellowString(release.TagName))
	fmt.Println(color.HiCyanString("Downloading and installing..."))

	downloadCmd := fmt.Sprintf(`curl -fsSL https://raw.githubusercontent.com/dd3vahmad/plicense/master/install.sh | bash`)
	cmd := exec.Command("bash", "-c", downloadCmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Println(color.HiRedString("Update failed: %v", err))
		return
	}

	fmt.Println(color.HiGreenString("pLICENSE updated successfully!"))
}

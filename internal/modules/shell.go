package modules

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var versionRe = regexp.MustCompile(`\b\d+\.\d+\.\d+\b`)

func extractVersion(output string) string {
	return versionRe.FindString(output)
}

func Shell() string {
	shellPath := os.Getenv("SHELL")
	shell := strings.ToLower(filepath.Base(shellPath))

	if shell == "" {
		return "unknown"
	}

	switch shell {

	case "bash":
		out, err := exec.Command("bash", "--version").Output()
		if err != nil {
			return "Bash"
		}

		if v := extractVersion(string(out)); v != "" {
			return fmt.Sprintf("Bash %s", v)
		}
		return "Bash"

	case "zsh":
		out, err := exec.Command("zsh", "--version").Output()
		if err != nil {
			return "Zsh"
		}

		if v := extractVersion(string(out)); v != "" {
			return fmt.Sprintf("Zsh %s", v)
		}
		return "Zsh"

	case "fish":
		out, err := exec.Command("fish", "--version").Output()
		if err != nil {
			return "Fish"
		}

		if v := extractVersion(string(out)); v != "" {
			return fmt.Sprintf("Fish %s", v)
		}
		return "Fish"

	default:
		return "unknown"
	}
}

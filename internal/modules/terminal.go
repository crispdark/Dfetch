package modules

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Terminal() string {
	if os.Getenv("ALACRITTY_SOCKET") != "" {
		out, err := exec.Command("alacritty", "--version").Output()
		if err != nil {
			return "Alacritty"
		}

		if v := extractVersion(string(out)); v != "" {
			return fmt.Sprintf("Alacritty %s", v)
		}
		return "Alacritty"
	}

	if os.Getenv("KITTY_PID") != "" {
		out, err := exec.Command("kitty", "--version").Output()
		if err != nil {
			return "Kitty"
		}

		if v := extractVersion(string(out)); v != "" {
			return fmt.Sprintf("Kitty %s", v)
		}
		return "Kitty"
	}

	if os.Getenv("TERM_PROGRAM") != "" {
		return strings.TrimSpace(os.Getenv("TERM_PROGRAM"))
	}

	return os.Getenv("TERM")
}

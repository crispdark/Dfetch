package modules

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func terminalWithVersion(cmd, name string, args ...string) string {
	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return name
	}

	if v := extractVersion(string(out)); v != "" {
		return fmt.Sprintf("%s %s", name, v)
	}

	return name
}

func Terminal() string {
	if os.Getenv("ALACRITTY_SOCKET") != "" {
		return terminalWithVersion("alacritty", "Alacritty", "--version")
	}

	if os.Getenv("GNOME_TERMINAL_SCREEN") != "" {
		return terminalWithVersion("gnome-terminal", "GNOME Terminal", "--version")
	}

	if os.Getenv("KITTY_PID") != "" {
		return terminalWithVersion("kitty", "Kitty", "--version")
	}

	if v := os.Getenv("KONSOLE_VERSION"); v != "" {
		return fmt.Sprintf("Konsole %s", v)
	}

	if termProg := os.Getenv("TERM_PROGRAM"); termProg != "" {
		switch strings.ToLower(strings.TrimSpace(termProg)) {
		case "vscode":
			return "VSCode"
		case "wezterm":
			return "WezTerm"
		default:
			return termProg
		}
	}

	term := os.Getenv("TERM")

	if strings.HasPrefix(term, "foot") {
		return terminalWithVersion("foot", "Foot", "--version")
	}

	switch term {
	case "":
		return "unknown"
	case "xterm":
		out, err := exec.Command("xterm", "-v").Output()
		if err != nil {
			return "XTerm"
		}

		var version int
		if _, err := fmt.Sscanf(strings.TrimSpace(string(out)), "XTerm(%d)", &version); err == nil {
			return fmt.Sprintf("XTerm %d", version)
		}

		return "XTerm"
	default:
		return term
	}
}

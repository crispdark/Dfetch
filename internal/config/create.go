package config

import (
	"dfetch/internal/getsysinfo"
	"os"
	"path/filepath"
	"strings"
)

func CreateConfigFile() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	appConfigDir := filepath.Join(configDir, "dfetch")

	// Create config directory if missing
	if err := os.MkdirAll(appConfigDir, 0700); err != nil {
		return err
	}

	configFile := filepath.Join(appConfigDir, "dfetch.conf")

	// Only create if missing
	if _, err := os.Stat(configFile); os.IsNotExist(err) {

		var config strings.Builder

		config.WriteString(
			"// Config file for Dfetch\n" +
				"// Lines starting with '//' will be ignored\n" +
				"// Default settings can be restored by deleting this file\n\n" +
				"// Change the color of the ASCII art. Supported colors are listed below.\n" +
				"color: default\n\n" +
				"// Supported colors:\n" +
				"// Black\n" +
				"// Red\n" +
				"// Green\n" +
				"// Yellow\n" +
				"// Blue\n" +
				"// Magenta\n" +
				"// Cyan\n" +
				"// White\n" +
				"// Bright_black / gray / grey\n" +
				"// Bright_red\n" +
				"// Bright_green\n" +
				"// Bright_yellow\n" +
				"// Bright_blue\n" +
				"// Bright_magenta\n" +
				"// Bright_cyan\n" +
				"// Bright_white\n\n" +
				"// Underneath a list of what information to show and in what order. Its recommended not to remove items but comment them out instead.\n\n" +
				"os\n" +
				"kernel\n" +
				"cpu\n" +
				"memory\n" +
				"localip\n" +
				"uptime\n" +
				"shell\n" +
				"//de\n",
		)

		// Add battery if available
		_, present := getsysinfo.Battery()
		if present != "unknown" {
			config.WriteString("battery\n")
		} else {
			config.WriteString("//battery\n")
		}

		err = os.WriteFile(configFile, []byte(config.String()), 0600)
		if err != nil {
			return err
		}
	}

	return nil
}

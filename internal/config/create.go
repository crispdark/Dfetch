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
				"// color: blue\n\n" +
				"// ASCII: ubuntu\n\n" +
				"// Info to fetch:\n" +
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

		file, err := os.Create(configFile)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.WriteString(config.String())
		if err != nil {
			return err
		}
	}

	return nil
}

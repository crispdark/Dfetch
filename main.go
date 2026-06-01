package main

import (
	"log"

	"dfetch/internal/config"
	"dfetch/internal/output"
	"dfetch/internal/sysinfo"
)

func main() {
	// Read or create the config file
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Collect necessary system info
	sys := sysinfo.CollectSystemInfo(cfg.EnabledModules)

	// Prepare the ASCII art
	asciiLines, asciiColor := output.LoadASCII(
		output.LogoFS,
		sys.ID,
		cfg.AsciiColor,
		cfg.AsciiSize,
	)

	accentColor := cfg.AccentColor
	if accentColor == "" || accentColor == "default" {
		accentColor = asciiColor
	}

	// Get the ANSI codes corresponding to the colors
	asciiColor = config.GetColorCode(asciiColor)
	accentColor = config.GetColorCode(accentColor)

	// Build the info lines
	infoLines := output.BuildInfoLines(
		sys,
		cfg.EnabledModules,
		accentColor,
	)

	// Put everything together and print it
	output.PrintOutput(asciiLines, infoLines, asciiColor)
}

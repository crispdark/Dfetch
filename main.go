package main

import (
	"dfetch/internal/config"
	"dfetch/internal/modules"
	"dfetch/internal/output"
	"log"
)

func main() {
	// Read or create the config file
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Collect necessary system info
	sys := modules.CollectSystemInfo(cfg.EnabledModules)

	// Prepare the ASCII art
	asciiLines, asciiColor := output.LoadASCII(
		output.LogoFS,
		sys.ID,
		cfg.AsciiColor,
		cfg.AsciiSize,
		cfg.CustomAscii,
	)

	if cfg.AccentColor == "" || cfg.AccentColor == "default" {
		cfg.AccentColor = asciiColor
	}

	// Get the ANSI codes corresponding to the colors
	asciiColor = config.GetColorCode(asciiColor)
	cfg.AccentColor = config.GetColorCode(cfg.AccentColor)

	// Build the info lines
	infoLines := output.BuildInfoLines(
		sys,
		cfg.EnabledModules,
		cfg.AccentColor,
	)

	// Put everything together and print it
	output.PrintOutput(asciiLines, infoLines, asciiColor)
}

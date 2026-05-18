package main

import (
	"Dfetch/internal/assets"
	"Dfetch/internal/config"
	"Dfetch/internal/customization"
	"Dfetch/internal/model"
	"fmt"
	"os"
)

func main() {
	lines, color := config.ReadConfig()

	noColor := false

	if len(os.Args) > 1 {
		if os.Args[1] == "--no-color" {
			noColor = true
			color = ""
		} else {
			fmt.Println("Invalid flag!")
			return
		}
	}

	sys := model.CollectSystemInfo()

	asciiLines, color := assets.LoadASCII(
		assets.LogoFS,
		sys.ID,
		color,
		noColor,
	)

	color = customization.GetColorCode(color)

	infoLines := buildInfoLines(sys, lines)

	printOutput(asciiLines, infoLines, color)
}

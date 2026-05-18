package main

import (
	"dfetch/internal/assets"
	"dfetch/internal/config"
	"dfetch/internal/customization"
	"dfetch/internal/model"
	"dfetch/internal/render"
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

	infoLines := render.BuildInfoLines(sys, lines)

	render.PrintOutput(asciiLines, infoLines, color)
}

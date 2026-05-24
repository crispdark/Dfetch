package main

import (
	"dfetch/internal/assets"
	"dfetch/internal/config"
	"dfetch/internal/customization"
	"dfetch/internal/model"
	"dfetch/internal/render"
)

func main() {
	lines, asciicolor := config.ReadConfig()

	sys := model.CollectSystemInfo()

	asciiLines, asciicolor := assets.LoadASCII(
		assets.LogoFS,
		sys.ID,
		asciicolor,
	)

	asciicolor = customization.GetColorCode(asciicolor)

	infoLines := render.BuildInfoLines(sys, lines)

	render.PrintOutput(asciiLines, infoLines, asciicolor)
}

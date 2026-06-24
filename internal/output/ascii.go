package output

import (
	"bufio"
	"embed"
	"fmt"
	"os"
	"strings"
)

func LoadASCII(fs embed.FS, distroID, customascii string) ([]string, string) {
	var scanner *bufio.Scanner
	var closer func()

	// custom file
	if customascii != "" && customascii != "default" {
		if f, err := os.Open(customascii); err == nil {
			scanner = bufio.NewScanner(f)
			closer = func() { f.Close() }
		} else {
			fmt.Printf("Error: Custom ascii path '%s' doesn't seem to exist.\n", customascii)
		}
	}

	// Build in ascii
	if scanner == nil {
		var file string

		file = fmt.Sprintf("logo/%s.txt", strings.ToLower(distroID))

		// Fallback to Linux logo
		if _, err := fs.Open(file); err != nil {
			file = "logo/linux.txt"
		}

		// Fallback to no ascii
		f, err := fs.Open(file)
		if err != nil {
			return nil, ""
		}

		scanner = bufio.NewScanner(f)
		closer = func() { f.Close() }
	}

	defer closer()

	var lines []string
	var accentColor string

	for scanner.Scan() {
		line := scanner.Text()

		// Take aside the accent_color
		if strings.HasPrefix(strings.ToLower(line), "accent_color:") {
			accentColor = strings.TrimSpace(strings.TrimPrefix(line, "accent_color:"))
			continue
		}
		lines = append(lines, line)
	}

	return lines, accentColor
}

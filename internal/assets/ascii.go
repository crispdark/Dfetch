package assets

import (
	"bufio"
	"embed"
	"fmt"
	"strings"
)

func LoadASCII(fs embed.FS, distroID, color string) ([]string, string) {
	file := fmt.Sprintf("logo/%s.txt", strings.ToLower(distroID))

	f, err := fs.Open(file)
	if err != nil {
		f, err = fs.Open("logo/linux.txt")
		if err != nil {
			return nil, color
		}
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "color:") {
			if color == "" {
				color = strings.TrimSpace(strings.TrimPrefix(line, "color:"))
			}
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, color
	}

	return lines, color
}

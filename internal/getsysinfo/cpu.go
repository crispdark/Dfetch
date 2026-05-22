// getsysinfo/getcpu.go
package getsysinfo

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func Cpu() string {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return "unknown"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var cpu string

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "model name	: ") {
			cpu = strings.TrimPrefix(line, "model name	: ")

			cpu = cleanCPU(cpu)

			return cpu
		}
	}
	return "unknown"
}

// Cleans up the CPU name
func cleanCPU(cpu string) string {
	// Lower-noise replacements
	replacements := []string{
		"(R)", "",
		"(TM)", "",
		"CPU", "",
		"Processor", "",
		"with Radeon Graphics", "",
	}

	for i := 0; i < len(replacements); i += 2 {
		cpu = strings.ReplaceAll(cpu, replacements[i], replacements[i+1])
	}

	// Remove everything after " w/"
	if strings.Contains(cpu, " w/") {
		cpu = strings.Split(cpu, " w/")[0]
	}

	// Remove clock speed
	reClock := regexp.MustCompile(`@\s*[0-9.]+\s*GHz`)
	cpu = reClock.ReplaceAllString(cpu, "")

	// Collapse multiple spaces
	reSpaces := regexp.MustCompile(`\s+`)
	cpu = reSpaces.ReplaceAllString(cpu, " ")

	return strings.TrimSpace(cpu)
}

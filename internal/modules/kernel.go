package modules

import (
	"os"
	"strings"
)

func Kernel() string {
	file, err := os.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return "unknown"
	}

	kernel := strings.TrimSpace(string(file))

	return kernel
}

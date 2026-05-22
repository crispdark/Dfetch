package getsysinfo

import (
	"os/exec"
	"strings"
)

func Shell() string {
	cmd := exec.Command(
		"bash",
		"-c",
		`shell=$(basename "$SHELL"); echo "$shell"`,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "unknown"
	}

	shellName := strings.TrimSpace(string(output))

	return string(shellName)
}

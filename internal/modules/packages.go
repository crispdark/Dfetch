package modules

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

func Packages() string {
	var cmd *exec.Cmd

	switch {
	case exists("dpkg-query"):
		cmd = exec.Command("dpkg-query", "-f", "${binary:Package}\n", "-W")

	case exists("rpm"):
		cmd = exec.Command("rpm", "-qa")

	case exists("pacman"):
		cmd = exec.Command("pacman", "-Qq")

	default:
		return "Unknown package manager"
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "unknown"
	}

	return strconv.Itoa(len(strings.Fields(out.String())))
}

func exists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

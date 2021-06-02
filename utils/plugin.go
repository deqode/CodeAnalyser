package utils

import "os/exec"

func CallPluginCommand(command string) *exec.Cmd {
	return exec.Command("sh", "-c", command)
}

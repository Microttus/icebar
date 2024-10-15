package app

import (
	"os/exec"
)

func Launch(execPath string) error {
	cmd := exec.Command(execPath)
	return cmd.Start()
}

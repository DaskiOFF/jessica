package command

import (
	"os"
	"os/exec"
)

func Execute(command string) error {
	cmd := exec.Command("sh", "-c", command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

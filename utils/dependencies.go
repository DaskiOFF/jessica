package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

func ExecCmd(name string, arg ...string) (string, error) {
	out, err := exec.Command(name, arg...).Output()
	return string(out), err
}

func InstallGemDependencies() error {
	cmd := exec.Command("sh", "-c", "sudo bundle install")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", out.String())
	return nil
}

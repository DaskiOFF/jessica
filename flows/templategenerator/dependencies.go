package templategenerator

import (
	"bytes"
	"fmt"
	"os/exec"
)

func installGemDependencies() error {
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

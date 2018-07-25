package templategenerator

import (
	"os/exec"
)

func execCmd(name string, arg ...string) (string, error) {
	out, err := exec.Command(name, arg...).Output()
	return string(out), err
}

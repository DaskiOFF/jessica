package utils

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
)

func ExecCmd(name string, arg ...string) (string, error) {
	out, err := exec.Command(name, arg...).Output()
	return string(out), err
}

func GitClone(url string, branch string, path string) (string, error) {
	destPath := filepath.Join(path, ".tmp")
	gitCloneCmd := "git clone " + url + " " + destPath
	if branch != "" {
		gitCloneCmd = gitCloneCmd + " --single-branch -b " + branch
	}
	moveCmd := "mv " + destPath + "/* " + path
	removeCmd := "rm -rf " + destPath

	cmd := exec.Command("sh", "-c", gitCloneCmd+";"+moveCmd+";"+removeCmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
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

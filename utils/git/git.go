package git

import (
	"path/filepath"

	"github.com/daskioff/jessica/utils/command"
)

func Clone(url string, branch string, path string) error {
	destPath := filepath.Join(path, ".tmpGit")

	gitCloneCmd := "git clone " + url + " " + destPath
	if branch != "" {
		gitCloneCmd = gitCloneCmd + " --single-branch -b " + branch
	}
	moveCmd := "mv " + destPath + "/* " + path
	removeCmd := "rm -rf " + destPath

	return command.Execute(gitCloneCmd + ";" + moveCmd + ";" + removeCmd)
}

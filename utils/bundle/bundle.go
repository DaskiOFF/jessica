package bundle

import "github.com/daskioff/jessica/utils/command"

func Install() error {
	return command.Execute("sudo bundle install")
}

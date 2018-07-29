package bundle

import (
	"github.com/daskioff/jessica/utils/command"
)

func Install() error {
	err := command.Execute("which bundle")
	if err != nil {
		err = command.Execute("sudo gem install bundler")
		if err != nil {
			return err
		}
	}

	return command.Execute("sudo bundle install")
}

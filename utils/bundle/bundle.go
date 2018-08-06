package bundle

import (
	"github.com/daskioff/jessica/utils/command"
)

func installBundleIfNeeded() error {
	err := command.Execute("which bundle")
	if err != nil {
		err = command.Execute("sudo gem install bundler")
		return err
	}

	return nil
}

func Install() error {
	err := installBundleIfNeeded()
	if err != nil {
		return err
	}

	return command.Execute("sudo bundle install")
}

func Update() error {
	err := installBundleIfNeeded()
	if err != nil {
		return err
	}

	return command.Execute("sudo bundle update")
}

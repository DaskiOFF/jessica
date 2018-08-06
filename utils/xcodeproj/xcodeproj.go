package xcodeproj

import "github.com/daskioff/jessica/utils/command"

func Install() error {
	err := command.Execute("which xcodeproj")
	if err != nil {
		err = command.Execute("sudo gem install xcodeproj")
		return err
	}

	return nil
}

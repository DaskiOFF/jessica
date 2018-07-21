package setup

import (
	"github.com/daskioff/jessica/utils"
)

func userName() string {
	return utils.AskQuestion("Your name (global): ", true)
}

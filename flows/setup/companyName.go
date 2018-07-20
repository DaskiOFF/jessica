package setup

import (
	"github.com/daskioff/jessica/utils"
)

func companyName() string {
	return utils.AskQuestion("Your company name (for project): ", true)
}

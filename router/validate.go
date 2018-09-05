package router

import (
	"errors"

	"github.com/daskioff/jessica/configs/models"
)

func (r *Router) validateConfigs() error {
	errorMessage := ""

	if err := r.globalConfig.Validate(); err != nil {
		errorMessage = err.Error() + "\n"
	}

	if err := r.projectConfig.Validate(); err != nil {
		errorMessage = errorMessage + err.Error() + "\n"
	}

	switch r.projectConfig.GetProjectType() {
	case models.ConfigProjectTypeIOS:
		if err := r.iosConfig.Validate(); err != nil {
			errorMessage = errorMessage + err.Error() + "\n"
		}
	case models.ConfigProjectTypeOther:
		if err := r.otherConfig.Validate(); err != nil {
			errorMessage = errorMessage + err.Error() + "\n"
		}
	}

	if errorMessage != "" {
		return errors.New(errorMessage + "Для начала необходимо настроить конфигурацию вызвав команду `jessica setup`")
	}

	return nil
}

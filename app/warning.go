package app

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

//----------------------------  Warner interface ----------------------------------------------
type Warner interface {
	Warning(data WarningData) error
}

type warning struct {
	logger *logrus.Logger
}


func (a *warning) Warning(data WarningData) error {
	if data.Type == WarningTypeSystem {
		a.logger.Error(data.Data)
	} else {
		a.logger.Error(fmt.Sprintf("task: %s, warning: %s", data.TaskName, data.Data))
	}

	return nil
}
//-----------------------------------------------------------------------------------------------
func NewDefaultWarner(logger *logrus.Logger) *warning {
	return &warning{logger: logger}
}


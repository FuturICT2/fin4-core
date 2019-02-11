package apperrors

import (
	"github.com/lytics/logrus"
)

//Critical reports a critical error
func Critical(errorID string, err error) {
	if err == nil {
		logrus.Error(errorID)
		return
	}
	logrus.WithFields(logrus.Fields{"e": err.Error()}).Error(errorID)
}

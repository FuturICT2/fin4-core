package apperrors

import (
	"github.com/FuturICT2/fin4-core/server/datatype"
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

//CriticalWithFields reports a critical error
func CriticalWithFields(errorID string, fields *datatype.Map) {
	if fields == nil {
		logrus.Error(errorID)
		return
	}
	logrus.WithFields(logrus.Fields(*fields)).Error(errorID)
}

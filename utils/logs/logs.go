package logs

import "github.com/sirupsen/logrus"

func Debug(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Debug(msg)
}
func Info(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Info(msg)
}

func Warning(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Warn(msg)
}

func Error(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Error(msg)
}

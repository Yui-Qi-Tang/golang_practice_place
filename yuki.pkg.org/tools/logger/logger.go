/*
    logrus sample
*/
package logger

import (
	"github.com/sirupsen/logrus"
)

func setLogInfo(msgField string, infoText string) {
    logrus.WithFields(logrus.Fields{
		"msg": msgField,
	}).Info(infoText)
}

func init() {
	setLogInfo("Initial phase of logger", "in tools/logger/logger.go")
}

func InfoLog(msg string, info string) {
	setLogInfo(msg, info)
}
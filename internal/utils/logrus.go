package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SendLogrusError(component string, action string, err error, message string) {
	logrus.WithFields(logrus.Fields{
		"component": component,
		"action":    action,
		"error":     err,
	}).Error(message)
}

func SendLogrusFatal(component string, action string, err error, message string) {
	logrus.WithFields(logrus.Fields{
		"component": component,
		"action":    action,
		"error":     err,
	}).Fatal(message)
}

func SendLogrusService(action string, c *gin.Context, err error, service string) {
	logrus.WithFields(logrus.Fields{
		"component": "service",
		"action":    action,
		"method":    c.Request.Method,
		"host":      c.Request.Host,
		"error":     err,
	}).Errorf("Error in service %s", service)
}

func SendLoggerHandlers(action string, err error, handler string) {
	logrus.WithFields(logrus.Fields{
		"component": "handler",
		"action":    action,
		"error":     err,
	}).Errorf("Error in handler %s", handler)
}

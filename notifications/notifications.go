package notifications

import (
	"it/apps/commons/logger"
)

func SendNotification(message string) (bool, error) {
	logger.Log.Info("sending notification", logger.Log.Args("message", message))
	return true, nil
}

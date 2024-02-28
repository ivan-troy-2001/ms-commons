package comments

import (
	"it/apps/commons/logger"
)

func PostComment(comment string) (bool, error) {
	logger.Log.Info("posting comment", logger.Log.Args("comment", comment))
	return true, nil
}

package utils

import (
	"go.uber.org/zap"
	"log"
)

//Logger todo: need to create/implement logger
func Logger(l ...interface{}) {
	log.Println(l...)
}

//Setting setting for logger
type Setting struct {
	Logger *zap.Logger
}

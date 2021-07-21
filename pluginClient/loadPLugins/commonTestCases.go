package loadPLugins

import (
	"code-analyser/utils"
	"go.uber.org/zap"
)

var logger, _ = zap.NewProduction()
var Set = utils.Setting{
	Logger: logger,
}

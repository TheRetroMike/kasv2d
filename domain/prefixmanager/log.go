package prefixmanager

import (
	"github.com/kasv2/kasv2d/infrastructure/logger"
	"github.com/kasv2/kasv2d/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)

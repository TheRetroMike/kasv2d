package rpcclient

import (
	"github.com/kasv2/kasv2d/infrastructure/logger"
	"github.com/kasv2/kasv2d/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)

package main

import (
	"github.com/kasv2/kasv2d/infrastructure/logger"
	"github.com/kasv2/kasv2d/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("IFLG")
	spawn      = panics.GoroutineWrapperFunc(log)
)

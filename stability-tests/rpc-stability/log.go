package main

import (
	"github.com/ixbasANT/gord/infrastructure/logger"
	"github.com/ixbasANT/gord/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("JSTT")
	spawn      = panics.GoroutineWrapperFunc(log)
)

package main

import (
	"github.com/ixbaseANT/gord/infrastructure/logger"
	"github.com/ixbaseANT/gord/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("IFLG")
	spawn      = panics.GoroutineWrapperFunc(log)
)

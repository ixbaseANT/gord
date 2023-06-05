package connmanager

import (
	"github.com/ixbasANT/gord/infrastructure/logger"
	"github.com/ixbasANT/gord/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)

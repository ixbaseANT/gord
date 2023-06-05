package prefixmanager

import (
	"github.com/ixbasANT/gord/infrastructure/logger"
	"github.com/ixbasANT/gord/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)

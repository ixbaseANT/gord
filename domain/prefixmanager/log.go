package prefixmanager

import (
	"github.com/ixbaseANT/gord/infrastructure/logger"
	"github.com/ixbaseANT/gord/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)

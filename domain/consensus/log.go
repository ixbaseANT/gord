package consensus

import (
	"github.com/ixbaseANT/gord/infrastructure/logger"
	"github.com/ixbaseANT/gord/util/panics"
)

var log = logger.RegisterSubSystem("BDAG")
var spawn = panics.GoroutineWrapperFunc(log)

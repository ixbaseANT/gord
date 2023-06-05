package grpcclient

import (
	"github.com/ixbasANT/gord/infrastructure/logger"
	"github.com/ixbasANT/gord/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)

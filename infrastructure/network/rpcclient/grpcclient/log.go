package grpcclient

import (
	"github.com/ixbaseANT/gord/infrastructure/logger"
	"github.com/ixbaseANT/gord/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.GoroutineWrapperFunc(log)

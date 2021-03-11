package service

import (
	"DTU/jobs/http_alleyes/biz"
	"DTU/jobs/http_alleyes/data"
	"github.com/google/wire"
)

var RecvAlleyesSet = wire.NewSet(data.NewAlleyesViolations, wire.Bind(new(biz.EventData), new(data.AlleyesViolations)), biz.NewAlleyesRecv, NewalleyesEventRecvData)

package service

import (
	"DTU/jobs/pulsar_tuya/biz"
	"DTU/jobs/pulsar_tuya/data"
	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(data.NewRespEc, data.NewTuyaRecv, wire.Bind(new(biz.EcMail), new(*data.RespEc)), wire.Bind(new(biz.TuyaPulsar), new(*data.TuyaRecv)), biz.NewTuyaPulsarData, biz.NewGreeterRepo, NewGreeterService)

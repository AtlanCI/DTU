package biz

import (
	"DTU/configs/conf_v1"
	"DTU/jobs/pulsar_tuya/data"
	"testing"
)

func TestGetEcEmail(t *testing.T) {
	config := conf_v1.NewConfig(conf_v1.ConfigTuyaEmailInterface("http://47.94.143.103:47001/gateway/getRestaurant"), conf_v1.ConfigTuyaUidPath("F:\\github\\uid.txt"))
	var a = &data.RespEc{}
	GetEcEmail(config.ConnConfig.TuyaConfig, a)
}

package biz

import (
	"DTU/configs/conf_v1"
	"DTU/jobs/pulsar_tuya/data"
	"github.com/tuya/tuya-cloud-sdk-go/api/common"
	config2 "github.com/tuya/tuya-cloud-sdk-go/config"
	"testing"
)

func TestStartTuyaPulsar(t *testing.T) {
	config := conf_v1.NewConfig(conf_v1.ConfigTuyaEmailInterface("http://47.94.143.103:47001/gateway/getRestaurant"), conf_v1.ConfigTuyaUidPath("F:\\github\\uid.txt"))
	var a = data.TuyaRecv{}
	var q = data.RespEc{}
	config2.SetEnv(common.URLCN, config.ConnConfig.TuyaConfig.TuyaIotCloudAccessId, config.ConnConfig.TuyaConfig.TuyaIotCloudAccessCRET)
	StartTuyaPulsar(config.ConnConfig.TuyaConfig, &a, &q)
}

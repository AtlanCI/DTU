package service

import (
	"DTU/configs/conf_v1"
	"DTU/jobs/pulsar_tuya/biz"
	"DTU/pkg"
	TuyaSDK "github.com/tuya/tuya-cloud-sdk-go/config"
)

type GreeterService struct {
	uc *biz.GreeterRepo
	uw *biz.TuyaPulsarData
}

func NewGreeterService(uc *biz.GreeterRepo, data *biz.TuyaPulsarData) *GreeterService {
	return &GreeterService{uc: uc, uw: data}
}

func StartPulsar(tuya *conf_v1.Config) {
	Service := initializeNewGreeterService()
	TuyaSDK.SetEnv("https://openapi.tuyacn.com", tuya.ConnConfig.TuyaConfig.TuyaIotCloudAccessId, tuya.ConnConfig.TuyaConfig.TuyaIotCloudAccessCRET)
	if !pkg.Exists(tuya.ConnConfig.TuyaConfig.TuyaUidCachePath) {
		biz.StartTuyaPulsar(tuya.ConnConfig.TuyaConfig, Service.uw.Repo(), Service.uc.Repo())
	}
}

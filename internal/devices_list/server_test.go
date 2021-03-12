package devices_list

import (
	"DTU/internal/devices_list/service"
	"DTU/pkg"
	"DTU/task/cronjob"
	TuyaSDK "github.com/tuya/tuya-cloud-sdk-go/config"
	"net/http"
	"testing"
)

func TestStartDeviceList(t *testing.T) {
	var CronManager *pkg.Crontab
	CronManager = pkg.NewCrontab()
	cronjob.NewAlleyesCron(CronManager)
	TuyaSDK.SetEnv("https://openapi.tuyacn.com", "5jh3wwnsa9wq5xda4hls", "bda847da6348420caae301def1f95c75")
	pkg.GlobleId = "ay1615257759386AMeJV"
	cronjob.NewTuyaCron(CronManager)
	CronManager.Start()
	router := pkg.NewHttpRouter()
	service.RegisterDeviceList(router)
	http.ListenAndServe(":8000", router)
}

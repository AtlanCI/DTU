package biz

import (
	"DTU/configs/conf_v1"
	"DTU/pkg"
	"encoding/json"
)

type DeviceEventSend struct {
	Path string
}

func (c *DeviceEventSend) DeviceEventCallBack(data interface{}) {
	if HealthData, ok := data.(*pkg.DeviceEventRecv); ok {
		Strdata, _ := json.Marshal(HealthData)
		c.SendEc(&Strdata)
	}
}
func (c *DeviceEventSend) SendEc(data *[]byte) {
	_, _, err := pkg.HttpJSON(c.Path, string(*data), 3, nil)
	if err != nil {
		panic(err)
	}
}
func NewDeviceEventSend(ec *conf_v1.Config) *DeviceEventSend {
	return &DeviceEventSend{ec.ConnConfig.EcConfig.HealthUpdateInterface}
}

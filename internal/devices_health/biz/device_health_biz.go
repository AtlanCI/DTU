package biz

import (
	"DTU/configs/conf_v1"
	"DTU/pkg"
	"encoding/json"
	"github.com/pkg/errors"
)

type DeviceHealthSend struct {
	Path string
}

func (c *DeviceHealthSend) HealthCallBack(data interface{}) {
	if HealthData, ok := data.(*pkg.DeviceList); ok {
		Strdata, _ := json.Marshal(HealthData)
		c.SendEc(&Strdata)
	}
}
func (c *DeviceHealthSend) SendEc(data *[]byte) {
	_, _, err := pkg.HttpJSON(c.Path, string(*data), 3, nil)
	if err != nil {
		pkg.LogSugar.Sugar().Warn(errors.Wrap(err, "send health to ec failed"))
	}
}
func NewDeviceHealthSend(ec *conf_v1.Config) *DeviceHealthSend {
	return &DeviceHealthSend{ec.ConnConfig.EcConfig.HealthUpdateInterface}
}

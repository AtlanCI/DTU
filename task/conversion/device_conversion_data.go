package conversion

import (
	"DTU/pkg"
	"strings"
	"time"
)

type DeviceCache []*pkg.DeviceCache

func (c DeviceCache) Exists(Code string) bool {
	for _, v := range c {
		if v.Sid == Code {
			return false
		}
	}
	return true
}

func (c DeviceCache) ConversionHealth() {
	for _, v := range c {
		var a = &pkg.DeviceHealth{
			DeviceCode:  v.Sid,
			StatusDate:  time.Now().Format("2006-01-02 15:04:05"),
			Manufactory: "tuya",
		}
		if v.Online != true {
			a.Status = 1
		}
		pkg.DeviceHealthchannel.Data = a
		pkg.DeviceHealthchannel.Notify()
	}
}

func (c DeviceCache) ConversionEvent() {

	for _, v := range c {
		// device not online and device is gateway be type
		// continue
		if v.Online == false || v.ProductName == "Zigbee智能网关" {
			continue
		}
		pptime := time.Now()
		var a = &pkg.DeviceViolationsEvent{
			EventId:          strings.ReplaceAll(pptime.Format("150405.9999"), ".", ""),
			EventType:        -1,
			EventDetail:      -1,
			EventDescription: v.ProductName + "异常",
			EventTime:        pptime.Format("2006-01-02 15:04:05"),
			DeviceCode:       v.Sid,
			DeviceTypeCode:   v.ProductId,
			Data:             v.Sensor,
			Manufactory:      "tuya",
		}
		pkg.DeviceEventChannel.Data = a
		pkg.DeviceEventChannel.Notify()
	}
}

func (c *DeviceCache) Append(cache *pkg.DeviceCache) {
	*c = append(*c, cache)
}

func (c *DeviceCache) TuyaCacheCallBack(data interface{}) {
	if Cache, ok := data.(*pkg.DeviceCache); ok {
		if !c.Exists(Cache.Sid) {
			return
		}
		c.Append(Cache)
	}
}

func NewDeviceCache() DeviceCache {
	return DeviceCache{}
}

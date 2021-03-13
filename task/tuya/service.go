package tuya

import (
	"DTU/pkg"
	"github.com/pkg/errors"
	"github.com/tuya/tuya-cloud-sdk-go/api/device"
	"github.com/tuya/tuya-cloud-sdk-go/api/user"
)

var cacheListId []string

func similarityId(id string) bool {
	for _, v := range cacheListId {
		if v == id {
			return true
		}
	}
	cacheListId = append(cacheListId, id)
	return false
}
func GetDeviceByUserId(uid string) {
	resp, err := user.GetDeviceListByUID(uid)
	if err != nil {
		pkg.LogSugar.Sugar().Warn(errors.Wrap(err, "Get device for device failed"))
		return
	}
	for _, v := range resp.Result {
		if recv, ok := v.(map[string]interface{}); ok {
			var a = &pkg.DeviceList{
				DeviceCode:     recv["id"].(string),
				DeviceTypeCode: recv["product_id"].(string),
				EventType:      -1,
				EventDetail:    -1,
			}
			if similarityId(a.DeviceCode) {
				return
			}
			pkg.DeviceListChannel.Data = a
			pkg.DeviceListChannel.Notify()
		}
	}
}

func conversion(d *device.GetDeviceResponse) *pkg.DeviceCache {
	a := make(map[string]interface{})
	switch d.Result.ProductName {
	case "温湿度":
		for _, v := range d.Result.Status {
			if v.Code == "va_temperature" {
				a["temperature"] = v.Value
			}
		}
		goto Loop
	case "智能插座":
		for _, v := range d.Result.Status {
			if v.Code == "cur_power" {
				a["load_power"] = v.Value
			}
		}
		goto Loop
	case "门磁":
		for _, v := range d.Result.Status {
			if v.Code == "doorcontact_state" {
				switch v.Value.(type) {
				case bool:
					flag := v.Value.(bool)
					if flag {
						a["status"] = "open"
						goto Loop
					}
					a["status"] = "close"
					goto Loop
				default:
					a["status"] = "close"
					goto Loop
				}
			}
		}
	}
Loop:
	return &pkg.DeviceCache{
		Sid:          d.Result.ID,
		Data:         "",
		Mode:         d.Result.ProductName,
		Sensor:       a,
		Manufacturer: "T",
		ProductId:    d.Result.ProductID,
		Online:       d.Result.Online,
		ProductName:  d.Result.ProductName,
	}
}

func GetDeviceInfo() {
	for _, v := range cacheListId {
		resp, err := device.GetDevice(v)
		if err != nil {
			panic(err)
		}
		devCache := conversion(resp)
		pkg.DeviceCacheChannel.Data = devCache
		pkg.DeviceCacheChannel.Notify()
	}
}

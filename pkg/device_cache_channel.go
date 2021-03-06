package pkg

import "sync"

type DeviceCache struct {
	Sid          string                 `json:"sid"`
	Data         string                 `json:"data"`
	Mode         string                 `json:"mode"`
	Manufacturer string                 `json:"-"`      // T == 涂鸦 M == 小米 A == Aqara
	Sensor       map[string]interface{} `json:"sensor"` //特殊字段 多功能设备时使用
	ProductId    string                 `json:"product_id,omitempty"`
}

type DeviceCacheRecv interface {
	TuyaCacheCallBack(data interface{})
}

type TuyaDeviceCacheSubject interface {
	AddListener(server DeviceCacheRecv)
	Notify()
}
type DeviceCacheInfoChannel struct {
	Observer sync.Map
	Data     *DeviceCache
}

func (i *DeviceCacheInfoChannel) AddListener(server DeviceCacheRecv) {
	i.Observer.Store(server, struct{}{})
}

func (i *DeviceCacheInfoChannel) Notify() {
	i.Observer.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(DeviceListRecv).TuyaListCallBack(i.Data)
		return true
	})
}

package pkg

import "sync"

type DeviceListRecv interface {
	TuyaListCallBack(data interface{})
}

type TuyaDeviceListSubject interface {
	AddListener(server DeviceListRecv)
	Notify()
}

type DeviceList struct {
	DeviceCode     string `json:"device_code"`
	DeviceTypeCode string `json:"device_type_code"`
	EventType      int    `json:"event_type"`
	EventDetail    int    `json:"event_detail"`
}

type DeviceListInfoChannel struct {
	Observer sync.Map
	Data     *DeviceList
}

func (i *DeviceListInfoChannel) AddListener(server DeviceListRecv) {
	i.Observer.Store(server, struct{}{})
}

func (i *DeviceListInfoChannel) Notify() {
	i.Observer.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(DeviceListRecv).TuyaListCallBack(i.Data)
		return true
	})
}

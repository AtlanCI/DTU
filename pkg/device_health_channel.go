package pkg

import "sync"

type DeviceHealthRecv interface {
	HealthCallBack(data interface{})
}

type DeviceListSubject interface {
	AddListener(server DeviceHealthRecv)
	Notify()
}

type DeviceHealth struct {
	DeviceCode  string `json:"device_code"`
	StatusDate  string `json:"status_date"`
	Manufactory string `json:"manufactory"` //故意拼写 错误 如要更改请通知Ec程序 负责人
	Status      int    `json:"status"`
}

type DeviceHealthInfoChannel struct {
	Observer sync.Map
	Data     *DeviceHealth
}

func (i *DeviceHealthInfoChannel) AddListener(server DeviceHealthRecv) {
	i.Observer.Store(server, struct{}{})
}

func (i *DeviceHealthInfoChannel) Notify() {
	i.Observer.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(DeviceListRecv).TuyaListCallBack(i.Data)
		return true
	})
}

package pkg

import "sync"

type DeviceViolationsEvent struct {
	EventId          string      `json:"event_id"`
	EventType        int         `json:"event_type"`
	EventDetail      int         `json:"event_detail"`
	EventDescription string      `json:"event_description"`
	EventTime        string      `json:"event_time"`
	EventImg         string      `json:"event_img"`
	EventVideo       string      `json:"event_video"`
	DeviceCode       string      `json:"device_code"`
	DeviceTypeCode   string      `json:"device_type_code"`
	Data             interface{} `json:"data"`
	Manufactory      string      `json:"manufactory"` //故意拼写 错误 如要更改请通知Ec程序 负责人
	ExtensionField   string      `json:"extension_field"`
}

type DeviceEventInfoChannel struct {
	Observer sync.Map
	Data     *DeviceViolationsEvent
}
type DeviceEventSubject interface {
	AddListener(server DeviceEventRecv)
	Notify()
}

type DeviceEventRecv interface {
	DeviceEventCallBack(data interface{})
}

func (i *DeviceEventInfoChannel) AddListener(server DeviceEventRecv) {
	i.Observer.Store(server, struct{}{})
}

func (i *DeviceEventInfoChannel) Notify() {
	i.Observer.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(DeviceEventRecv).DeviceEventCallBack(i.Data)
		return true
	})
}

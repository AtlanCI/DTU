package biz

import "DTU/pkg"

type EventData interface {
	Id() string
	Type() int
	Detail() int
	Description() string
	Time() string
	Img() string
	Video() string
	Uid() string
}

func NewDeviceEvent(Data EventData) {
	var a = &pkg.DeviceViolationsEvent{
		EventId:          Data.Id(),
		EventType:        Data.Type(),
		EventDetail:      Data.Detail(),
		EventDescription: Data.Description(),
		EventTime:        Data.Time(),
		EventImg:         Data.Img(),
		EventVideo:       Data.Video(),
		DeviceCode:       Data.Uid(),
		DeviceTypeCode:   "alleyes-ipc",
		Data:             nil,
		Manufactory:      "alleyes",
		ExtensionField:   "",
	}
	pkg.DeviceEventChannel.Data = a
	pkg.DeviceEventChannel.Notify()
}

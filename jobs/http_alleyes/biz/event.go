package biz

import (
	"DTU/pkg"
	"fmt"
)

type EventData interface {
	Id() string
	Type() int
	Detail() int
	Description() string
	Time() string
	Img() string
	Video() string
	Uid() string
	ApplyToEvent([]byte)
}

func (Data AlleyesRecv) ConversionEvent() {
	var a = &pkg.DeviceViolationsEvent{
		EventId:          Data.Event.Id(),
		EventType:        Data.Event.Type(),
		EventDetail:      Data.Event.Detail(),
		EventDescription: Data.Event.Description(),
		EventTime:        Data.Event.Time(),
		EventImg:         Data.Event.Img(),
		EventVideo:       Data.Event.Video(),
		DeviceCode:       Data.Event.Uid(),
		DeviceTypeCode:   "alleyes-ipc",
		Data:             nil,
		Manufactory:      "alleyes",
		ExtensionField:   "",
	}
	fmt.Println("转换后", a)
	pkg.DeviceEventChannel.Data = a
	pkg.DeviceEventChannel.Notify()
}

type AlleyesRecv struct {
	Event EventData
}

func NewAlleyesRecv(data EventData) *AlleyesRecv {
	return &AlleyesRecv{Event: data}
}

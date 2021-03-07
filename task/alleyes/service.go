package alleyes

import (
	"DTU/pkg"
	"encoding/json"
	"time"
)

var cacheIpcId []string

func similarityId(id string) bool {
	for _, v := range cacheIpcId {
		if v == id {
			return true
		}
	}
	cacheIpcId = append(cacheIpcId, id)
	return false
}
func conversionHealth(ipc IpcStatus) *pkg.DeviceHealth {
	return &pkg.DeviceHealth{
		DeviceCode:  ipc.EquipmentId,
		StatusDate:  time.Now().Format("2006-01-02 15:04:05"),
		Manufactory: "alleyes",
		Status:      ipc.IpcOnline,
	}
}
func conversionList(ipc IpcStatus) *pkg.DeviceList {
	return &pkg.DeviceList{
		DeviceCode:     ipc.EquipmentId,
		DeviceTypeCode: ipc.IpcName,
		EventType:      ipc.EventType,
		EventDetail:    IpcStatus{}.EventDetail,
	}
}
func GetIpcStatus(path string) {
	var recv AlleyesStatus
	_, body, err := pkg.HttpGET(path, nil, 3, nil)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(body, &recv); err != nil {
		panic(err)
	}
	for _, v := range recv.EquipmentStatus {
		// conversion ipc to be device Health object,and notify observer
		h := conversionHealth(v)
		pkg.DeviceHealthchannel.Data = h
		pkg.DeviceHealthchannel.Notify()

		// similarity id
		if similarityId(v.EquipmentId) {
			continue
		}

		// conversion ipc to be device list object,and notify observer
		a := conversionList(v)
		pkg.DeviceListChannel.Data = a
		pkg.DeviceListChannel.Notify()
	}
}

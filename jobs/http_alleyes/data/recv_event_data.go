package data

import "time"

// AlleyesViolations 视频事件类，定义视频事件详细信息
type AlleyesViolations struct {
	EventId      string    `json:"event_id"`
	EquipmentId  string    `json:"equipment_id"`
	OccurredTime time.Time `json:"occurred_time"`
	EventType    int       `json:"event_type"`
	EventDetail  int       `json:"event_detail"`
	IpcType      int       `json:"ipc_type"`
	IpcName      string    `json:"ipc_name"`
	EventDes     string    `json:"event_des"`
	EventImg     string    `json:"event_img"`
	EventVideo   string    `json:"event_video"`
} //接受白目发送的违规事件结构

func NewAlleyesViolations() AlleyesViolations {
	return AlleyesViolations{}
}

func (p AlleyesViolations) Id() string {
	return p.EventId
}

func (p AlleyesViolations) Type() int {
	return p.EventType
}

func (p AlleyesViolations) Detail() int {
	return p.EventDetail
}

func (p AlleyesViolations) Description() string {
	return p.EventDes
}

func (p AlleyesViolations) Time() string {
	ctime := time.FixedZone("CST", 8*3600)
	return p.OccurredTime.In(ctime).Format("2006-01-02 15:04:05")
}

func (p AlleyesViolations) Img() string {
	return p.EventImg
}

func (p AlleyesViolations) Video() string {
	return p.EventVideo
}

func (p AlleyesViolations) Uid() string {
	return p.EquipmentId
}

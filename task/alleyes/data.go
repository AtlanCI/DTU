package alleyes

type AlleyesStatus struct {
	Code            int
	EquipmentStatus []IpcStatus
	Msg             string
}

type IpcStatus struct {
	EquipmentId string `json:"Equipment_id"`
	IpcOnline   int    `json:"Ipc_status"`
	IpcName     string `json:"Ipc_name"`
	IpcId       string `json:"Ipc_id"`
	EventType   int    `json:"Event_type"`
	EventDetail int    `json:"Event_detail"`
}

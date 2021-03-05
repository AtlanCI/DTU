package data

import "encoding/json"

type TuyaRecv struct {
	BizCode    string  `json:"bizCode"`
	BizData    BizData `json:"bizData"`
	DevID      string  `json:"devId"`
	ProductKey string  `json:"productKey"`
	Ts         int64   `json:"ts"`
	UUID       string  `json:"uuid"`
}

type BizData struct {
	DevID   string `json:"devId"`
	UID     string `json:"uid"`
	OwnerID string `json:"ownerId"`
	UUID    string `json:"uuid"`
	Token   string `json:"token"`
}

func (p *TuyaRecv) Parse(q []byte) bool {
	err := json.Unmarshal(q, p)
	if err != nil {
		return false
	}
	return true
}

func (p *TuyaRecv) UidToString() string {
	return p.BizData.UID
}
func (p *TuyaRecv) ModeToString() string {
	return p.BizCode
}

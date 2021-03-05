package data

import (
	"encoding/json"
)

type RespEc struct {
	Code int        `json:"code"`
	Data RespEcData `json:"data"`
	Msg  string     `json:"msg"`
}

type RespEcData struct {
	Email        string `json:"email"`
	LocalId      string `json:"localId"`
	IsEntered    int    `json:"isEntered"`
	RestaurantId string `json:"restaurantId"`
	PrivateKey   string `json:"privateKey"`
	RestName     string `json:"restName"`
	RestCode     string `json:"restCode"`
}

//return merchant email address
func (p *RespEc) Mail(q []byte) string {
	err := json.Unmarshal(q, p)
	if err != nil {
		return ""
	}
	return p.Data.Email
}

func (p *RespEc) ToString() string {
	return p.Data.Email
}

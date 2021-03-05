package biz

import (
	"fmt"
	"github.com/tuya/tuya-cloud-sdk-go/api/common"
)

type GetUserInfoReq struct {
	Uid string
}

func (t *GetUserInfoReq) Method() string {
	return common.RequestGet
}

func (t *GetUserInfoReq) API() string {
	return fmt.Sprintf("/v1.0/users/%s/infos", t.Uid)
}

func GetUserInfo(uId string) (*GetUserInfoResponse, error) {
	a := &GetUserInfoReq{Uid: uId}
	fmt.Println(a)
	resp := &GetUserInfoResponse{}
	fmt.Println("hhhh")
	err := common.DoAPIRequest(a, resp)
	fmt.Println("wwwww")
	return resp, err
}

type GetUserInfoResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  struct {
		Avatar       string `json:"avatar"`
		Country_code string `json:"country_code"`
		Create_time  int64  `json:"create_time"`
		Email        string `json:"email"`
		Mobile       string `json:"mobile"`
		Nick_name    string `json:"nick_name"`
		Temp_unit    int    `json:"temp_unit"`
		Uid          string `json:"uid"`
		Update_time  int64  `json:"update_time"`
		UserName     string `json:"user_name"`
	} `json:"result"`
	//error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

package service

import (
	"DTU/configs/conf_v1"
	"DTU/internal/devices_list/biz"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func AlleyesRecvEventHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(res))
	recv := initNewalleyesEventRecvData()
	recv.BzData.Event.ApplyToEvent(res)
	recv.BzData.ConversionEvent()
}

func RegisterAlleyesEvent(route *httprouter.Router, config *conf_v1.Config) {
	route.POST(config.ConnConfig.AlleyesConfig.EventInfoRecvInterface, biz.DeviceListHandler)
}

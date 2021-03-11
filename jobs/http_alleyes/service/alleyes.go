package service

import (
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

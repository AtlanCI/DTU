package biz

import (
	"DTU/pkg"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var DevicelistResp *CacheBiz

func init() {
	DevicelistResp = initDeviceDataSet()
	pkg.DeviceListChannel.AddListener(DevicelistResp)
}

func DeviceListHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	respData, _ := json.Marshal(DevicelistResp)
	w.Write(respData)
}

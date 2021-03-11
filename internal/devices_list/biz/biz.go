package biz

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var DevicelistResp CacheBiz

func init() {
	initDeviceDataSet()
}

func DeviceListHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

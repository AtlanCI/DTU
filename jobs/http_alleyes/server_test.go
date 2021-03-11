package http_alleyes

import (
	"DTU/jobs/http_alleyes/service"
	"DTU/pkg"
	"net/http"
	"testing"
)

func TestStartAlleyesRecv(t *testing.T) {
	router := pkg.NewHttpRouter()
	router.POST("/v1/wntime/agent/violations/receive", service.AlleyesRecvEventHandler)
	http.ListenAndServe(":8000", router)
}

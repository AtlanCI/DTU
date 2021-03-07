package service

import "github.com/julienschmidt/httprouter"

func NewAlleyesEventRecv(route *httprouter.Router) {
	route.POST("")
}

package pkg

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.r.ServeHTTP(w, r)
}

func NewHttpRouter() *httprouter.Router {
	router := httprouter.New()
	return router
}

func SendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	_, _ = io.WriteString(w, resp)
}

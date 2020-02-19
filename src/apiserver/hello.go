package apiserver

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type HelloRouter struct {
	apiserver *APIServer
	muxRouter *mux.Router
}

// Constructor
func CreateHelloRouter(apiserver *APIServer) *HelloRouter {
	var router = &HelloRouter{
		apiserver: apiserver,
		muxRouter: apiserver.MuxRouter,
	}

	router.init()
	return router
}

// Private
func (router *HelloRouter) handleHello(res http.ResponseWriter, req *http.Request) {
	if _, err := io.WriteString(res, "Hello World"); err != nil {
		router.apiserver.Logger.Info(err)
	}
}

func (router *HelloRouter) init() {
	router.muxRouter.HandleFunc("/hello", router.handleHello)
}

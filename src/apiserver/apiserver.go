package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type APIServer struct {
	config    *APIConfig
	Logger    *logrus.Logger
	MuxRouter *mux.Router
}

func ConfigureServer(config *APIConfig) *APIServer {
	return &APIServer{
		config:    config,
		Logger:    logrus.New(),
		MuxRouter: mux.NewRouter(),
	}
}

func (configuredServer *APIServer) StartServer() error {
	if err := configuredServer.configureLogger(); err != nil {
		return err
	}
	configuredServer.configureRouter()
	configuredServer.Logger.Info("Started GO API server...")

	return http.ListenAndServe(configuredServer.config.Port, configuredServer.MuxRouter)
}

func (configuredServer *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(configuredServer.config.LogLevel)

	if err != nil {
		return err
	}

	configuredServer.Logger.SetLevel(level)
	return nil
}

func (configuredServer *APIServer) configureRouter() {
	CreateHelloRouter(configuredServer)
}

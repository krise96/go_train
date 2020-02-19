package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gostartup/src/routers"
	"net/http"
)

type APIServer struct {
	config *APIConfig
	logger *logrus.Logger
	router *mux.Router
}

func ConfigureServer(config *APIConfig) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (configuredServer *APIServer) StartServer() error {
	if err := configuredServer.configureLogger(); err != nil {
		return err
	}
	configuredServer.configureRouter()
	configuredServer.logger.Info("Started GO API server...")

	return http.ListenAndServe(configuredServer.config.Port, configuredServer.router)
}

func (configuredServer *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(configuredServer.config.LogLevel)

	if err != nil {
		return err
	}

	configuredServer.logger.SetLevel(level)
	return nil
}

func (configuredServer *APIServer) configureRouter() {
	configuredServer.router.HandleFunc("/hello", routers.HandleHello())
}

package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gostartup/src/repositories"
	"net/http"
)

type APIServer struct {
	config    *APIConfig
	Logger    *logrus.Logger
	MuxRouter *mux.Router
	Store     *repositories.Store
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

	if err := configuredServer.configureStore(); err != nil {
		return err
	}

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

func (configuredServer *APIServer) configureStore() error {
	store := repositories.CreateStore(configuredServer.config.Store)

	if err := store.Open(); err != nil {
		return err
	}

	configuredServer.Store = store

	return nil
}

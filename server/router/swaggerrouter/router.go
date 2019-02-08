package swaggerrouter

import (
	"github.com/andriolisp/hangman/application/applicationservice"
	"github.com/andriolisp/hangman/infra"
	"github.com/gorilla/mux"
)

const (
	swaggerRouter = "/swagger"
)

//Register will add all the necessary routes
func Register(svc *mux.Router, base infra.BaseConfig, app *applicationservice.Service) {
	ctrl := controllerInstance(svc, base)

	svc.HandleFunc(swaggerRouter, ctrl.getSwagger)
}

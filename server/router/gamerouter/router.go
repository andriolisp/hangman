package gamerouter

import (
	"net/http"

	"github.com/andriolisp/hangman/application/applicationservice"
	"github.com/andriolisp/hangman/infra"
	"github.com/gorilla/mux"
)

const (
	gameRoute       = "/game"
	gameDetailRoute = "/game/{id}"
)

//Register will add all the necessary routes
func Register(svc *mux.Router, base infra.BaseConfig, app *applicationservice.Service) {
	ctrl := controllerInstance(svc, base, app)

	svc.HandleFunc(gameRoute, ctrl.postGame).Methods(http.MethodPost, http.MethodOptions)
	svc.HandleFunc(gameDetailRoute, ctrl.getGame).Methods(http.MethodGet, http.MethodOptions)
	svc.HandleFunc(gameDetailRoute, ctrl.putGame).Methods(http.MethodPut, http.MethodOptions)
}

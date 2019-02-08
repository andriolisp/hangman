package router

import (
	"github.com/andriolisp/hangman/application/applicationservice"
	"github.com/andriolisp/hangman/infra"
	"github.com/andriolisp/hangman/server/router/gamerouter"
	"github.com/andriolisp/hangman/server/router/swaggerrouter"
	"github.com/gorilla/mux"
)

//Register will set all the routes for the API
func Register(router *mux.Router, base infra.BaseConfig, app *applicationservice.Service) {
	gamerouter.Register(router, base, app)
	swaggerrouter.Register(router, base, app)
}

package infrautils

import (
	"net/http"

	"github.com/andriolisp/hangman/infra"
	"github.com/gorilla/mux"
)

//RouterInterface has the required methods to add a route
type RouterInterface interface {
	Handle(pattern string, handlerFunc http.HandlerFunc)
	Router() *mux.Router
	infra.BaseConfig
}

//Middleware has all the steps to intercept any request
type Middleware func(next http.Handler) http.Handler

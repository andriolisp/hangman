package server

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"sync"

	"github.com/andriolisp/hangman/application/applicationservice"
	"github.com/andriolisp/hangman/infra"
	"github.com/andriolisp/hangman/infra/infrautils"
	"github.com/andriolisp/hangman/server/router"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	once     sync.Once
	instance *Server
)

// Server has all the server's info
type Server struct {
	middleware infrautils.Middleware
	httpServer *mux.Router
	httpPrefix *mux.Router
	app        *applicationservice.Service
	infra.BaseConfig
}

//Router returns an instance of the mux.Router
func (s *Server) Router() *mux.Router {
	return s.httpPrefix
}

//Handle add a route an a action
func (s *Server) Handle(pattern string, handleFunc http.HandlerFunc) {
	s.httpServer.HandleFunc(pattern, handleFunc)
}

//App return the application instance
func (s *Server) App() *applicationservice.Service {
	return s.app
}

//RunStatic will run a static server in a different port
func (s *Server) RunStatic() error {
	dir, err := os.Getwd()
	if err != nil {
		s.Log().Error(err)
	}
	filePath := path.Join(dir, "client/build")
	fs := http.FileServer(http.Dir(filePath))
	http.Handle("/", fs)
	return http.ListenAndServe(":7000", nil)
}

// Run start the http server
func (s *Server) Run() error {
	s.httpServer.Use(s.Log().Middleware)
	router.Register(s.Router(), s.BaseConfig, s.app)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Access-Control-Allow-Origin"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	return http.ListenAndServe(fmt.Sprintf(":%v", s.Config().App.Port), handlers.CORS(headersOk, originsOk, methodsOk)(s.httpServer))
}

// Instance create a new instance of the Server struct
func Instance(base infra.BaseConfig) *Server {
	once.Do(func() {
		httpServer := mux.NewRouter().StrictSlash(true)
		httpPrefix := httpServer.PathPrefix(base.Config().App.Prefix).Subrouter()

		app, err := applicationservice.New(base)
		if err != nil {
			panic(err)
		}

		instance = &Server{
			nil,
			httpServer,
			httpPrefix,
			app,
			base,
		}
	})
	return instance
}

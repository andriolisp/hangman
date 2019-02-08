package swaggerrouter

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"

	"github.com/andriolisp/hangman/infra"
	"github.com/andriolisp/hangman/server/serverutil"
	"github.com/gorilla/mux"
)

var (
	once     sync.Once
	instance *controller
)

type controller struct {
	*mux.Router
	infra.BaseConfig
}

func controllerInstance(svc *mux.Router, base infra.BaseConfig) *controller {
	once.Do(func() {
		instance = &controller{
			svc,
			base,
		}
	})

	return instance
}

func (c *controller) getSwagger(w http.ResponseWriter, r *http.Request) {
	dir, err := os.Getwd()
	if err != nil {
		c.Log().Error(err)
	}
	filePath := path.Join(dir, "server/router/swaggerrouter/static/index.html")

	f, err := os.Open(filePath)
	if err != nil {
		c.Log().Error(err)
		serverutil.ResponseAPINotFoundError(w, r)
		return
	}
	defer f.Close()

	htmlFile, err := ioutil.ReadAll(f)
	if err != nil {
		c.Log().Error(err)
		serverutil.ResponseAPINotFoundError(w, r)
		return
	}

	serverutil.ResponseHTMLOK(w, r, htmlFile)
}

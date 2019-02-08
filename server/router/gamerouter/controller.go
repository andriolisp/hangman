package gamerouter

import (
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/andriolisp/hangman/application/applicationservice"
	"github.com/andriolisp/hangman/application/entity"
	"github.com/andriolisp/hangman/infra"
	"github.com/andriolisp/hangman/server/serverutil"
	"github.com/gorilla/mux"
	"github.com/pquerna/ffjson/ffjson"
)

var (
	once     sync.Once
	instance *controller
)

type controller struct {
	*mux.Router
	infra.BaseConfig
	app    *applicationservice.Service
	helper *helper
}

func controllerInstance(svc *mux.Router, base infra.BaseConfig, app *applicationservice.Service) *controller {
	once.Do(func() {
		instance = &controller{
			svc,
			base,
			app,
			helperInstance(base),
		}
	})

	return instance
}

func (c *controller) postGame(w http.ResponseWriter, r *http.Request) {
	game := new(entity.Game)

	bGame, err := ioutil.ReadAll(r.Body)
	if err != nil {
		serverutil.ResponseAPIError(w, r, err.Error())
		return
	}
	defer r.Body.Close()

	if err := ffjson.Unmarshal(bGame, game); err != nil {
		serverutil.ResponseAPIError(w, r, err.Error())
		return
	}

	game, err = c.app.Game.Add(game)
	if err != nil {
		serverutil.ResponseAPIError(w, r, err.Error())
		return
	}

	serverutil.ResponseAPIOK(w, r, c.helper.GameHelper(game))
}

func (c *controller) getGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		serverutil.ResponseAPINotFoundError(w, r)
		return
	}

	game, err := c.app.Game.Get(id)
	if err != nil {
		serverutil.ResponseAPIError(w, r, err.Error())
		return
	}

	serverutil.ResponseAPIOK(w, r, c.helper.GameHelper(game))
}

func (c *controller) putGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		serverutil.ResponseAPINotFoundError(w, r)
		return
	}

	game, err := c.app.Game.Get(id)
	if err != nil {
		serverutil.ResponseAPIError(w, r, err.Error())
		return
	}

	detail := entity.Detail{}
	bDetail, err := ioutil.ReadAll(r.Body)
	if err != nil {
		serverutil.ResponseAPIError(w, r, err.Error())
		return
	}
	defer r.Body.Close()

	if err := ffjson.Unmarshal(bDetail, &detail); err != nil {
		serverutil.ResponseAPIError(w, r, err.Error())
		return
	}

	detail.Letter = strings.TrimSpace(strings.ToUpper(detail.Letter))

	game, err = c.app.Game.Put(game, &detail)
	if err != nil {
		serverutil.ResponseAPIError(w, r, err.Error())
		return
	}

	serverutil.ResponseAPIOK(w, r, c.helper.GameHelper(game))
}

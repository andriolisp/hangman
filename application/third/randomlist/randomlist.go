package randomlist

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/andriolisp/hangman/application/third/randomlist/randomlistentity"
	"github.com/andriolisp/hangman/infra"
	"github.com/pquerna/ffjson/ffjson"
)

type RandomListService struct {
	base  infra.BaseConfig
	words []string
}

func NewRandomListService(base infra.BaseConfig) *RandomListService {
	svc := &RandomListService{
		base: base,
	}

	list, err := svc.GetWordList()
	if err != nil {
		base.Log().Error("Error getting words: ", err)
	}

	svc.words = list.Data

	return svc
}

func (s *RandomListService) GetRandomWord() string {
	if len(s.words) == 0 {
		s.GetWordList()
	}

	rand.Seed(time.Now().UnixNano())
	return strings.TrimSpace(strings.ToUpper(s.words[rand.Intn(len(s.words)-1)]))
}

func (s *RandomListService) GetWordList() (*randomlistentity.WordList, error) {
	res, err := http.Get("https://www.randomlists.com/data/words.json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bReturn, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response randomlistentity.WordList
	err = ffjson.Unmarshal(bReturn, &response)
	if err != nil {
		return nil, err
	}

	s.words = response.Data

	return &response, nil
}

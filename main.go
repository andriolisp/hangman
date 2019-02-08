package main

import (
	"fmt"
	"log"

	"github.com/andriolisp/hangman/infra"
	"github.com/andriolisp/hangman/server"
)

// our main function
func main() {
	base, err := infra.New()
	if err != nil {
		fmt.Println("Could not create the base file: ", err.Error())
	}

	base.Log().Printf("Initializing the %s application", base.Config().App.Name)
	base.Log().Println("Debug Mode:", base.Config().App.Debug)
	base.Log().Println("Startint application on port:", base.Config().App.Port)

	s := server.Instance(base)

	go func() {
		log.Fatal(s.Run())
	}()

	log.Fatal(s.RunStatic())
	select {}
}

package main

import (
	"log"

	"github.com/fedir/gmm02/buffaloapi/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}

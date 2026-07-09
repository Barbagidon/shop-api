package main

import (
	"log"

	"github.com/Barbagidon/shop-api/internal/app"
)

func main() {
	app := app.New()

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

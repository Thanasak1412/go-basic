package main

import (
	"database/sql"
	"github.com/Thanasak1412/go-social/internal/env"
	"github.com/Thanasak1412/go-social/internal/store"
	"log"
)

func main() {
	config := config{
		addr: env.GetString("API_ADDR", ":8080"),
	}

	db := sql.DB{}

	app := &application{
		config: config,
		store:  store.NewStorage(&db),
	}

	mux := app.mount()

	log.Fatal(app.run(mux))

}

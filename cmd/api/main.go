package main

import (
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := &application{
		config: config{addr: ":" + port, env: "development"},
	}

	if err := app.run(app.mount()); err != nil {
		log.Fatal(err)
	}
}

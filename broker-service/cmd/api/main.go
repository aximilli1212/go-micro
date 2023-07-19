package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

type Config struct{}

func main() {

	app := Config{}

	log.Printf("Starting broker service on Port: %s \n ", webPort)

	//define Http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	//Start Server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

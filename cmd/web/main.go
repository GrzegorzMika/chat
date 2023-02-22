package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

const webPort = ":8080"

func main() {
	mux := routes()

	log.WithField("port", webPort).Info("Starting web server")

	log.Fatal(http.ListenAndServe(webPort, mux))
}

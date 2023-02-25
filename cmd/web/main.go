package main

import (
	"chat/internal/handlers"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const webPort = ":8080"

func main() {
	log.SetReportCaller(true)
	mux := routes()

	log.Info("Starting channel listener")
	go handlers.ListenToWsChannel()

	log.WithField("port", webPort).Info("Starting web server")

	log.Fatal(http.ListenAndServe(webPort, mux))
}

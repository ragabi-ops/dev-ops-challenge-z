package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *App) start() {
	log.Println("Server is Running on port 5000")
	var it Item
	var hc HealthCheck
	a.r.HandleFunc("/", hc.GetHealthCheckJson).Methods("GET")
	a.r.HandleFunc("/health", hc.GetHealthCheckJson).Methods("GET")
	a.r.HandleFunc("/secret", it.GetItemJson).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", a.r))
}

func main() {
	app := App{
		r: mux.NewRouter(),
	}
	app.start()
}

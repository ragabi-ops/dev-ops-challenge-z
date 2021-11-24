package main

import (
	"github.com/gorilla/mux"
)

type Item struct {
	CodeName   string `json:"codeName"`
	SecretCode string `json:""secretCode"`
}

type HealthCheck struct {
	Status     string `json:"status"`
	DockerRepo string `json:"dockerRepo"`
}

type App struct {
	r *mux.Router
}

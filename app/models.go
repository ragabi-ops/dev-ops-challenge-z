package main

import (
	"github.com/gorilla/mux"
)

type ItemDAO interface {
	GetItemJson(item Item)
}

type Item struct {
	CodeName   string `json:"codeName"`
	SecretCode string `json:"secretCode"`
}

type HealthCheckDAO interface {
	ListTable()
	GetHealthCheckJson(hc HealthCheck)
}
type HealthCheck struct {
	Status     string `json:"status"`
	DockerRepo string `json:"dockerRepo"`
}

type App struct {
	r *mux.Router
}

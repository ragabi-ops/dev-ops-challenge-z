package main

import (
	"github.com/gorilla/mux"
)

type App struct {
	r *mux.Router
}

type GenResponse struct {
	Status    string `json:"status"`
	CodeName  string `json:"codeName"`
	Container string `json:"container"`
}

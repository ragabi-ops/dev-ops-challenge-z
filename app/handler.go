package main

import (
	"encoding/json"
	"net/http"
)

func (it *Item) GetItemJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	item, err := it.getItem()

	if err != nil {
		sendWebErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		sendWebErr(w, http.StatusInternalServerError, err.Error())
	}
}

func sendWebErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

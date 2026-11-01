package main

import (
	"encoding/json"
	"net/http"
)

// why do we set the writeheader whats the purpose of this impl in the http interface
func (app *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeJSON(w http.ResponseWriter, val interface{}) {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(val)
	if err != nil {
		log.Warnf("error encode json: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(b); err != nil {
		log.Warnf("error flush response to browser: %s", err)
	}
}

func write400(w http.ResponseWriter, err error) {
	log.Warnf("invalid request: %s", err)
	w.WriteHeader(http.StatusBadRequest)
	// TODO: might return a json message if request is from js
	fmt.Fprintf(w, "invalid request: %s", err)
}

func write500(w http.ResponseWriter, err error) {
	log.Warnf("internal error: %s", err)
	w.WriteHeader(http.StatusInternalServerError)
	// TODO: might return a json message if request is from js
	fmt.Fprintf(w, "internal error: %s", err)
}

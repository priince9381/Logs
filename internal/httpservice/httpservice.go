package httpservice

import (
	"assisment/internal/models"
	"assisment/pkg/buffer"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func RouteHanler(router *mux.Router, log *buffer.LogBuffer) {
	router.HandleFunc("/", HandlePingRequest).Methods("GET")

	router.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		HandleLogRequest(w, r, log)
	}).Methods("POST")
}

func HandlePingRequest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome To Log Saver"))
	return
}

func HandleLogRequest(w http.ResponseWriter, r *http.Request, b *buffer.LogBuffer) {
	// Parse the log data from the request body
	var log models.Log
	err := json.NewDecoder(r.Body).Decode(&log)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid log data"))
		return
	}

	// Store the log in the database
	b.AddToBuffer(&log)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to store log in the database"))
		return
	}
	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Log received and stored in the database with id " + strconv.Itoa(log.GetId())))
}

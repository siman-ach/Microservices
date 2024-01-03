package resources

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var logger *log.Logger

// Initialize sets up the logger for the CoolTown microservice
func Initialize() {
	logger = log.New(os.Stdout, "", log.LstdFlags)
}

// GetRouter returns the router with the endpoint "/cooltown" configured
func GetRouter() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/cooltown", handleRequest).Methods("POST")
	return router
}

// handleRequest processes the incoming POST request and returns the audio for the song snippet
func handleRequest(w http.ResponseWriter, r *http.Request) {
	input := map[string]interface{}{}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Printf("Error decoding request body: %v", err)
		return
	}

	// Check for the presence of the "Audio" or "audio" key in the JSON object
	audio, ok := input["Audio"]
	if !ok || audio == "" {
		audio, ok = input["audio"]
		if !ok || audio == "" {
			w.WriteHeader(http.StatusBadRequest)
			logger.Printf("Error: Audio field missing or empty")
			return
		}
	}

	// Prepare the request body for the search service
	searchRequestBody, err := json.Marshal(map[string]interface{}{"Audio": audio})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Printf("Error marshaling search request body: %v", err)
		return
	}

	// Send a POST request to the search service on port 3001
	searchRes, err := http.Post("http://127.0.0.1:3001/search", "application/json", bytes.NewBuffer(searchRequestBody))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Printf("Failed HTTP request to search microservice: %v", err)
		return
	}
	defer searchRes.Body.Close()

	// Check the status code from the search service
	if searchRes.StatusCode != http.StatusOK {
		w.WriteHeader(searchRes.StatusCode)
		logger.Printf("Unexpected status code from search service: %v", searchRes.Status)
		return
	}
	// Decode the response body from the search service
	searchBody := map[string]interface{}{}
	err = json.NewDecoder(searchRes.Body).Decode(&searchBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Printf("Error decoding search service response: %v", err)
		return
	}
	// Extract the track ID from the search service response and prepare the track URL
	trackID, ok := searchBody["Id"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Printf("Error: Id field missing in search service response")
		return
	}
	trackURL := "http://127.0.0.1:3000/tracks/" + strings.Replace(trackID.(string), " ", "+", -1)
	// Send a GET request to the tracks service to fetch the track audio
	tracksRes, err := http.Get(trackURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Printf("Error making request to tracks service: %v", err)
		return
	}
	defer tracksRes.Body.Close()

	// Decode the response body from the tracks service
	tracksBody := map[string]interface{}{}
	err = json.NewDecoder(tracksRes.Body).Decode(&tracksBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Printf("Error decoding tracks service response: %v", err)
		return
	}

	if tracksRes.StatusCode != http.StatusOK {
		w.WriteHeader(tracksRes.StatusCode)
		logger.Printf("Unexpected status code from tracks service: %v", tracksRes.Status)
		return
	}

	// Extract the track audio from the tracks service response
	trackAudio, ok := tracksBody["Audio"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Printf("Error: Audio field missing in tracks service response")
		return
	}
	// Prepare the response JSON object with the track audio
	response := map[string]interface{}{"Audio": trackAudio}

	// Encode the response JSON object and write it to the response writer
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Printf("Error encoding response: %v", err)
		return
	}

	// Set the response status code to 200 OK
	w.WriteHeader(http.StatusOK)
}

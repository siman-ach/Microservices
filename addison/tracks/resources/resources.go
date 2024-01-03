package resources

import (
	"encoding/json"
	"net/http"
	"tracks/repository"

	"github.com/gorilla/mux"
)

// updateTrack updates or inserts a music track to the repository
func updateTrack(w http.ResponseWriter, r *http.Request) {
	// Parse the track ID from the URL parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Decode the track from the request body
	var t repository.Track
	if err := json.NewDecoder(r.Body).Decode(&t); err == nil {

		// Check if the track ID in the request body matches the URL parameter
		if id == t.Id {
			// Update the existing track or insert a new track if it doesn't exist
			if n := repository.Update(t); n > 0 {
				w.WriteHeader(204) /* No Content */
			} else if n := repository.Insert(t); n > 0 {
				w.WriteHeader(201) /* Created */
			} else {
				w.WriteHeader(500) /* Internal Server Error */
			}
		} else {
			w.WriteHeader(400) /* Bad Request */
		}
	} else {
		w.WriteHeader(400) /* Bad Request */
	}
}

// readTrack retrieves a music track from the repository by ID
func readTrack(w http.ResponseWriter, r *http.Request) {

	// Parse the track ID from the URL parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Read the track from the repository
	if t, n := repository.Read(id); n > 0 {
		// If the track exists, return it as a JSON response with status code 200 OK
		w.WriteHeader(200) // OK
		json.NewEncoder(w).Encode(t)
	} else if n == 0 {
		// If the track doesn't exist, return a 404 Not Found response
		w.WriteHeader(404) // Not Found
	} else {
		// If there was an error reading the track from the repository, return a 500 Internal Server Error response
		w.WriteHeader(500) // Internal Server Error
	}
}

// tracksList function retrieves a list of all music tracks from the repository
func tracksList(w http.ResponseWriter, r *http.Request) {
	// Get the list of tracks from the repository
	tracks := repository.ListTracks()
	if tracks != nil {
		// If the list of tracks is not empty, return it as a JSON response with status code 200 OK
		w.WriteHeader(200) /* OK */
		json.NewEncoder(w).Encode(tracks)
	} else {
		// If there was an error retrieving the list of tracks from the repository, return a 500 Internal Server Error response
		w.WriteHeader(500) /* Internal Server Error */
	}
}

// deleteTrack function deletes a music track from the repository by ID
func deleteTrack(w http.ResponseWriter, r *http.Request) {
	// Parse the track ID from the URL parameters
	vars := mux.Vars(r)
	id := vars["id"]
	// Delete the track from the repository
	n := repository.Delete(id)

	if n > 0 {
		// If the track was successfully deleted, return a 204 No Content response
		w.WriteHeader(204) // No Content
	} else if n == 0 {
		// If the track doesn't exist, return a 404 Not Found response
		w.WriteHeader(404) // Not Found
	} else {
		// If there was an error deleting the track from the repository, return 500 Internal Server Error response
		w.WriteHeader(500) // Internal Server Error
	}
}

// Router returns a new router instance for the resources package
func Router() http.Handler {
	r := mux.NewRouter()
	// Define routes and their associated handlers

	/* Store - Handle PUT requests for updating a track by ID*/
	r.HandleFunc("/tracks/{id}", updateTrack).Methods("PUT")
	/* Document - Handle GET requests for retrieving a track by ID */
	r.HandleFunc("/tracks/{id}", readTrack).Methods("GET")
	/* List - Handle GET requests for retrieving a list of all tracks*/
	r.HandleFunc("/tracks", tracksList).Methods("GET")
	/* Delete - Handle DELETE requests for deleting a track by ID*/
	r.HandleFunc("/tracks/{id}", deleteTrack).Methods("DELETE")

	return r
}

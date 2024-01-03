package resources

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// APIResponse defines the response structure from audd.io API
type APIResponse struct {
	Status string    `json:"status"`
	Result APIResult `json:"result"`
}

// APIResult defines the result structure from audd.io API
type APIResult struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
}

var token string

const KEY = "810b9adf0d7817575580355d3c513bad" //replace with own API key here !!!

var logger *log.Logger

// Init initializes the resources package
// For the purposes of this assignment, the API key will be hardcoded
func Init() int {
	token = KEY
	logger = log.New(os.Stdout, "", 0)
	return 0
}

// Router returns a new router instance for the resources package
func Router() http.Handler {
	router := mux.NewRouter()

	// searching URL
	router.HandleFunc("/search", search).Methods("POST")

	return router
}

// search sends an HTTP POST request to audd.io API to recognize the given audio
// fragment and returns the recognized track title as JSON
func search(w http.ResponseWriter, r *http.Request) {

	// parse the request body
	requestBody := map[string]interface{}{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "failed to decode request body: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// get the audio fragment
	audio, ok := requestBody["Audio"]
	if !ok {
		http.Error(w, "audio is missing", http.StatusBadRequest)
		return
	}

	// create a request for the audd.io API
	apiReqBody := map[string]interface{}{"api_token": token, "audio": audio}
	marshalledBody, err := json.Marshal(apiReqBody)
	if err != nil {
		http.Error(w, "failed to marshal API request body: "+err.Error(), http.StatusInternalServerError)
		return
	}

	apiSendData := bytes.NewBuffer(marshalledBody)
	apiRes, err := http.Post("https://api.audd.io/recognize", "application/json", apiSendData)
	if err != nil {
		http.Error(w, "http request to audd.io failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer apiRes.Body.Close()

	if apiRes.StatusCode != http.StatusOK {
		http.Error(w, "api request failed with code "+apiRes.Status, http.StatusInternalServerError)
		return
	}

	// read the response body
	apiResBodyMarshalled, err := io.ReadAll(apiRes.Body)
	if err != nil {
		http.Error(w, "failed to read http response body of api request to audd.io: "+err.Error(), http.StatusInternalServerError)
		return
	}

	apiResBody := APIResponse{}
	if err := json.Unmarshal(apiResBodyMarshalled, &apiResBody); err != nil {
		http.Error(w, "failed to unmarshal response from audd.io api response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if apiResBody.Status != "success" {
		log.Printf("API response error: %v", apiResBody.Status)
		log.Printf("API response body: %v", string(apiResBodyMarshalled))
		http.Error(w, "api response error: "+apiResBody.Status, http.StatusInternalServerError)
		return
	}
	// construct the response to the user
	userRes := map[string]interface{}{"Id": apiResBody.Result.Title}
	if err := json.NewEncoder(w).Encode(userRes); err != nil {
		http.Error(w, "failed to encode response to user: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

# Addison Microservices

This project contains three microservices - Tracks, Search, and CoolTown.  

Located in addison/tracks, addison/search and addison/cooltown respectively.

# Tracks Microservice

The Tracks microservice stores music tracks in an SQL database and listens on port 3000. It has the following features:

-   **Creating Tracks:** Allows music tracks to be created by sending a PUT request to `/tracks/id`. The request body should contain a JSON object with an `Id` property (string) and an `Audio` property (WAV file encoded in Base64). Successful responses are `201 Created` and `204 No Content`. Unsuccessful responses are `400 Bad Request` and `500 Internal Server Error`.
-   **Listing Tracks:** Allows all music tracks to be listed by sending a GET request to `/tracks`. A successful response is `200 OK` accompanied by a list of strings. An unsuccessful response is `500 Internal Server Error`.
-   **Reading Tracks:** Allows a music track to be read by sending a GET request to `/tracks/id`. A successful response is `200 OK` accompanied by a JSON object with an `Id` property (string) and an `Audio` property (WAV file encoded in Base64). Unsuccessful responses are `404 Not Found` and `500 Internal Server Error`.
-   **Deleting Tracks:** Allows a music track to be deleted by sending a DELETE request to `/tracks/id`. A successful response is `204 No Content`. Unsuccessful responses are `404 Not Found` and `500 Internal Server Error`

## How to Run Tracks

To run the Tracks microservice, navigate to the `addison/tracks` directory and run the following command:

`sh run_tracks.sh`

The above script is a simple helper script for initialising & running the tracks microservice. 

Alternatively you may use the following commands to run the tracks microservice: 

Navigate to the tracks folder:
`cd addison/tracks`
Remove existing go.mod file:
`rm -f go.mod`
Initialize the Go module:
`go mod init tracks`
Update dependencies:
`go mod tidy`
Finally, run:
`go run main.go`

## Search Microservice

The Search microservice provides music recognition and listens on port 3001. It has the following features:

-   **Recognising Tracks:** Allows music fragments to be recognised by sending a POST request to `/search`. The request body should contain a JSON object with an `Audio` property (WAV file encoded in Base64). A successful response is `200 OK` accompanied by a JSON object with an `Id` property (string). Unsuccessful responses are `400 Bad Request`, `404 Not Found`, and `500 Internal Server Error`.

### How to Run

To run the Search microservice, navigate to the `addison/search` directory and run the following command:

`sh run_search.sh`

Alternatively you may use the following commands to run the search microservice: 

Navigate to the tracks folder:
`cd addison/search`
Remove existing go.mod file:
`rm -f go.mod`
Initialize the Go module:
`go mod init search`
Update dependencies:
`go mod tidy`
Finally, run:
`go run main.go`

**Note:** The Search microservice uses an API key that is currently **hardcoded** in the `addison/search/resources/resources.go` file. You will need to replace it with your own API key.

It is readily located under `resources.go `:
const  KEY = "...."

## CoolTown Microservice

The CoolTown microservice provides device integration and listens on port 3002. It has the following features:

-   **Retrieving Tracks:** Allows music tracks to be retrieved using music fragments by sending a POST request to `/cooltown`. The request body should contain a JSON object with an `Audio` property (WAV file encoded in Base64). A successful response is `200 OK` accompanied by a JSON object with an `Audio` property (WAV file encoded in Base64). Unsuccessful responses are `400 Bad Request`, `404 Not Found`, and `500 Internal Server Error`.

### How to Run

To run the CoolTown microservice, navigate to the `addison/cooltown` directory and run the following command:
`sh run_cooltown.sh`

Alternatively you may use the following commands to run the CoolTown microservice: 

Navigate to the tracks folder:
`cd addison/cooltown`
Remove existing go.mod file:
`rm -f go.mod`
Initialize the Go module:
`go mod init search`
Update dependencies:
`go mod tidy`
Finally, run:
`go run main.go`

## 
There are also some test scripts provided for each microservice (`script1.sh`, `script2.sh`, `script3.sh`, `script4.sh`, `script5.sh`, `script6.sh`)

-   `script1.sh`: This script tests the `tracks` microservice by creating a new track using a PUT request to `localhost:3000/tracks/id` with a JSON payload containing the track ID and audio data.

-   `script2.sh`: This script tests the `tracks` microservice by listing all tracks using a GET request to `localhost:3000/tracks`.

-   `script3.sh`: This script tests the `tracks` microservice by reading a specific track using a GET request to `localhost:3000/tracks/id`.

-   `script4.sh`: This script tests the `tracks` microservice by deleting a specific track using a DELETE request to `localhost:3000/tracks/id`.

-   `script5.sh`: This script tests the `search` microservice by recognizing a music fragment using a POST request to `localhost:3001/search` with a JSON payload containing the audio data of the fragment.

-   `script6.sh`: This script tests the `cooltown` microservice by retrieving a music track using a POST request to `localhost:3002/cooltown` with a JSON payload containing the audio data of a music fragment.

package main

import (
	"log"
	"net/http"

	"cooltown/resources"
)

func main() {
	// Initialize the logger for the CoolTown microservice
	resources.Initialize()

	// Get the router with the endpoint "/cooltown" configured
	router := resources.GetRouter()

	// Start the CoolTown microservice and listen for incoming requests on port 3002
	log.Println("Starting CoolTown microservice on port 3002...")
	log.Fatal(http.ListenAndServe(":3002", router))
}

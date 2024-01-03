#!/bin/sh

#The following is a simple helper script for initialising & running the tracks microservice

# Navigate to the tracks folder
cd addison/tracks


# Remove existing go.mod file
rm -f go.mod

# Initialize the Go module
go mod init tracks

# Update dependencies
go mod tidy

# Run 
go run main.go
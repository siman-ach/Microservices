#!/bin/sh

#The following is a simple helper script for initialising & running the search microservice

# Navigate to the search folder
cd addison/search

# Remove existing go.mod file, if any
rm -f go.mod

# Initialize the Go module
go mod init search

# Update dependencies
go mod tidy

# Run 
go run main.go

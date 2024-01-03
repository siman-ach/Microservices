#!/bin/sh

#The following is a simple helper script for initializing & running the CoolTown microservice

# Navigate to the cooltown folder
cd addison/cooltown

# Remove existing go.mod file
rm -f go.mod

# Initialize the Go module
go mod init cooltown

# Update dependencies
go mod tidy

# Run 
go run main.go

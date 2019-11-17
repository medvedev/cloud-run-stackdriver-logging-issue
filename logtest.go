package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/logging"
)

func main() {
	loggingClient, err := stackdriverClient()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		os.Exit(1)
	}
	logger := loggingClient.Logger("default")

	log.Println("Before stackdriver logging")

	logger.StandardLogger(logging.Info).Println("Stackdriver log")

	if err = logger.Flush(); err != nil {
		log.Fatalf("Failed to flush client: %v", err)
	}

	if err = loggingClient.Close(); err != nil {
		log.Fatalf("Failed to close client: %v", err)
	}

	log.Println("After stackdriver logging")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func stackdriverClient() (client *logging.Client, err error) {
	var projectID string
	if projectID, err = metadata.ProjectID(); err == nil {
		client, err = logging.NewClient(context.Background(), projectID)
	}
	return
}

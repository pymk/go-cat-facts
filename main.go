// Package main provides a simple client for fetching cat facts from an API.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	url         = "https://meowfacts.herokuapp.com/"
	httpTimeOut = 10 * time.Second
)

// Response represents the structure of the API response.
type Response struct {
	Data []string `json:"data"`
}

// main is the entry point of the program.
// It fetches a cat fact and prints it, or prints an error if one occurs.
func main() {
	fact, err := getCatFact()
	if err != nil {
		fmt.Printf("Error getting cat fact: %v\n", err)
		return
	}

	fmt.Println(fact)
}

// getCatFact fetches a single cat fact from the API.
// It returns the fact as a string and any error encountered.
func getCatFact() (string, error) {
	// Create an HTTP client with a timeout.
	client := &http.Client{Timeout: httpTimeOut}

	// Make a GET request to the API.
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close() // Ensure the response body is closed after we're done.

	// Check if the status code indicates success.
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Decode the JSON response directly from the response body.
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", fmt.Errorf("failed to decode JSON: %w", err)
	}

	// Check if we received any facts.
	if len(response.Data) == 0 {
		return "", fmt.Errorf("no cat facts received")
	}

	// Return the first fact from the response.
	return response.Data[0], nil
}

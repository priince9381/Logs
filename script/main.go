package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
)

type Log struct {
	ID        int    `json:"id"`
	UnixTs    int64  `json:"unix_ts"`
	UserID    int    `json:"user_id"`
	EventName string `json:"event_name"`
}

func main() {
	host := "api:8000"
	timeout := time.Second * 5
	maxRetries := 10
	retryInterval := time.Second * 3
	// Ping the database to ensure a successful connection
	for i := 0; i < maxRetries; i++ {
		err := ping(host, timeout)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to pod. Retrying in %v...", retryInterval)
		time.Sleep(retryInterval)

	}
	// Start sending requests
	sendRequests(10000) // Specify the desired number of requests

	fmt.Println("Finished sending requests.")
}

func ping(host string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", host, timeout)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

func sendRequests(numRequests int) {
	client := &http.Client{}
	payload := createPayload()
	counter := 0

	for counter < numRequests {
		// Convert payload to JSON
		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Failed to marshal payload:", err)
			continue
		}

		// Send POST request to the /logs endpoint
		resp, err := client.Post("http://api:8000/logs", "application/json", bytes.NewBuffer(jsonPayload))
		if err != nil {
			fmt.Println("Request failed:", err)
			continue
		}

		// Check response status code
		if resp.StatusCode != http.StatusOK {
			fmt.Println("Request failed with status code:", resp.StatusCode)
			continue
		}

		// Increment the counter
		counter++

		// Sleep for 1 millisecond (to achieve 1K requests per second)
		time.Sleep(time.Millisecond)
	}
}

func createPayload() Log {
	// Create a dynamic payload with current timestamp and random ID
	return Log{
		ID:        generateRandomID(),
		UnixTs:    time.Now().Unix(),
		UserID:    generateUserID(),
		EventName: getRandomEvent(),
	}
}

func generateRandomID() int {
	// Generate a random ID between 1000 and 50000
	return 1000 + rand.Intn(50000)
}

func generateUserID() int {
	// Generate a random ID till 50000
	return rand.Intn(50000)
}

func getRandomEvent() string {
	random := rand.Intn(50000)
	if random%2 == 0 {
		return "login"
	}
	return "logout"
}

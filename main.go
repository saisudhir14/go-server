package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// Slice to store user data (using a mutex for concurrent access)
var userData []string
var mu sync.Mutex

// formHandler handles form submission
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Extract form values
	name := r.FormValue("name")
	password := r.FormValue("password")

	// Go routine to handle data storage asynchronously
	go storeUserData(name, password)

	// Send the response to the client
	fmt.Fprintf(w, "POST request Successful\n")
	fmt.Fprintf(w, "name = %s\n", name)
	fmt.Fprintf(w, "password = %s\n", password)
}

// storeUserData appends user data to a slice with mutex for concurrency control
func storeUserData(name string, password string) {
	// Using mutex to protect concurrent access to userData
	mu.Lock()
	defer mu.Unlock()

	// Simulate time-consuming task (e.g., saving to a database)
	time.Sleep(2 * time.Second)

	// Append the data
	userData = append(userData, fmt.Sprintf("name: %s, password: %s", name, password))
	fmt.Printf("User data stored: %v\n", userData)
}

// helloHandler handles the /hello endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Handle 404 error if the path is incorrect
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// Ensure that only GET method is accepted
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}

	// Respond with a hello message
	fmt.Fprintf(w, "hello!")
}

// backgroundTask starts a Go routine for a background task (example for channels)
func backgroundTask(done chan bool) {
	time.Sleep(3 * time.Second) // Simulate background work
	fmt.Println("Background task completed")
	done <- true
}

func main() {
	// Setting up the file server for serving static files
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Handling the form and hello endpoints
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// Channel for managing background tasks
	done := make(chan bool)

	// Starting a background Go routine
	go backgroundTask(done)

	// Wait for the background task to complete
	go func() {
		<-done
		fmt.Println("Received completion signal from background task")
	}()

	// Starting the HTTP server
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

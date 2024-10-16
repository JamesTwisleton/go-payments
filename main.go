package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

const creditScoreMin = 500
const creditScoreMax = 999
const port = 8080

type creditRating struct {
	CreditRating int `json:"credit_rating"`
}

func main() {

	// uncomment this line to trigger linting error in pre-commit hook/CI
	// log.Printf(fmt.Sprintf("Starting server on :%d", port))

	// uncomment this line to trigger build error in pre-commit hook/CI
	// notBuildable.Printf("Starting server on :%d", port)

	log.Printf("Starting server on :%d", port)
	handleRequests()
}

func handleRequests() {
	// * when http.HandlerFunc is called, it converts getCreditScore into a http.Handler
	// which implements ServeHTTP method.
	// * the HTTP server passes in w and r to the getCreditScore method
	// TODO: learn more about how HandlerFunc and similar work
	http.Handle("/creditscore", http.HandlerFunc(getCreditScore))

	// * nil here means DefaultServeMux is used, Go's default HTTP request multiplexer/router.
	// * if http.ListenAndServe encounters an error, it will call log.Fatal to log the error
	// and terminate the program
	// * Formatting port with : prefix here means that it will bind to all network interfaces,
	// I could provide "127.0.0.1:%d" here for example
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

// TODO: read up on pointers like used in * here
func getCreditScore(w http.ResponseWriter, r *http.Request) {

	// construct new variable of type creditRating
	creditRating := creditRating{
		CreditRating: (rand.Intn(creditScoreMax-creditScoreMin) + creditScoreMin),
	}

	// write the status code to the ResponseWriter
	w.WriteHeader(http.StatusOK)

	// Convert the creditRating variable to JSON and send it to the ResponseWriter `w`.
	// `json.NewEncoder(w).Encode(creditRating)` does this by encoding the `creditRating` struct as JSON and writing it to `w`,
	// which sends the response back to the client.
	//
	// Go functions often return an error type if something goes wrong. In this case, `Encode` will return an error if the
	// struct cannot be properly converted to JSON.
	//
	// The `if err := ...; err != nil` pattern declares and checks the error in a single step.
	// If `err` is not `nil`, it means something went wrong, and the code inside the `if` block will run.
	// In this case, we return an HTTP 500 (Internal Server Error) to the client if JSON encoding fails.
	if err := json.NewEncoder(w).Encode(creditRating); err != nil {

		// Send a 500 Internal Server Error response if the JSON encoding fails.
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

}

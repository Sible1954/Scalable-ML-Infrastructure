
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting Scalable ML Infrastructure Server...")

	http.HandleFunc("/predict", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Prediction endpoint: Model serving is active!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

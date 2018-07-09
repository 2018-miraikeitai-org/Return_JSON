package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Result string
}

type Responses []Response

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(Response{Result: "Hello net/http"})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

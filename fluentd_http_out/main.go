package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleRequest(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusMethodNotAllowed)
		res.Write([]byte("Method not allowed"))
		return
	}

	if req.Header.Get("Content-type") != "application/x-ndjson" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Only application/x-ndjson content is allowed"))
		return
	}

	decoder := json.NewDecoder(req.Body)
	for decoder.More() {
		var value interface{}
		if err := decoder.Decode(&value); err != nil {
			fmt.Errorf("parse error: %w\n", err)
		} else {
			fmt.Printf("value: %#v\n", value)
		}
	}

	fmt.Fprintf(res, "OK")
}

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

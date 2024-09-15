package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Fatal("faild to marshal json response: %v, Error: %v\n", payload, err)
		w.WriteHeader(500) // if an error happend while marshaling the payload we will return a 500 status code
	}
	w.Header().Add("Content-Type", "application/json") // response with json
	w.WriteHeader(200)
	w.Write(data)

}

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	response := Response{Message: "Hello World!"}
	json.NewEncoder(w).Encode(response)
}
func checkHandler(w http.ResponseWriter, r *http.Request){
	response := Response{Message: "cheking World!"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/check", checkHandler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}
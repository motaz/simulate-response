package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("simulate-response")
	fmt.Println("http://localhost:2025")
	http.HandleFunc("/", handleRequests)
	err := http.ListenAndServe(":2025", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

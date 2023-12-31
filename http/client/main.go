package main

import (
	"bytes"
	"net/http"
)

func main() {
	client := http.Client{}

	//
	_, err := client.Get("http://localhost:8080/")
	if err != nil {
		return
	}

	//
	request, err := http.NewRequest("GET", "http://localhost:8080/", bytes.NewReader([]byte("test")))
	if err != nil {
		return
	}
	_, err = client.Do(request)
	if err != nil {
		return
	}
}

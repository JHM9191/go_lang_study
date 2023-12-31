package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", GetData)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return
	}
}

func GetData(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Printf("Could not read body: %s\n", err)
		}
		defer request.Body.Close()

		if len(body) > 0 {
			fmt.Println("has req body")
		}

		writer.Write([]byte("data"))
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

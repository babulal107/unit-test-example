package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Simple http server example server request and response data
func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		// log `Hello word` message on server
		log.Println("Hello word")

		// read request body data
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "Oops invalid request", http.StatusBadRequest)
			return
		}

		// write response back to write handler
		fmt.Fprintf(writer, "Response data : %s", data)
	})

	http.ListenAndServe(":8080", nil)
}

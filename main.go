package main

import (
	"log"
	"net/http"

	"server-rest-study/internal/data"
)

func main() {
	err := data.InitStorage()
	if err != nil {
		log.Println(
			"Storage initialize " +
				"finished with error: " + err.Error(),
		)
	}

	//http.HandleFunc("endpoint", handler)
	const PORT = ":4444"

	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Println("Server fell with: " + err.Error())
	}
}

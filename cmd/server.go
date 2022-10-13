package main

import (
	"log"
	"net/http"

	"server-rest-study/api/files"
	"server-rest-study/api/session"
	"server-rest-study/api/user"
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

	http.HandleFunc("/api/user", user.Handler)
	http.HandleFunc("/api/session", session.Handler)
	http.HandleFunc("/api/files/", files.Handler)

	const PORT = ":4444"

	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Println("Server fell with: " + err.Error())
	}
}

package main

import (
	"log"
	"net/http"

	"server-rest-study/api/files"
	giveAccess "server-rest-study/api/give_access"
	"server-rest-study/api/session"
	"server-rest-study/api/test"
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

	const PORT = ":4444"
	//http.HandleFunc("endpoint", handler)

	http.HandleFunc("/api/user", user.Handler)
	http.HandleFunc("/api/session", session.Handler)
	http.HandleFunc("/api/files/", files.Handler)
	http.HandleFunc("/api/give_access", giveAccess.Handler)
	http.HandleFunc("/api/test", test.Handler)

	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Println("Server fell with: " + err.Error())
	}
}

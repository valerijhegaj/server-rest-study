package user

import (
	"io"
	"log"
	"net/http"

	"server-rest-study/internal/data"
	"server-rest-study/pkg/file"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println(
			"Bad method for new user, request method:",
			r.Method,
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to create new user:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	parsedBody, err := file.Parse(body)
	username, password := parsedBody.UserName, parsedBody.Password
	if username == "" || err != nil {
		log.Print("Failed to create new user: ")
		if err == nil {
			log.Println(err.Error())
		} else {
			log.Println("unresolved username")
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storage := data.GetStorage()
	_, err = storage.NewUser(username, password)
	if err != nil {
		log.Println("Failed to create new user: " + err.Error())
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusCreated)
	log.Println("Successfuly new user: " + username)
}

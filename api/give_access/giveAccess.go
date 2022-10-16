package giveAccess

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
			"Bad method for set rights, request method:",
			r.Method,
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to set rights:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	parsedBody, err := file.Parse(body)
	if err != nil {
		log.Println("Failed set rights: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	username, path, rights :=
		parsedBody.UserName, parsedBody.Path, parsedBody.Rights

	var token string
	cookies := r.Cookies()
	for _, c := range cookies {
		if c.Name == "token" {
			token = c.Value
		}
	}

	storage := data.GetStorage()
	err = storage.SetRights(token, username, path, rights)
	if err != nil {
		if err.Error() == data.NotAuthorized {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		log.Println("Failed set rights: " + err.Error())
		return
	}
	log.Printf(
		"Successful set rights owner: %s, user: %s, path: %s, rights: %s",
		token, username, path, rights,
	)
	log.Println()
	w.WriteHeader(http.StatusCreated)
}

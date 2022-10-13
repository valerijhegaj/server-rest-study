package user

import (
	"encoding/json"
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
	user, password := parsedBody.User, parsedBody.Password
	if user == "" || err != nil {
		log.Print("Failed to create new user: ")
		if err == nil {
			log.Println(err.Error())
		} else {
			log.Println("unresolved name")
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storage := data.GetStorage()
	userID, err := storage.NewUser(user, password)
	if err != nil {
		log.Println("Failed to create new user: " + err.Error())
		w.WriteHeader(http.StatusForbidden)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(representAsJson(userID))
	log.Println("Successfuly new user: " + user)
}

func representAsJson(userID int) []byte {
	data, _ := json.Marshal(
		struct {
			UserID int `json:"user_id"`
		}{userID},
	)
	return data
}

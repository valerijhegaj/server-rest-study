package session

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
			"Bad method for new session, request method:",
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
	userID, password := parsedBody.UserID, parsedBody.Password
	if err != nil {
		log.Println("Failed new session: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	storage := data.GetStorage()

	token, err := storage.NewToken(userID, password)
	if err != nil {
		log.Println("Failed new session: " + err.Error())
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = storage.NewSession(userID, token)
	if err != nil {
		log.Println("Failed new session: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(representAsJson(token))
	log.Printf("Succesfuly new session: %d\n", userID)
}

func representAsJson(token string) []byte {
	data, _ := json.Marshal(
		struct {
			Token string `json:"token"`
		}{token},
	)
	return data
}

package files

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"server-rest-study/internal/data"
	"server-rest-study/pkg/file"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Failed to", r.Method, "file:", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var parsedBody file.Helper
	if len(body) != 0 {
		parsedBody, err = file.Parse(body)
		file.Parse(body) //delete
		if err != nil {
			log.Println("Failed to", r.Method, "file:", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	token := parsedBody.Token
	userID, name, err := parseURL(r.URL.Path)
	if err != nil {
		log.Println(
			"Failed to", r.Method, "file:", err.Error(),
			r.URL.Path,
		)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		get(userID, name, token, w, r)
	case http.MethodPost:
		post(userID, name, token, parsedBody.FileData, w, r)
	case http.MethodPut:
		put(userID, name, token, parsedBody.FileData, w, r)
	case http.MethodDelete:
		del(userID, name, token, w, r)
	default:
		log.Println(
			"Bad method for work with files, request method:",
			r.Method,
		)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func parseURL(path string) (
	int,
	string,
	error,
) {
	const commonSize int = 11
	path = path[commonSize:]

	ptr := len(path)
	for i := 0; i < len(path); i++ {
		if path[i] == '/' {
			ptr = i
			break
		}
	}
	userIDstring := path[:ptr]
	name := path[ptr+1:]
	userID, err := strconv.Atoi(userIDstring)

	return userID, name, err
}

func get(
	userID int, name, token string,
	w http.ResponseWriter, r *http.Request,
) {
	storage := data.GetStorage()
	//storage.CheckAccess(token, userID, name, "r") not implemented
	rc, err := storage.GetFile(userID, name)
	if err != nil {
		log.Println("Failed to get file:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	content, err := io.ReadAll(rc)
	rc.Close()
	if err != nil {
		log.Println("Failed to get file:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(content)
	log.Println("Successful get file", r.URL.Path)
}

func post(
	userID int, name, token, fileData string,
	w http.ResponseWriter, r *http.Request,
) {
	storage := data.GetStorage()
	//storage.CheckAccess(token, userID, name, "w") not implemented
	err := storage.NewFile(file.NewReadCloser(fileData), userID, name)
	if err != nil {
		log.Println("Failed to post file:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func put(
	userID int, name, token, fileData string,
	w http.ResponseWriter, r *http.Request,
) {
	post(userID, name, token, fileData, w, r)
}

func del(
	userID int, name, token string,
	w http.ResponseWriter, r *http.Request,
) {
	storage := data.GetStorage()

	//storage.CheckAccess(token, userID, name, "w") not implemented
	err := storage.DeleteFile(userID, name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

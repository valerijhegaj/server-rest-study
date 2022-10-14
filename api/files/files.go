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
	userID, path, err := parseURL(r.URL.Path)
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
		get(userID, path, token, w)
	case http.MethodPost:
		post(userID, path, token, parsedBody.FileData, w)
	case http.MethodPut:
		put(userID, path, token, parsedBody.FileData, w)
	case http.MethodDelete:
		del(userID, path, token, w)
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
	path = path[ptr+1:]
	userID, err := strconv.Atoi(userIDstring)

	return userID, path, err
}

func get(userID int, path, token string, w http.ResponseWriter) {
	storage := data.GetStorage()

	isHasAccess, err := storage.CheckAccess(token, userID, path, "r")
	if err != nil {
		log.Println("Failed to get file:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isHasAccess {
		log.Println("Failed to get file: Permission denied")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	rc, err := storage.GetFile(userID, path)
	if err != nil {
		log.Println("Failed to get file:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	content, err := io.ReadAll(rc)
	rc.Close() //need to handle this error
	if err != nil {
		log.Println("Failed to get file:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(content) //need to handle this error
	log.Println("Successful get file")
}

func post(
	userID int, path, token, fileData string, w http.ResponseWriter,
) {
	storage := data.GetStorage()

	isHasAccess, err := storage.CheckAccess(token, userID, path, "w")
	if err != nil {
		log.Println("Failed to get file:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isHasAccess {
		log.Println("Failed to get file: Permission denied")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = storage.NewFile(
		file.NewReadCloserFromString(fileData),
		userID,
		path,
	)
	if err != nil {
		log.Println("Failed to post file:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	log.Println("Success posted file")
}

func put(
	userID int, path, token, fileData string, w http.ResponseWriter,
) {
	post(userID, path, token, fileData, w)
	log.Println("Success put file")
}

func del(userID int, path, token string, w http.ResponseWriter) {
	storage := data.GetStorage()

	isHasAccess, err := storage.CheckAccess(token, userID, path, "w")
	if err != nil {
		log.Println("Failed to get file:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isHasAccess {
		log.Println("Failed to get file: Permission denied")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = storage.DeleteFile(userID, path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Success del file")
}

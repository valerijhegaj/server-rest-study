package files

import (
	"io"
	"log"
	"net/http"

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
		if err != nil {
			log.Println("Failed to", r.Method, "file:", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	var token string
	//if error -> notAuthorized
	cookie, err := r.Cookie("token")
	if err == nil {
		token = cookie.Value
	}

	path, err := parseURL(r.URL.Path)
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
		get(path, token, w)
	case http.MethodPost:
		post(path, token, parsedBody.FileData, w)
	case http.MethodPut:
		put(path, token, parsedBody.FileData, w)
	case http.MethodDelete:
		del(path, token, w)
	default:
		log.Println(
			"Bad method for work with files, request method:",
			r.Method,
		)
		w.WriteHeader(http.StatusBadRequest)
	}
}

func parseURL(path string) (string, error) {
	const commonSize int = 11 // cat /api/files/
	path = path[commonSize:]

	return path, nil
}

func get(path, token string, w http.ResponseWriter) {
	storage := data.GetStorage()

	isHasAccess, err := storage.CheckAccess(token, path, "r")
	if err != nil {
		log.Println("Failed to get file:", err.Error(), token)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isHasAccess {
		log.Println("Failed to get file: Permission denied", token)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	rc, err := storage.GetFile(path)
	if err != nil {
		log.Println("Failed to get file:", err.Error())
		w.WriteHeader(http.StatusNotFound)
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
	path, token, fileData string,
	w http.ResponseWriter,
) {
	storage := data.GetStorage()

	isHasAccess, err := storage.CheckAccess(token, path, "w")
	if err != nil {
		log.Println("Failed to post file:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isHasAccess {
		log.Println("Failed to post file: Permission denied")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = storage.NewFile(
		file.NewReadCloserFromString(fileData),
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
	path, token, fileData string,
	w http.ResponseWriter,
) {
	log.Print("define post put /// ")
	post(path, token, fileData, w)
}

func del(path, token string, w http.ResponseWriter) {
	storage := data.GetStorage()

	isHasAccess, err := storage.CheckAccess(token, path, "w")
	if err != nil {
		log.Println("Failed to del file:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !isHasAccess {
		log.Println("Failed to del file: Permission denied")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err = storage.DeleteFile(path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Success del file")
}

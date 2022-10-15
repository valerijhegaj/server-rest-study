package test

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Header.Get("Cookie"))
	w.Header().Add("Set-Cookie", "token=token; max-age=20")
	w.Write([]byte(r.Header.Get("Cookie")))
}

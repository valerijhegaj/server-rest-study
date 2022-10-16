package test

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Header.Get("Cookie"))
	c := &http.Cookie{Name: "token", Value: "token", MaxAge: 60}
	http.SetCookie(w, c)
	w.Write([]byte(r.Header.Get("Cookie")))
}

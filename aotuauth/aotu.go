package main

import (
	"fmt"
	"log"
	"net/http"
)

func aotuAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Auth")
		if err != nil && cookie.Value != "Authentic" {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		h(w, r)
	}
}
func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Hello,world!"+r.URL.Path)
}
func main() {
	http.HandleFunc("/hello", aotuAuth(Hello))
	err := http.ListenAndServe(":5566", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

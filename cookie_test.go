package Golang_web

import (
	"fmt"
	"net/http"
)

//menulis cookie response http
func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X=PZN-Name"
	cookie.Value = r.URL.Query().Get("name")

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "success create cookie")
}

//membaca cookie response http
func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("X-PZN-Name")
	if err != nil {
		fmt.Fprint(w, "No Cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(w, "Hello %s", name)
	}
}

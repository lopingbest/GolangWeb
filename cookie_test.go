package Golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//menulis cookie response http
func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
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

//dummie server
func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	//dummie server
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//test cookie
func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http:localhost:8080/?name=Galih", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s:%s \n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http:localhost:8080/?name=Galih", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-PZN-Name"
	cookie.Value = "Galih"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

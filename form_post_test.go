package Golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	//parsin
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	//ambil tanpa melakukan parsing
	//r.PostFormValue("first_name")

	//ambil
	first_name := r.PostForm.Get("first_name")
	last_name := r.PostForm.Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", first_name, last_name)
}

func TestFormPost(t *testing.T) {
	requestbody := strings.NewReader("first_name=Galih&last_name=setiadi")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", requestbody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

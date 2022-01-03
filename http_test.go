package Golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//implementasi http_test tanpa running server
func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

func TestHelloHandler(t *testing.T) {
	//nil karena enggak butuh body
	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	//memanggil func diatas
	HelloHandler(recorder, request)

	//Cek hasil test
	response := recorder.Result()
	//io.readAll untuk membaca semua
	body, _ := io.ReadAll(response.Body)
	//konversi body yang semula bentuk byte ke string
	bodyString := string(body)

	fmt.Println(bodyString)
}

package Golang_web

import (
	fmt "fmt"
	"net/http"
	"testing"
)

//Handler berguna untuk menerima HTTP requestyang masuk ke server
func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		//logic web
		//untuk menghindari konversi byte secara manual
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

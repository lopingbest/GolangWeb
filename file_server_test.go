package Golang_web

import (
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileserver := http.FileServer(directory)

	mux := http.NewServeMux()
	//menghapus prefixdi url
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

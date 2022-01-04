package Golang_web

import (
	"embed"
	"io/fs"
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

//memakai embed ke dalam binary distribution file, untuk menghindari copy file secara static
//go:embed resources
var resources embed.FS

func TestFileServerGolangEmbed(t *testing.T) {

	directory, _ := fs.Sub(resources, "resources")
	//mengkonversi resources bawaan golang embed menuju http filesistem
	fileserver := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	//menghapus prefix di url
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

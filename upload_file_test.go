package Golang_web

import (
	"io"
	"net/http"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "upload.form.gohtml", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	//r.ParseMultipartForm(100<<200)  menetapkan masimum memori
	file, fileheader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	//tujuan, lalu nama file
	fileDestionation, err := os.Create("./resources/" + fileheader.Filename)
	if err != nil {
		panic(err)
	}

	//simpan file, ke destinasi file
	_, err = io.Copy(fileDestionation, file)
	if err != nil {
		panic(err)
	}
	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileheader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    ("localhost:8080"),
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

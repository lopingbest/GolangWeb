package Golang_web

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
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

//go:embed resources/pp.jpeg
var uploadfiletest []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	//ngisi field
	writer.WriteField("name", "Galih Setiadi")
	//ngisi file yang akan diupload
	file, _ := writer.CreateFormFile("file", "CONTOHUPLOAD.PNG")
	file.Write(uploadfiletest)

	//untuk memastikan tidak ada memory yang menggantung
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}

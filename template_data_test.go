package Golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./template/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{
		"Title": " Template Data Map",
		"Name":  "Galih",
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Page struct {
	Title string
	Name  string
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./template/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Templae Data Struct",
		Name:  "Galih",
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

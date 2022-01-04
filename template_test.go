package Golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	//variabel untuk menyimpan data dynamic
	templateText := `<html><body>{{.}}</body></html>`
	//parsing
	//t, err := template.New("SIMPLE").Parse(templateText)
	//if err != nil {
	//	panic(err)
	//}

	//parsing versi simple menggunakan must
	t := template.Must(template.New("SIMPLE").Parse(templateText))
	t.ExecuteTemplate(w, "SIMPLE", "Hallo HTML Template")
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

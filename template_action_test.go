package Golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./template/if.gohtml"))
	t.ExecuteTemplate(w, "if.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "Galih",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionOperator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./template/comparator.gohtml"))
	t.ExecuteTemplate(w, "comparator.gohtml", map[string]interface{}{
		"Title":      "Template Data Struct",
		"FinalValue": 90,
	})
}

func TestTemplateActionOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./template/range.gohtml"))
	t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Title": "Template Action Range",
		"Hobbies": []string{
			"Game", "Read", "Code",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./template/address.gohtml"))
	t.ExecuteTemplate(w, "address.gohtml", map[string]interface{}{
		"Title": "Template Action Range",
		"Name":  "Galih",
		"Address": map[string]interface{}{
			"Street": "Jalanin aja dulu",
			"City":   "Semarang",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

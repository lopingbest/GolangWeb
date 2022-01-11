package Golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

//LogMiddleware dijadikan pointer agar tidak duplikat terus
func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("After Execute Handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("Before Execute Handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//ditangkap di defer sebelum selesai
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddlewre(t *testing.T) {
	//mux nerima request
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello Middleware")
	})
	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("foo Executed")
		fmt.Fprint(writer, "Hello Foo")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("foo Executed")
		panic("Ups")
	})
	//log middleware kirim ke mux
	LogMiddleware := &LogMiddleware{
		Handler: mux,
	}
	//error handler kiri mlog ke log middleware
	ErrorHandler := &ErrorHandler{
		Handler: LogMiddleware,
	}

	//request pertama masuk ke server
	server := http.Server{
		Addr: "localhost:8080",
		//dikirim ke error handler
		Handler: ErrorHandler,
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

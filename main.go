package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	s := &http.Server{
		Addr:           ":3000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<html><body><h2>Hello World</h2></body></html>"))
}

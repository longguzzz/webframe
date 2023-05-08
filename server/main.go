package main

import (
	"fmt"
	"log"
	"net/http"

	webframe "web/server/frame"
)

//	curl http://localhost:54321/

func main() {
	engine := webframe.NewEngine()
	engine.GET("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "r: %q\nw: %q\n", r, w)
		if err != nil {
			webframe.ErrBacktrace(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		for key, value := range r.Header {
			_, err := fmt.Fprintf(w, "Header[%q]: %q\n", key, value)
			if err != nil {
				webframe.ErrBacktrace(err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}
	})
	engine.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "r: %q\nw: %q\n", r, w)
		if err != nil {
			webframe.ErrBacktrace(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		for key, value := range r.Header {
			_, err := fmt.Fprintf(w, "Header[%q]: %q\n", key, value)
			if err != nil {
				webframe.ErrBacktrace(err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}
	})
	log.Fatal(http.ListenAndServe("localhost:54321", engine))
}

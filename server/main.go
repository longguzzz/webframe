package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

//	curl http://localhost:54321/

func errBacktrace(err error) {
	bufSize := 1024
	for {
		traceBuf := make([]byte, bufSize)
		n := runtime.Stack(traceBuf, false)
		if n < bufSize {
			traceStr := string(traceBuf[:n])
			log.Printf("Error: %v\nStack Trace:\n%s", err, traceStr)
			break
		}
		bufSize <<= 1
	}

}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello! ", w, r)
		_, err := fmt.Fprintf(w, "r: %q\nw: %q", r, w)
		if err != nil {
			errBacktrace(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("There!")
		_, err := fmt.Fprintf(w, "r: %q\nw: %q", r, w)
		if err != nil {
			errBacktrace(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:54321", nil))
}

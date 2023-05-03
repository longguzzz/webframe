package main

import (
	"fmt"
	"log"
	"net/http"
	"webframe"
)

//	curl http://localhost:54321/

func main() {
	engine := webframe.NewEngine()
	engine.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello! ", w, r)
		// Hello!  &{0xc00009a000 0xc0000a0000 {} 0x4fe640 false false false false {{} 0} {0 0} 0xc0000a2040 {0xc0000ac000 map[] false false} map[] false 0 -1 0 false false [] {{} 0} [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0] 0xc0000aa000 {{} 0}} &{GET / HTTP/1.1 1 1 map[Accept-Encoding:[gzip] User-Agent:[Go-http-client/1.1]] {} <nil> 0 [] false localhost:54321 map[] map[] <nil> map[] 127.0.0.1:51144 / <nil> <nil> <nil> 0xc0000900a0}
		_, err := fmt.Fprintf(w, "r: %q\nw: %q", r, w)
		if err != nil {
			webframe.ErrBacktrace(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
	engine.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("There!")
		_, err := fmt.Fprintf(w, "r: %q\nw: %q", r, w)
		if err != nil {
			webframe.ErrBacktrace(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:54321", engine))
}

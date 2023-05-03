package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:54321")
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}

	// 读取响应头后立即关闭连接
	io.CopyN(io.Discard, resp.Body, int64(resp.ContentLength))
	resp.Body.Close()
}

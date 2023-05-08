package frame

import (
	"log"
	"runtime"
)

func ErrBacktrace(err error) {
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

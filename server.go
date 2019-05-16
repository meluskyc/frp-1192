package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// https://stackoverflow.com/questions/26769626/send-a-chunked-http-response-from-a-go-server?lq=1
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			panic("expected http.ResponseWriter to be an http.Flusher")
		}
		w.Header().Set("X-Content-Type-Options", "nosniff")
		for i := 1; i <= 10000; i++ {
			fmt.Fprintf(w, "Chunk #%d\n", i)
			flusher.Flush() // Trigger "chunked" encoding and send a chunk...
			time.Sleep(500 * time.Millisecond)
		}
	})

	log.Print("Listening on localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

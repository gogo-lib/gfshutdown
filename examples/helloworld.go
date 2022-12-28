package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gogo-lib/gfshutdown"
)

func main() {
	http.HandleFunc("/history", func(w http.ResponseWriter, r *http.Request) {
		// Log GET / history  before shutdown
		gfshutdown.ExecBeforeShutDown(func() {
			now := time.Now()
			log.Printf("GET / at: %s", now)
		})
		w.Write([]byte("Hello World"))
	})

	go http.ListenAndServe(":8080", nil)

	// log exec before shutdown when (ctrl+C)
	gfshutdown.Wait(func() {
		log.Println("exec before shutdown")
	})
}

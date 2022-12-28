package main

import (
	"log"
	"net/http"

	"github.com/gogo-lib/gfshutdown"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	go http.ListenAndServe(":8080", nil)

	// log exec before shutdown when (ctrl+C)
	gfshutdown.Wait(func() {
		log.Println("exec before shutdown")
	})
}

package main

import (
	"log"
	"net/http"

	"github.com/hoanhan101/ultimate-go/go/testing/web_server/handlers"
)

func main() {
	handlers.Routes()

	log.Println("listener : Started : Listening on: http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}

package main

import (
	"log"

	"github.com/Joshdike/proglog/Internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}

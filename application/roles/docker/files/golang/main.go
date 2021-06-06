package main

import (
	"log"
	"net/http"

	"github.com/yvalencia91/devops_test/tree/main/application/roles/docker/files/golang/server"
)

func main() {
	s := server.NewServer()
	log.Println("Started, serving on port 8080")
	err := http.ListenAndServe(":8080", s)

	if err != nil {
		log.Fatal(err.Error())
	}
}

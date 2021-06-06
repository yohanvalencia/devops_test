package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/yvalencia91/devops_test/tree/main/application/roles/docker/files/golang/client"
)

func appHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(time.Now(), "Hello from my new fresh server")

}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func getComments(ctx context.Context) func(w http.ResponseWriter, r *http.Request) {
	jsonPlaceHolder := client.NewClient()
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := jsonPlaceHolder.GetComments(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	ctx := context.Background()

	http.HandleFunc("/", appHandler)
	http.HandleFunc("/_healthcheck", healthCheck)
	http.HandleFunc("/external", getComments(ctx))

	log.Println("Started, serving on port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err.Error())
	}
}

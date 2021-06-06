package server

import (
	"context"
	"encoding/json"
	"fmt"
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

type Server struct {
	http.Handler
}

func NewServer() *Server {
	ctx := context.Background()
	server := new(Server)
	router := http.NewServeMux()
	router.HandleFunc("/", appHandler)
	router.HandleFunc("/_healthcheck", healthCheck)
	router.HandleFunc("/external", getComments(ctx))

	server.Handler = router
	return server
}

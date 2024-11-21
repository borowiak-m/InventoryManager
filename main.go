package main

import (
	"net/http"
)

type api struct {
	addr string
}

func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get users"))
}

func (a *api) createUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created a user"))
}

func main() {
	api := &api{addr: ":8080"}
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}
	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

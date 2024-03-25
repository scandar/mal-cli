package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/scandar/mal-cli/auth"
)

func Start() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		state := r.URL.Query().Get("state")

		fmt.Printf("Code: %v", code)
		fmt.Printf("State: %v", state)

		token := auth.Exchange(state, code)

		fmt.Printf("Token: %v", token)
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":9090", r)

}

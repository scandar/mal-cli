package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/scandar/mal-cli/auth"
	"github.com/scandar/mal-cli/logger"
)

func Start() {
	log := logger.Instance
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		state := r.URL.Query().Get("state")

		log.Debug("Code: ", code)
		log.Debug("State: ", state)

		token := auth.Exchange(state, code)
		log.Debug("Token: ", token)

		w.Write([]byte("<h1>Success! You can close this page and go back to the CLI.</h1>"))
	})
	http.ListenAndServe(":9090", r)
}

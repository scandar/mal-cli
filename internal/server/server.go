package server

import (
	"net/http"
	"time"

	"github.com/scandar/mal-cli/internal/logger"
)

type Code struct {
	Code  string
	State string
}

var code = make(chan Code)

func GetCode() Code {
	s := http.Server{Addr: ":9090"}
	http.HandleFunc("/", rootHandler)
	go s.ListenAndServe()

	return waitForCode()
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log := logger.Instance
	c := r.URL.Query().Get("code")
	s := r.URL.Query().Get("state")

	log.Debug("Code: ", c)
	log.Debug("State: ", s)

	code <- Code{Code: c, State: s}

	w.Write([]byte("<h1>Success! You can close this page and go back to the CLI.</h1>"))
}

func waitForCode() Code {
	log := logger.Instance
	log.Debug("Waiting for code")
	for c := range code {
		time.Sleep(10 * time.Millisecond)

		if c.Code != "" {
			return c
		}
	}

	return Code{}
}

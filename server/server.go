package server

import (
	"fmt"
	"net/http"

	"github.com/scandar/mal-cli/auth"
	"github.com/scandar/mal-cli/logger"
	"github.com/scandar/mal-cli/secrets"
	"github.com/scandar/mal-cli/services/user_service"
)

func Start() {
	http.HandleFunc("/", getRoot)
	http.ListenAndServe(":9090", nil)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	log := logger.Instance
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	log.Debug("Code: ", code)
	log.Debug("State: ", state)

	token := auth.Exchange(state, code)
	secrets.Set("access_token", token.AccessToken)
	secrets.Set("refresh_token", token.RefreshToken)
	log.Debug("Token saved")

	userInfo, err := user_service.GetUserInfo()
	if err != nil {
		log.Error(err)
	}

	fmt.Printf("Logged in as: %s\n", userInfo.Name)

	w.Write([]byte("<h1>Success! You can close this page and go back to the CLI.</h1>"))
}

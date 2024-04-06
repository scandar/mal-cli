package cli

import (
	"fmt"
	"os/exec"

	"github.com/scandar/mal-cli/internal/auth"
	"github.com/scandar/mal-cli/internal/logger"
	"github.com/scandar/mal-cli/internal/secrets"
	"github.com/scandar/mal-cli/internal/server"
	"github.com/scandar/mal-cli/internal/services/user_service"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

var authCMD = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with MyAnimeList",
	Run: func(cmd *cobra.Command, args []string) {
		isDebug, _ := cmd.Flags().GetBool("debug")
		logger.InitLogger(isDebug)
		login()
	},
}

func init() {
	rootCmd.AddCommand(authCMD)
}

func login() {
	log := logger.Instance
	log.Debug("Logging in")
	url := auth.GenerateRedirectURI()
	exec.Command("open", url).Start()

	code := server.GetCode()
	token := exchangeToken(code)
	saveToken(token)
	getUserInfoAndPrint()
}

func exchangeToken(code server.Code) *oauth2.Token {
	return auth.Exchange(code.State, code.Code)
}

func saveToken(token *oauth2.Token) {
	log := logger.Instance
	secrets.Set("access_token", token.AccessToken)
	secrets.Set("refresh_token", token.RefreshToken)
	log.Debug("Token saved")
}

func getUserInfoAndPrint() {
	log := logger.Instance
	userInfo, err := user_service.GetUserInfo()
	if err != nil {
		log.Error(err)
	}

	fmt.Printf("Logged in as: %s\n", userInfo.Name)
}

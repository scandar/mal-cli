package cmd

import (
	"fmt"
	"os/exec"

	"github.com/scandar/mal-cli/auth"
	"github.com/scandar/mal-cli/server"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with MyAnimeList",
	RunE: func(cmd *cobra.Command, args []string) error {
		login()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}

func login() {
	fmt.Println("Logging in")
	url := auth.GenerateRedirectURI()
	exec.Command("open", url).Start()
	server.Start()
}

//
// https://myanimelist.net/v1/oauth2/authorize?
// response_type=code
// &client_id=YOUR_CLIENT_ID
// &state=YOUR_STATE
// &redirect_uri=YOUR_REDIRECT_URI
// &code_challenge=YOUR_PKCE_CODE_CHALLENGE
// &code_challenge_method=plain

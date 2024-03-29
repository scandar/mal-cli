package cmd

import (
	"os/exec"

	"github.com/scandar/mal-cli/auth"
	"github.com/scandar/mal-cli/logger"
	"github.com/scandar/mal-cli/server"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with MyAnimeList",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.InitLogger(isDev)
		login()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}

func login() {
	log := logger.Instance
	log.Debug("Logging in")
	url := auth.GenerateRedirectURI()
	exec.Command("open", url).Start()
	server.Start()
}

package cli

import (
	"os/exec"

	"github.com/scandar/mal-cli/internal/auth"
	"github.com/scandar/mal-cli/internal/logger"
	"github.com/scandar/mal-cli/internal/server"
	"github.com/spf13/cobra"
)

var authCMD = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with MyAnimeList",
	Run: func(cmd *cobra.Command, args []string) {
		logger.InitLogger(isDev)
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
	server.Start()
}

package cmd

import (
	"fmt"

	"github.com/scandar/mal-cli/logger"
	"github.com/scandar/mal-cli/services/anime_service"
	"github.com/spf13/cobra"
)

var status *string
var listAnimeCMD = &cobra.Command{
	Use:     "list-anime",
	Aliases: []string{"la"},
	Short:   "Authenticated user's anime list",
	Run: func(cmd *cobra.Command, args []string) {
		logger.InitLogger(isDev)
		getAnimeList(anime_service.Status(*status), p)
	},
}

func init() {
	status = listAnimeCMD.Flags().StringP("status", "s", "", "Anime status")
	rootCmd.AddCommand(listAnimeCMD)
}

func getAnimeList(s anime_service.Status, p int) {
	log := logger.Instance
	log.Debug("Getting user anime list")
	log.Debugf("Status: %s, Page: %d", s, p)
	animeList, err := anime_service.GetUserAnimeList(s, p)
	if err != nil {
		log.Error(err)
		return
	}

	for _, anime := range animeList.Data {
		fmt.Printf("ID: %d, Title: %s\n", anime.Node.ID, anime.Node.Title)
	}
}

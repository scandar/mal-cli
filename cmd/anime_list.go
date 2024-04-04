package cmd

import (
	"fmt"

	"github.com/scandar/mal-cli/internal/logger"
	"github.com/scandar/mal-cli/services"
	"github.com/scandar/mal-cli/services/anime_service"
	"github.com/spf13/cobra"
)

var listAnimeCMD = &cobra.Command{
	Use:     "anime-list",
	Aliases: []string{"al"},
	Short:   "Authenticated user's anime list",
	Run: func(cmd *cobra.Command, args []string) {
		logger.InitLogger(isDev)
		status, _ := cmd.Flags().GetString("status")
		p, _ := cmd.Flags().GetInt("page")
		getAnimeList(services.AnimeStatus(status), p)
	},
}

func init() {
	listAnimeCMD.Flags().IntP("page", "p", 0, "Page number zero indexed")
	listAnimeCMD.Flags().StringP("status", "s", "", "Anime status")
	rootCmd.AddCommand(listAnimeCMD)
}

func getAnimeList(s services.AnimeStatus, p int) {
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

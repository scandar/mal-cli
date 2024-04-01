package cmd

import (
	"fmt"
	"os"

	"github.com/scandar/mal-cli/logger"
	"github.com/scandar/mal-cli/services/anime_service"
	"github.com/spf13/cobra"
)

var p *int
var animeCMD = &cobra.Command{
	Use:   "anime [QUERY]",
	Short: "Authenticate with MyAnimeList",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logger.InitLogger(isDev)
		if len(args) == 0 {
			fmt.Println("Please provide a query")
			os.Exit(1)
		}
		searchAnime(concatArgs(args), *p)
	},
}

func init() {
	p = animeCMD.Flags().IntP("page", "p", 0, "Page number zero indexed")
	rootCmd.AddCommand(animeCMD)
}

func searchAnime(q string, p int) {
	log := logger.Instance
	log.Debug("Searching anime")

	animeList, err := anime_service.SearchAnime(q, p)
	if err != nil {
		log.Error(err)
		return
	}

	for _, anime := range animeList.Data {
		fmt.Printf("ID: %d, Title: %s\n", anime.Node.ID, anime.Node.Title)
	}
}

func concatArgs(args []string) string {
	q := ""
	for _, arg := range args {
		q += arg + " "
	}
	return q
}

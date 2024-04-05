package cli

import (
	"fmt"
	"os"

	"github.com/scandar/mal-cli/internal/logger"
	"github.com/scandar/mal-cli/internal/services/anime_service"
	"github.com/scandar/mal-cli/internal/utils"
	"github.com/spf13/cobra"
)

var findAnimeCMD = &cobra.Command{
	Use:     "find-anime [QUERY]",
	Aliases: []string{"fa"},
	Short:   "Search for anime by name",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p, _ := cmd.Flags().GetInt("page")
		isDebug, _ := cmd.Flags().GetBool("debug")
		logger.InitLogger(isDebug)
		if len(args) == 0 {
			fmt.Println("Please provide a query")
			os.Exit(1)
		}

		searchAnime(utils.ConcatArgs(args), p)
	},
}

func init() {
	findAnimeCMD.Flags().IntP("page", "p", 0, "Page number zero indexed")
	rootCmd.AddCommand(findAnimeCMD)
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

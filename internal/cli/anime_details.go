package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/scandar/mal-cli/internal/logger"
	"github.com/scandar/mal-cli/internal/services/anime_service"
	"github.com/spf13/cobra"
)

var animeDetailsCMD = &cobra.Command{
	Use:     "anime-details [ID]",
	Aliases: []string{"ad"},
	Args:    cobra.ExactArgs(1),
	Short:   "Get anime details",
	Run: func(cmd *cobra.Command, args []string) {
		isDebug, _ := cmd.Flags().GetBool("debug")
		logger.InitLogger(isDebug)
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid anime ID")
			os.Exit(1)
		}
		getAnimeDetails(id)
	},
}

func init() {
	rootCmd.AddCommand(animeDetailsCMD)
}

func getAnimeDetails(id int) {
	log := logger.Instance
	log.Debug("Getting anime details")
	log.Debugf("Anime ID: %d", id)
	animeDetails, err := anime_service.GetAnimeDetails(id)
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Printf("ID: %d\n", animeDetails.ID)
	fmt.Printf("Title: %s\n", animeDetails.Title)
	fmt.Printf("Episodes: %d\n", animeDetails.NumEpisodes)
	fmt.Printf("Synopsis: %s\n", animeDetails.Synopsis)
	fmt.Printf("Score: %f\n", animeDetails.Mean)
	fmt.Printf("Status: %s\n", animeDetails.Status)
	fmt.Printf("Rank: %d\n", animeDetails.Rank)

	fmt.Println("Genres:")
	for _, genre := range animeDetails.Genres {
		fmt.Printf("  %s\n", genre.Name)
	}

	fmt.Println("My List Status:")
	fmt.Printf("  Episodes Watched: %d\n", animeDetails.MyListStatus.NumEpisodesWatched)
	fmt.Printf("  Score: %d\n", animeDetails.MyListStatus.Score)
	fmt.Printf("  Status: %s\n", animeDetails.MyListStatus.Status)
	fmt.Printf("  Updated At: %s\n", animeDetails.MyListStatus.UpdatedAt)
}

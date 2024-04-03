package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/scandar/mal-cli/logger"
	"github.com/scandar/mal-cli/services/anime_service"
	"github.com/spf13/cobra"
)

var episodes *int
var score *int

var updateAnimeCMD = &cobra.Command{
	Use:     "update-list [ANIME_ID]",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"ul"},
	Short:   "Update an entry in the user's anime list",
	Run: func(cmd *cobra.Command, args []string) {
		logger.InitLogger(isDev)
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid anime ID")
			os.Exit(1)
		}
		updateAnimeList(id, anime_service.Status(*status), *episodes, *score)
	},
}

func init() {
	status = updateAnimeCMD.Flags().StringP("status", "s", "", "Anime status")
	episodes = updateAnimeCMD.Flags().IntP("episodes", "e", 0, "Number of episodes watched")
	score = updateAnimeCMD.Flags().IntP("score", "c", 0, "Score given to the anime")
	rootCmd.AddCommand(updateAnimeCMD)
}

func updateAnimeList(id int, s anime_service.Status, episodes int, score int) {
	log := logger.Instance
	log.Debug("Updating user anime list")
	log.Debugf("Anime ID: %d, Status: %s, Episodes: %d, Score: %d", id, s, episodes, score)
	res, err := anime_service.UpdateUserAnimeList(id, s, episodes, score)
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Printf("Episodes: %d\n", res.NumEpisodesWatched)
	fmt.Printf("Score: %d\n", res.Score)
	fmt.Printf("Status: %s\n", res.Status)
	fmt.Printf("Updated at: %s\n", res.UpdatedAt)
}

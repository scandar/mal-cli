package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/scandar/mal-cli/internal/logger"
	"github.com/scandar/mal-cli/services"
	"github.com/scandar/mal-cli/services/anime_service"
	"github.com/spf13/cobra"
)

var updateAnimeCMD = &cobra.Command{
	Use:     "update-anime [ANIME_ID]",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"ua"},
	Short:   "Update an entry in the user's anime list",
	Run: func(cmd *cobra.Command, args []string) {
		isDebug, _ := cmd.Flags().GetBool("debug")
		logger.InitLogger(isDebug)
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid anime ID")
			os.Exit(1)
		}
		status, _ := cmd.Flags().GetString("status")
		episodes, _ := cmd.Flags().GetInt("episodes")
		score, _ := cmd.Flags().GetInt("score")
		updateAnimeList(id, services.AnimeStatus(status), episodes, score)
	},
}

func init() {
	updateAnimeCMD.Flags().StringP("status", "s", "", "Anime status")
	updateAnimeCMD.Flags().IntP("episodes", "e", 0, "Number of episodes watched")
	updateAnimeCMD.Flags().IntP("score", "o", 0, "Score given to the anime")
	rootCmd.AddCommand(updateAnimeCMD)
}

func updateAnimeList(id int, s services.AnimeStatus, episodes int, score int) {
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

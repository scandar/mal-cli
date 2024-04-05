package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/scandar/mal-cli/internal/logger"
	"github.com/scandar/mal-cli/internal/services/anime_service"
	"github.com/spf13/cobra"
)

var deleteAnimeCMD = &cobra.Command{
	Use:     "delete-anime [ANIME_ID]",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"da"},
	Short:   "Delete an entry from the user's anime list",
	Run: func(cmd *cobra.Command, args []string) {
		isDebug, _ := cmd.Flags().GetBool("debug")
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid anime ID")
			os.Exit(1)
		}

		logger.InitLogger(isDebug)
		deleteAnime(id)
	},
}

func init() {
	rootCmd.AddCommand(deleteAnimeCMD)
}

func deleteAnime(id int) {
	log := logger.Instance
	log.Debug("Deleting anime from user anime list")
	log.Debugf("Anime ID: %d", id)
	success, err := anime_service.DeleteAnime(id)
	if err != nil || !success {
		log.Error(err)
		os.Exit(1)
	}

	fmt.Println("Anime deleted successfully")
}

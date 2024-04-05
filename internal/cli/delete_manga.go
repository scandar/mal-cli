package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/scandar/mal-cli/internal/logger"
	"github.com/scandar/mal-cli/services/manga_service"
	"github.com/spf13/cobra"
)

var deleteMangaCMD = &cobra.Command{
	Use:     "delete-manga [MANGA_ID]",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"dm"},
	Short:   "Delete an entry from the user's manga list",
	Run: func(cmd *cobra.Command, args []string) {
		isDebug, _ := cmd.Flags().GetBool("debug")
		logger.InitLogger(isDebug)
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid manga ID")
			os.Exit(1)
		}
		deleteManga(id)
	},
}

func init() {
	rootCmd.AddCommand(deleteMangaCMD)
}

func deleteManga(id int) {
	log := logger.Instance
	log.Debug("Deleting manga from user manga list")
	log.Debugf("Manga ID: %d", id)
	success, err := manga_service.DeleteManga(id)
	if err != nil || !success {
		log.Error(err)
		os.Exit(1)
	}

	fmt.Println("Manga deleted successfully")
}

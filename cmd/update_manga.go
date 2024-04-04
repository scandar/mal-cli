package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/scandar/mal-cli/internal/logger"
	"github.com/scandar/mal-cli/services"
	"github.com/scandar/mal-cli/services/manga_service"
	"github.com/spf13/cobra"
)

var updateMangaCMD = &cobra.Command{
	Use:     "update-manga [MANGA_ID]",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"um"},
	Short:   "Update an entry in the user's manga list",
	Run: func(cmd *cobra.Command, args []string) {
		logger.InitLogger(isDev)
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid manga ID")
			os.Exit(1)
		}
		status, _ := cmd.Flags().GetString("status")
		chapters, _ := cmd.Flags().GetInt("chapters")
		volumes, _ := cmd.Flags().GetInt("volumes")
		score, _ := cmd.Flags().GetInt("score")
		updateMangaList(id, services.MangaStatus(status), chapters, volumes, score)
	},
}

func init() {
	updateMangaCMD.Flags().StringP("status", "s", "", "Manga status")
	updateMangaCMD.Flags().IntP("chapters", "c", 0, "Number of chapters read")
	updateMangaCMD.Flags().IntP("volumes", "v", 0, "Number of volumes read")
	updateMangaCMD.Flags().IntP("score", "o", 0, "Score given to the manga")
	rootCmd.AddCommand(updateMangaCMD)
}

func updateMangaList(id int, s services.MangaStatus, chapters int, volumes int, score int) {
	log := logger.Instance
	log.Debug("Updating user manga list")
	log.Debugf("Manga ID: %d, Status: %s, Chapters: %d, Volumes: %d, Score: %d", id, s, chapters, volumes, score)
	res, err := manga_service.UpdateUserMangaList(id, s, chapters, volumes, score)
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Printf("Chapters: %d\n", res.NumChaptersRead)
	fmt.Printf("Volumes: %d\n", res.NumVolumesRead)
	fmt.Printf("Score: %d\n", res.Score)
	fmt.Printf("Status: %s\n", res.Status)
	fmt.Printf("Updated at: %s\n", res.UpdatedAt)
}

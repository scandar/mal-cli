package cli

import (
	"fmt"
	"os"
	"strconv"

	"github.com/scandar/mal-cli/internal/logger"
	"github.com/scandar/mal-cli/internal/services/manga_service"
	"github.com/spf13/cobra"
)

var mangaDetailsCMD = &cobra.Command{
	Use:     "manga-details [ID]",
	Aliases: []string{"md"},
	Args:    cobra.ExactArgs(1),
	Short:   "Get manga details",
	Run: func(cmd *cobra.Command, args []string) {
		isDebug, _ := cmd.Flags().GetBool("debug")
		logger.InitLogger(isDebug)
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid manga ID")
			os.Exit(1)
		}
		getMangaDetails(id)
	},
}

func init() {
	rootCmd.AddCommand(mangaDetailsCMD)
}

func getMangaDetails(id int) {
	log := logger.Instance
	log.Debug("Getting manga details")
	log.Debugf("Manga ID: %d", id)
	mangaDetails, err := manga_service.GetMangaDetails(id)
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Printf("ID: %d\n", mangaDetails.ID)
	fmt.Printf("Title: %s\n", mangaDetails.Title)
	fmt.Printf("Chapters: %d\n", mangaDetails.NumChapters)
	fmt.Printf("Volumes: %d\n", mangaDetails.NumVolumes)
	fmt.Printf("Synopsis: %s\n", mangaDetails.Synopsis)
	fmt.Printf("Score: %f\n", mangaDetails.Mean)
	fmt.Printf("Status: %s\n", mangaDetails.Status)
	fmt.Printf("Rank: %d\n", mangaDetails.Rank)

	fmt.Println("Genres:")
	for _, genre := range mangaDetails.Genres {
		fmt.Printf("  %s\n", genre.Name)
	}

	fmt.Println("My List Status:")
	fmt.Printf("  Chapters Read: %d\n", mangaDetails.MyListStatus.NumChaptersRead)
	fmt.Printf("  Volumes Read: %d\n", mangaDetails.MyListStatus.NumVolumesRead)
	fmt.Printf("  Score: %d\n", mangaDetails.MyListStatus.Score)
	fmt.Printf("  Status: %s\n", mangaDetails.MyListStatus.Status)
	fmt.Printf("  Updated At: %s\n", mangaDetails.MyListStatus.UpdatedAt)
}

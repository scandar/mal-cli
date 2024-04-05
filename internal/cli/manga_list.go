package cli

import (
	"fmt"

	"github.com/scandar/mal-cli/internal/logger"
	"github.com/scandar/mal-cli/services"
	"github.com/scandar/mal-cli/services/manga_service"
	"github.com/spf13/cobra"
)

var listMangaCMD = &cobra.Command{
	Use:     "manga-list",
	Aliases: []string{"ml"},
	Short:   "Authenticated user's manga list",
	Run: func(cmd *cobra.Command, args []string) {
		isDebug, _ := cmd.Flags().GetBool("debug")
		status, _ := cmd.Flags().GetString("status")
		p, _ := cmd.Flags().GetInt("page")
		logger.InitLogger(isDebug)

		getMangaList(services.MangaStatus(status), p)
	},
}

func init() {
	listMangaCMD.Flags().IntP("page", "p", 0, "Page number zero indexed")
	listMangaCMD.Flags().StringP("status", "s", "", "Manga status")
	rootCmd.AddCommand(listMangaCMD)
}

func getMangaList(s services.MangaStatus, p int) {
	log := logger.Instance
	log.Debug("Getting user manga list")
	log.Debugf("Status: %s, Page: %d", s, p)
	mangaList, err := manga_service.GetUserMangaList(s, p)
	if err != nil {
		log.Error(err)
		return
	}

	for _, manga := range mangaList.Data {
		fmt.Printf("ID: %d, Title: %s\n", manga.Node.ID, manga.Node.Title)
	}
}

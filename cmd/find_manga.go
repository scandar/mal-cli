package cmd

import (
	"fmt"
	"os"

	"github.com/scandar/mal-cli/logger"
	"github.com/scandar/mal-cli/services/manga_service"
	"github.com/spf13/cobra"
)

var findMangaCMD = &cobra.Command{
	Use:     "find-manga [QUERY]",
	Aliases: []string{"fm"},
	Short:   "Search for manga by name",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logger.InitLogger(isDev)
		if len(args) == 0 {
			fmt.Println("Please provide a query")
			os.Exit(1)
		}
		p, _ := cmd.Flags().GetInt("page")
		searchManga(concatArgs(args), p)
	},
}

func init() {
	findMangaCMD.Flags().IntP("page", "p", 0, "Page number zero indexed")
	rootCmd.AddCommand(findMangaCMD)
}

func searchManga(q string, p int) {
	log := logger.Instance
	log.Debug("Searching manga")

	mangaList, err := manga_service.SearchManga(q, p)
	if err != nil {
		log.Error(err)
		return
	}

	for _, manga := range mangaList.Data {
		fmt.Printf("ID: %d, Title: %s\n", manga.Node.ID, manga.Node.Title)
	}
}

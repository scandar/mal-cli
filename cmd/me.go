package cmd

import (
	"fmt"

	"github.com/scandar/mal-cli/internal/logger"
	"github.com/scandar/mal-cli/services/user_service"
	"github.com/spf13/cobra"
)

var meCMD = &cobra.Command{
	Use:   "me",
	Short: "Get authenticated user info",
	Run: func(cmd *cobra.Command, args []string) {
		logger.InitLogger(isDev)
		userInfo()
	},
}

func init() {
	rootCmd.AddCommand(meCMD)
}

func userInfo() {
	log := logger.Instance
	log.Debug("Fetching user info")

	user, err := user_service.GetUserInfo()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("id:", user.ID)
	fmt.Println("name:", user.Name)
	fmt.Println("location:", user.Location)
	fmt.Println("joined at:", user.JoinedAt)
	fmt.Println("anime statistics:")
	fmt.Println("  num items watching:", user.AnimeStatistics.NumItemsWatching)
	fmt.Println("  num items completed:", user.AnimeStatistics.NumItemsCompleted)
	fmt.Println("  num items on hold:", user.AnimeStatistics.NumItemsOnHold)
	fmt.Println("  num items dropped:", user.AnimeStatistics.NumItemsDropped)
	fmt.Println("  num items plan to watch:", user.AnimeStatistics.NumItemsPlanToWatch)
	fmt.Println("  num items:", user.AnimeStatistics.NumItems)
	fmt.Println("  num days watched:", user.AnimeStatistics.NumDaysWatched)
	fmt.Println("  num days watching:", user.AnimeStatistics.NumDaysWatching)
	fmt.Println("  num days completed:", user.AnimeStatistics.NumDaysCompleted)
	fmt.Println("  num days on hold:", user.AnimeStatistics.NumDaysOnHold)
	fmt.Println("  num days dropped:", user.AnimeStatistics.NumDaysDropped)
	fmt.Println("  num days:", user.AnimeStatistics.NumDays)
	fmt.Println("  num episodes:", user.AnimeStatistics.NumEpisodes)
	fmt.Println("  num times rewatched:", user.AnimeStatistics.NumTimesRewatched)
	fmt.Println("  mean score:", user.AnimeStatistics.MeanScore)

}

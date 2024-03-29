package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var isDev bool
var rootCmd = &cobra.Command{
	Use:   "mal-cli",
	Short: "A CLI tool for MyAnimeList",
	Long:  `A CLI tool for MyAnimeList`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&isDev, "dev", "d", false, "Enable development logs")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var isDev bool

var rootCmd = &cobra.Command{
	Use:   "mal-cli",
	Short: "A CLI tool to managa your MyAnimeList lists",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&isDev, "dev", "d", false, "Enable development logs")
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tmdbcli",
	Short: "TMDB CLI is a command line interface for The Movie Database",
	Long: `A command line interface application that allows you to interact 
with The Movie Database (TMDB) API to search for movies, TV shows, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to TMDB CLI!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
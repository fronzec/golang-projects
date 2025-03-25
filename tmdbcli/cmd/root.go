package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "tmdb-app",
	Short: "TMDB CLI is a command line interface for The Movie Database",
	Long: `A command line interface application that allows you to interact 
with The Movie Database (TMDB) API to search for movies, TV shows, and more.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

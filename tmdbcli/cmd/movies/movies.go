package movies

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	movieType string
	MoviesCmd = &cobra.Command{
		Use:   "movies",
		Short: "Get information about movies from TMDB",
		Long:  `Retrieve information about movies based on different categories like now playing, popular, top rated, and upcoming.`,
		Run: func(cmd *cobra.Command, args []string) {
			switch movieType {
			case "playing":
				fmt.Println("Fetching now playing movies...")
			case "popular":
				fmt.Println("Fetching popular movies...")
			case "top":
				fmt.Println("Fetching top rated movies...")
			case "upcoming":
				fmt.Println("Fetching upcoming movies...")
			default:
				fmt.Printf("Invalid movie type: %s\nValid types are: playing, popular, top, upcoming\n", movieType)
				os.Exit(1)
			}
		},
	}
)

func init() {
	MoviesCmd.Flags().StringVarP(&movieType, "type", "t", "", "Type of movies to fetch (playing, popular, top, upcoming)")
	MoviesCmd.MarkFlagRequired("type")
}

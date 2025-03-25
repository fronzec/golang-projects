package movies

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// Mock data structure for demonstration
type Movie struct {
	ID          int
	Title       string
	ReleaseDate string
	Popularity  float64
}

var (
	movieType string
	MoviesCmd = &cobra.Command{
		Use:   "movies",
		Short: "Get information about movies from TMDB",
		Long:  `Retrieve information about movies based on different categories like now playing, popular, top rated, and upcoming.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Mock data for demonstration
			movies := []Movie{
				{1, "The Shawshank Redemption", "1994-09-23", 9.3},
				{2, "The Godfather", "1972-03-14", 9.2},
				{3, "The Dark Knight", "2008-07-18", 9.0},
			}

			// Create a new table
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)

			// Set table headers
			t.AppendHeader(table.Row{"ID", "Title", "Release Date", "Popularity"})

			// Add rows to the table
			for _, m := range movies {
				t.AppendRow(table.Row{m.ID, m.Title, m.ReleaseDate, m.Popularity})
			}

			// Set some style options
			t.SetStyle(table.StyleRounded)

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
			fmt.Println()
			// Render the table
			t.Render()
		},
	}
)

func init() {
	MoviesCmd.Flags().StringVarP(&movieType, "type", "t", "", "Type of movies to fetch (playing, popular, top, upcoming)")
	MoviesCmd.MarkFlagRequired("type")
}

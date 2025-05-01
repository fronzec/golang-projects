package movies

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/fronzec/golang-projects/tmdbcli/internal/tmdb"
)

var (
	movieType string
	MoviesCmd = &cobra.Command{
		Use:   "movies",
		Short: "Get information about movies from TMDB",
		Long:  `Retrieve information about movies based on different categories like now playing, popular, top rated, and upcoming.`,
		Run: func(cmd *cobra.Command, args []string) {
			t := table.NewWriter()
			t.SetOutputMirror(os.Stdout)
			t.AppendHeader(table.Row{"Title", "Release Date", "Popularity"})

			if movieType == "top" {
				apiKey := tmdb.GetAPIKeyFromEnv()
				if apiKey == "" {
					fmt.Println("TMDB_API_KEY no está configurada en el entorno")
					os.Exit(1)
				}
				client := tmdb.NewClient(apiKey)
				resp, err := client.GetTopRatedMovies(1)
				if err != nil {
					fmt.Printf("Error obteniendo top rated movies: %v\n", err)
					os.Exit(1)
				}
				for _, m := range resp.Results {
					t.AppendRow(table.Row{m.Title, m.ReleaseDate, m.Popularity})
				}
				fmt.Println("Top Rated Movies:")
			}else if movieType == "playing" {
				apiKey := tmdb.GetAPIKeyFromEnv()
				if apiKey == "" {
					fmt.Println("TMDB_API_KEY no está configurada en el entorno")
					os.Exit(1)
				}
				client := tmdb.NewClient(apiKey)
				resp, err := client.GetNowPlayingMovies(1)
				if err != nil {
					fmt.Printf("Error obteniendo now playing movies: %v\n", err)
					os.Exit(1)
				}
				for _, m := range resp.Results {
					t.AppendRow(table.Row{m.Title, m.ReleaseDate, m.Popularity})
				}
				fmt.Println("Now Playing Movies:")

			 }else {
				fmt.Printf("Tipo de película no soportado aún: %s\n", movieType)
				os.Exit(1)
			}
			t.SetStyle(table.StyleRounded)
			t.Render()
		},
	}
)

func init() {
	MoviesCmd.Flags().StringVarP(&movieType, "type", "t", "", "Type of movies to fetch (playing, popular, top, upcoming)")
	MoviesCmd.MarkFlagRequired("type")
}

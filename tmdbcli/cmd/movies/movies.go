package movies

import (
	"fmt"
	"os"

	"github.com/fronzec/golang-projects/tmdbcli/internal/tmdb"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type movieType string

const (
	MovieTypeTop    movieType = "top"
	MovieTypePlaying movieType = "playing"
	MovieTypePopular movieType = "popular"
	MovieTypeUpcoming movieType = "upcoming"
)

var (
	movType string
	MoviesCmd = &cobra.Command{
		Use:   "movies",
		Short: "Get information about movies from TMDB",
		Long:  `Retrieve information about movies based on different categories like now playing, popular, top rated, and upcoming.`,
		Run: func(cmd *cobra.Command, args []string) {
			t := createTable()
			tmdbClient := tmdb.NewClient(getApiKey(), tmdb.NewProdHTTPClient())
			if movType == string(MovieTypeTop) {
				resp, err := tmdbClient.GetTopRatedMovies(1)
				handleErrorOrExit(err, MovieTypeTop)
				fillTable(t, MovieTypeTop, resp.Results)
			}else if movType == string(MovieTypePlaying) {
				resp, err := tmdbClient.GetNowPlayingMovies(1)
				handleErrorOrExit(err, MovieTypePlaying)
				fillTable(t, MovieTypePlaying, resp.Results)
			 } else if movType == string(MovieTypePopular) {
				resp, err := tmdbClient.GetPopularMovies(1)
				handleErrorOrExit(err, MovieTypePopular)
				fillTable(t, MovieTypePopular, resp.Results)
			 } else if movType == string(MovieTypeUpcoming) {
				resp, err := tmdbClient.GetUpcomingMovies(1)
				handleErrorOrExit(err, MovieTypeUpcoming)
				fillTable(t, MovieTypeUpcoming, resp.Results)
			 } else {
				fmt.Printf("Movie type not supported yet: %s\n", movType)
				os.Exit(1)
			}

			printTable(t)
		},
	}
)

func handleErrorOrExit(err error, movType movieType) {
	if err != nil {
		fmt.Printf("Error fetching %s movies\n", movType)
		os.Exit(1)
	}
}

func fillTable(t table.Writer, movType movieType, resp []tmdb.Movie) {
	for _, m := range resp {
		t.AppendRow(table.Row{m.Title, m.ReleaseDate, m.Popularity})
	}
	fmt.Printf("%s Movies:\n", cases.Title(language.English).String(string(movType)))
}

func printTable(t table.Writer) {
	t.Render()
}

func createTable() table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Title", "Release Date", "Popularity"})
	t.SetStyle(table.StyleRounded)
	return t
}

func getApiKey() string {
	apiKey := tmdb.GetAPIKeyFromEnv()
	if apiKey == "" {
		fmt.Println("Auth error: TMDB_API_KEY is not configured in the environment")
		os.Exit(1)
	}
	return apiKey
}

func init() {
	MoviesCmd.Flags().StringVarP(&movType, "type", "t", "", "Type of movies to fetch (playing, popular, top, upcoming)")
	MoviesCmd.MarkFlagRequired("type")
}

package main

import (
	"github.com/fronzec/golang-projects/tmdbcli/cmd"
	"github.com/fronzec/golang-projects/tmdbcli/cmd/movies"
)

func init() {
	cmd.RootCmd.AddCommand(movies.MoviesCmd)
}

func main() {
	cmd.Execute()
}
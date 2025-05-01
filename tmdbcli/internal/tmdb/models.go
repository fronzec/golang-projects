package tmdb

// Movie representa una película individual en la respuesta de top-rated.
type Movie struct {
	Adult            bool     `json:"adult"`
	BackdropPath     string   `json:"backdrop_path"`
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id"`
	OriginalLanguage string   `json:"original_language"`
	OriginalTitle    string   `json:"original_title"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       string   `json:"poster_path"`
	ReleaseDate      string   `json:"release_date"`
	Title            string   `json:"title"`
	Video            bool     `json:"video"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
}

// TopRatedResponse representa la respuesta completa del endpoint top-rated.
type TopRatedResponse struct {
	Page         int     `json:"page"`
	Results      []Movie `json:"results"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
}

// NowPlayingResponse representa la respuesta completa del endpoint now_playing.

type NowPlayingResponse struct {
	Page         int     `json:"page"`
	Results      []Movie `json:"results"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
	Dates        Dates   `json:"dates"`
}

type Dates struct {
	Maximum string `json:"maximum"`
	Minimum string `json:"minimum"`
}
package tmdb

// Movie represents an individual movie in the top-rated response.
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

// TopRatedResponse represents the full response from the top-rated endpoint.
type TopRatedResponse struct {
	Page         int     `json:"page"`
	Results      []Movie `json:"results"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
}

// NowPlayingResponse represents the full response from the now_playing endpoint.

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

// PopularResponse represents the full response from the popular endpoint.
type PopularResponse struct {
	Page         int     `json:"page"`
	Results      []Movie `json:"results"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
}

// UpcomingResponse represents the full response from the upcoming endpoint.
type UpcomingResponse struct {
	Page         int     `json:"page"`
	Results      []Movie `json:"results"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
	Dates        Dates   `json:"dates"`
}

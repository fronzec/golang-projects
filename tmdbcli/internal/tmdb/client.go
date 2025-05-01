package tmdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Client es el cliente para interactuar con la API de TheMovieDB.
type Client struct {
	APIKey  string
	BaseURL string
	Client  *http.Client
}

// NewClient crea una nueva instancia del cliente TMDB.
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: "https://api.themoviedb.org/3",
		Client:  http.DefaultClient,
	}
}

// GetTopRatedMovies obtiene las películas top-rated desde la API.
func (c *Client) GetTopRatedMovies(page int) (*TopRatedResponse, error) {
	url := fmt.Sprintf("%s/movie/top_rated?language=en-US&page=%d", c.BaseURL, page)
	// Crear request manualmente para agregar headers
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la respuesta de TMDB: %s", resp.Status)
	}

	var result TopRatedResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAPIKeyFromEnv obtiene la API Key desde la variable de entorno TMDB_API_KEY.
func GetAPIKeyFromEnv() string {
	return os.Getenv("TMDB_API_KEY")
}

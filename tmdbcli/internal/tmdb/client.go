package tmdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// HTTPClient is an interface to allow mocking of http.Client in tests.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is the client to interact with TheMovieDB API.
type Client struct {
	APIKey  string
	BaseURL string
	Client  HTTPClient
}

// NewClient creates a new instance of the TMDB client.
func NewClient(apiKey string, httpClient HTTPClient) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{
		APIKey:  apiKey,
		BaseURL: "https://api.themoviedb.org/3",
		Client:  httpClient,
	}
}

// GetTopRatedMovies fetches the top-rated movies from the API.
func (c *Client) GetTopRatedMovies(page int) (*TopRatedResponse, error) {
	url := fmt.Sprintf("%s/movie/top_rated?language=en-US&page=%d", c.BaseURL, page)
	// Create request manually to add headers
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	addHeaders(req, c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error in TMDB response: %s", resp.Status)
	}

	var result TopRatedResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetNowPlayingMovies(page int) (*NowPlayingResponse, error) {
	url := fmt.Sprintf("%s/movie/now_playing?language=en-US&page=%d", c.BaseURL, page)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	addHeaders(req, c.APIKey)
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error in TMDB response: %s", resp.Status)
	}

	var result NowPlayingResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetPopularMovies(page int) (*PopularResponse, error) {
	url := fmt.Sprintf("%s/movie/popular?language=en-US&page=%d", c.BaseURL, page)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	addHeaders(req, c.APIKey)
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error in TMDB response: %s", resp.Status)
	}

	var result PopularResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) GetUpcomingMovies(page int) (*UpcomingResponse, error) {
	url := fmt.Sprintf("%s/movie/upcoming?language=en-US&page=%d", c.BaseURL, page)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	addHeaders(req, c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error in TMDB response: %s", resp.Status)
	}

	var result UpcomingResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func addHeaders(req *http.Request, apiKey string) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
}

// GetAPIKeyFromEnv gets the API Key from the TMDB_API_KEY environment variable.
func GetAPIKeyFromEnv() string {
	return os.Getenv("TMDB_API_KEY")
}

func NewProdHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

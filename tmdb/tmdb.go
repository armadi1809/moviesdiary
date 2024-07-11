package tmdb

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type TmdbClient struct {
	apiKey string
}

type TmdbMovie struct {
	Title       string `json:"title"`
	PosterPath  string `json:"poster_path,omitempty"`
	ReleaseDate string `json:"release_date,omitempty"`
	Overview    string `json:"overview,omitempty"`
}

const baseUri = "https://api.themoviedb.org/3/"
const PosterBasePath = "https://image.tmdb.org/t/p/w500"

func NewTmdbClient(apiKey string) *TmdbClient {
	return &TmdbClient{apiKey: apiKey}
}

func (c *TmdbClient) GetNowPlayingMovies() ([]TmdbMovie, error) {
	request, err := http.NewRequest("GET", baseUri+"movie/now_playing", nil)
	if err != nil {
		return nil, err
	}

	q := &url.Values{
		"api_key": []string{c.apiKey},
	}
	request.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	res := struct {
		Results []TmdbMovie `json:"results"`
	}{
		Results: []TmdbMovie{},
	}

	json.Unmarshal(body, &res)

	return res.Results, nil
}

func (c *TmdbClient) GetMovies(query string) ([]TmdbMovie, error) {
	request, err := http.NewRequest("GET", baseUri+"search/movie", nil)
	if err != nil {
		return nil, err
	}

	q := &url.Values{
		"api_key": []string{c.apiKey},
		"query":   []string{query},
	}
	request.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	res := struct {
		Results []TmdbMovie `json:"results"`
	}{
		Results: []TmdbMovie{},
	}
	json.Unmarshal(body, &res)

	return res.Results, nil
}

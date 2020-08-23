package models

type MovieModels struct {
	MovieID       string `json:"movie_id"`
	MovieName     string `json:"movie_name"`
	MovieDuration string `json:"movie_duration"`
	MovieUrlPath  string `json:"movie_url_path"`
	MovieSynopsis string `json:"movie_synopsis"`
}

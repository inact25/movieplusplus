package repositories

import (
	"database/sql"
	"github.com/inact25/movieplusplus/masters/apis/models"
	"log"
)

type MovieRepoImpl struct {
	db *sql.DB
}

func (m MovieRepoImpl) GetAllMovie() ([]*models.MovieModels, error) {
	var dataMovies []*models.MovieModels
	query := "select movieID,movieName,movieDuration,movieUrlPath,movieSinopsis from movie;"
	data, err := m.db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for data.Next() {
		movies := models.MovieModels{}
		err := data.Scan(&movies.MovieID, &movies.MovieName, &movies.MovieDuration, &movies.MovieUrlPath, &movies.MovieSynopsis)
		if err != nil {
			log.Fatal(err)
			return nil, err

		}
		dataMovies = append(dataMovies, &movies)
	}
	return dataMovies, nil
}

func (m MovieRepoImpl) GetMovieByID(id string) ([]*models.MovieModels, error) {
	var dataMovies []*models.MovieModels
	query := "select movieID, movieName,movieDuration,movieUrlPath,movieSinopsis from movie where movieID = ?;"
	data, err := m.db.Query(query, id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for data.Next() {
		movies := models.MovieModels{}
		err := data.Scan(&movies.MovieID, &movies.MovieName, &movies.MovieDuration, &movies.MovieUrlPath, &movies.MovieSynopsis)
		if err != nil {
			log.Fatal(err)
			return nil, err

		}
		dataMovies = append(dataMovies, &movies)
	}
	return dataMovies, nil
}

func (m MovieRepoImpl) AddNewMovie(movies *models.MovieModels) (string, error) {
	tx, err := m.db.Begin()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	addMovies, err := m.db.Prepare("insert into movie (movieName, movieDuration, movieUrlPath, movieSinopsis) values (?,?,?,?)")
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer addMovies.Close()
	if _, err := addMovies.Exec(movies.MovieName, movies.MovieDuration, movies.MovieUrlPath, movies.MovieSynopsis); err != nil {
		tx.Rollback()
		return "", err
	}
	return "", tx.Commit()
}

func InitMovieRepoImpl(db *sql.DB) MoviesRepositories {
	return &MovieRepoImpl{db}
}

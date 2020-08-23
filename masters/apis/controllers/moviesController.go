package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/inact25/movieplusplus/masters/apis/models"
	"github.com/inact25/movieplusplus/masters/apis/usecases"
	"net/http"
)

type MoviesHandler struct {
	MoviesUsecases usecases.MovieUseCases
}

func (h MoviesHandler) GetAllMovies(writer http.ResponseWriter, request *http.Request) {
	menu, err := h.MoviesUsecases.GetAllMovie()
	if err != nil {
		writer.Write([]byte("Data Not Found"))
	}
	var resp = Res{Msg: "getAllMovies", Data: menu}
	byteOfMenu, err := json.Marshal(resp)
	if err != nil {
		writer.Write([]byte("Something when Wrong"))
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(byteOfMenu)
}

func (h MoviesHandler) GetMoviesbyID(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]
	movie, err := h.MoviesUsecases.GetMovieByID(id)
	if err != nil {
		writer.Write([]byte("Data Not Found"))
	}
	var resp = Res{Msg: "getAllMenu", Data: movie}
	byteOfMovie, err := json.Marshal(resp)
	if err != nil {
		writer.Write([]byte("Something when Wrong"))
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(byteOfMovie)
}

func (h MoviesHandler) AddNewMovies(writer http.ResponseWriter, request *http.Request) {
	var resp = Res{}
	movies := &models.MovieModels{}
	writer.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(request.Body).Decode(&movies)
	if err != nil {
		resp = Res{"Decode Failed", nil}
	}
	_, err = h.MoviesUsecases.AddNewMovie(movies)
	if err != nil {
		resp = Res{err.Error(), nil}
	} else {
		resp = Res{"Movies Successfully Added", movies}
	}
	byte, _ := json.Marshal(resp)
	writer.Write(byte)
}

func MoviesControll(r *mux.Router, service usecases.MovieUseCases) {
	UsersHandler := MoviesHandler{service}
	r.HandleFunc("/movies", UsersHandler.GetAllMovies).Methods(http.MethodGet)
	r.HandleFunc("/movies/", UsersHandler.GetAllMovies).Methods(http.MethodGet)

	r.HandleFunc("/movies/{id}", UsersHandler.GetMoviesbyID).Methods(http.MethodGet)
	r.HandleFunc("/movies/{id}", UsersHandler.GetMoviesbyID).Methods(http.MethodGet)

	r.HandleFunc("/movies", UsersHandler.AddNewMovies).Methods(http.MethodPost)
	r.HandleFunc("/movies/", UsersHandler.AddNewMovies).Methods(http.MethodPost)

}

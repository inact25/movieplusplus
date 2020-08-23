package apis

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/inact25/movieplusplus/masters/apis/controllers"
	"github.com/inact25/movieplusplus/masters/apis/repositories"
	"github.com/inact25/movieplusplus/masters/apis/usecases"
)

func Init(r *mux.Router, db *sql.DB) {
	movieRepo := repositories.InitMovieRepoImpl(db)
	movieUseCase := usecases.InitMovieUseCase(movieRepo)
	controllers.MoviesControll(r, movieUseCase)
}

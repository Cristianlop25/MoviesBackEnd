package repository

import (
	"database/sql"

	"backend/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies(genre ...int) ([]*models.Movie, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id int) (*models.User, error)

	OneMovieForEdit(id int) (*models.Movie, []*models.Genre, error)
	OneMovie(id int) (*models.Movie, error)
	AllGenres() ([]*models.Genre, error)
	InsertMovie(movie models.Movie) (int, error)
	UpdateMovieGenres(id int, genresIds []int) error
	UpdateMovie(movie models.Movie) error
	DeleteMovie(id int) error
}

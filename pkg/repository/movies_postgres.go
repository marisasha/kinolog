package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/marisasha/kinolog/pkg/models"
)

type MoviePostgres struct {
	db *sqlx.DB
}

func NewMoviePostgres(db *sqlx.DB) *MoviePostgres {
	return &MoviePostgres{db: db}
}

func (r *MoviePostgres) AddMovie(movie *models.Movie) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	var movieId int
	query := fmt.Sprintf("INSERT INTO %s (type,title,year,description) VALUES ($1,$2,$3,$4) RETURNING id", movieTable)
	row := tx.QueryRow(query, movie.Type, movie.Title, movie.Year, movie.Description)
	if err := row.Scan(&movieId); err != nil {
		return 0, err
	}

	for _, actor := range movie.Actors {
		query = fmt.Sprintf("INSERT INTO %s (movie_id,role,first_name,last_name,bio_url) VALUES ($1,$2,$3,$4,$5)", ActorTable)
		_, err := tx.Exec(query, movieId, actor.Role, actor.FirstName, actor.LastName, actor.BioUrl)
		if err != nil {
			return 0, err
		}
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return movieId, nil
}
func (r *MoviePostgres) GetAllMovies(userId *int) ([]*models.Movie, error) {

	var movies []*models.Movie

	query := fmt.Sprintf(
		`SELECT
			m.id,
			m.type,
			m.title,
			m.year,
			m.description,
			COALESCE(m.poster_url, '') as poster_url,
			um.status,
			um.mark,
			um.review
        FROM %s m
        INNER JOIN %s um ON m.id = um.movie_id
        WHERE um.user_id = $1
        ORDER BY m.id`, movieTable, UserMovieTable)

	err := r.db.Select(&movies, query, *userId)
	if err != nil {
		return nil, err
	}

	for i := range movies {
		actorsQuery := fmt.Sprintf(
			`SELECT id , role, first_name, last_name, bio_url
	     FROM %s WHERE movie_id = $1`, ActorTable)

		err := r.db.Select(&movies[i].Actors, actorsQuery, movies[i].Id)
		if err != nil {
			return nil, err
		}
	}

	return movies, nil
}

func (r *MoviePostgres) GetMovie(movie_id *int) (*models.Movie, error) {
	var movie models.Movie
	query := fmt.Sprintf(
		`SELECT
			m.id,
			m.type,
			m.title,
			m.year,
			m.description,
			COALESCE(m.poster_url, '') as poster_url,
			um.status,
			um.mark,
			um.review
        FROM %s m
        INNER JOIN %s um ON m.id = um.movie_id
        WHERE m.id = $1
        ORDER BY m.id`, movieTable, UserMovieTable)

	err := r.db.Get(&movie, query, *movie_id)
	if err != nil {
		return nil, err
	}

	actorsQuery := fmt.Sprintf(
		`SELECT id , role, first_name, last_name, bio_url
	     FROM %s WHERE movie_id = $1`, ActorTable)

	err = r.db.Select(&movie.Actors, actorsQuery, movie_id)
	if err != nil {
		return nil, err
	}
	return &movie, nil

}

func (r *MoviePostgres) DeleteMovie(movieId *int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", movieTable)
	_, err := r.db.Exec(query, *movieId)
	if err != nil {
		return err
	}
	return nil
}

func (r *MoviePostgres) ChangeMovieStatus(userId, movieId, mark *int, newStatus, review *string) error {
	query := fmt.Sprintf("UPDATE %s SET status=$1 ,mark=$2 ,review=$3 WHERE user_id = $4 AND movie_id = $5", UserMovieTable)
	_, err := r.db.Exec(query, *newStatus, *mark, *review, *userId, *movieId)
	if err != nil {
		return err
	}
	return nil
}

func (r *MoviePostgres) SearchMovieInDB(title *string, year *int) (int, error) {
	var movieId int
	query := fmt.Sprintf("SELECT id FROM %s WHERE title=$1 AND year=$2", movieTable)
	err := r.db.Get(&movieId, query, *title, *year)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return 0, nil
		}
		return 0, err
	}
	return movieId, nil
}

func (r *MoviePostgres) AddUserMovie(userId, movieId *int) error {
	var userMovieId int
	query := fmt.Sprintf("INSERT INTO %s (user_id,movie_id) VALUES ($1,$2) RETURNING ID", UserMovieTable)
	row := r.db.QueryRow(query, *userId, *movieId)
	if err := row.Scan(&userMovieId); err != nil {
		return err
	}
	return nil
}

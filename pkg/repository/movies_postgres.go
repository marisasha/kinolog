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

func (r *MoviePostgres) AddMovie(movie *models.Movie, user_id *int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var movie_id int
	query := fmt.Sprintf("INSERT INTO %s (type,title,year,description) VALUES ($1,$2,$3,$4) RETURNING id", movieTable)
	row := tx.QueryRow(query, movie.Type, movie.Title, movie.Year, movie.Description)
	if err := row.Scan(&movie_id); err != nil {
		return err
	}

	for _, actor := range movie.Actors {
		query = fmt.Sprintf("INSERT INTO %s (movie_id,role,first_name,last_name,bio_url) VALU ($1,$2,$3,$4,$5)", ActorTable)
		_, err := tx.Exec(query, movie_id, actor.Role, actor.FirstName, actor.LastName, actor.BioUrl)
		if err != nil {
			return err
		}
	}

	query = fmt.Sprintf("INSERT INTO %s (user_id,movie_id,status,mark,review) VALUES ($1,$2,$3,$4,$5)", UserMovieTable)
	_, err = tx.Exec(query, user_id, movie_id, movie.Status, movie.Mark, movie.Review)
	if err != nil {
		return err
	}

	return tx.Commit()
}
func (r *MoviePostgres) GetAllMovies(user_id *int) ([]*models.Movie, error) {

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

	err := r.db.Select(&movies, query, *user_id)
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

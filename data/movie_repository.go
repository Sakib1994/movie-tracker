package data

import (
	"database/sql"
	"errors"

	"sakib.com/reelingit/logger"
	"sakib.com/reelingit/models"
)

type MovieRepository struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewMovieRepository(db *sql.DB, log *logger.Logger) (*MovieRepository, error) {
	return &MovieRepository{db: db, logger: log}, nil
}

const defaultLimit = 20

func (r *MovieRepository) GetTopMovies() ([]models.Movie, error) {
	// Fetch movies
	query := `
		SELECT id, tmdb_id, title, tagline, release_year, overview, score, 
		       popularity, language, poster_url, trailer_url
		FROM movies
		ORDER BY popularity DESC
		LIMIT $1
	`
	return r.getMovies(query)
}
func (r *MovieRepository) GetRandomMovies() ([]models.Movie, error) {
	// Fetch random movies
	query := `
		SELECT id, tmdb_id, title, tagline, release_year, overview, score, 
		       popularity, language, poster_url, trailer_url
		FROM movies
		ORDER BY random() DESC
		LIMIT $1
	`
	return r.getMovies(query)
}

func (r *MovieRepository) getMovies(query string) ([]models.Movie, error) {
	rows, err := r.db.Query(query, defaultLimit)
	if err != nil {
		r.logger.Error("Failed to query movies", err)
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var m models.Movie
		if err := rows.Scan(
			&m.ID, &m.TMDB_ID, &m.Title, &m.Tagline, &m.ReleaseYear,
			&m.Overview, &m.Score, &m.Popularity, &m.Language,
			&m.PosterURL, &m.TrailerURL,
		); err != nil {
			r.logger.Error("Failed to scan movie row", err)
			return nil, err
		}
		movies = append(movies, m)
	}

	return movies, nil
}
func (r *MovieRepository) getMovie(query string, id int) (*models.Movie, error) {
	rows, err := r.db.Query(query, id)
	if err != nil {
		r.logger.Error("Failed to query movies", err)
		return nil, err
	}
	defer rows.Close()

	var movies models.Movie
	for rows.Next() {
		var m models.Movie
		if err := rows.Scan(
			&m.ID, &m.TMDB_ID, &m.Title, &m.Tagline, &m.ReleaseYear,
			&m.Overview, &m.Score, &m.Popularity, &m.Language,
			&m.PosterURL, &m.TrailerURL,
		); err != nil {
			r.logger.Error("Failed to scan movie row", err)
			return nil, err
		}
		movies = m
	}

	return &movies, nil
}

var (
	ErrMovieNotFound = errors.New("movie not found")
)

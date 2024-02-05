package repository

import (
	"database/sql"
	"fmt"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/dto"
	"time"
)

// MovieRepository handles movie-related database operations
type MovieRepository struct {
	db *sql.DB
}

// NewMovieRepository creates a new MovieRepository instance
func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{
		db: db,
	}
}

func (r *MovieRepository) SaveMovie(movie dto.Movie) error {
	query := `
        INSERT INTO movies (
            title, year, rated, released, runtime, genre, writer, actors, plot,
            language, country, awards, poster, metascore, imdb_rating, imdb_votes,
            imdb_id, type, box_office, response, created_at, updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)
        RETURNING id
    `

	// Execute the SQL query
	err := r.db.QueryRow(
		query,
		movie.Title, movie.Year, movie.Rated, movie.Released, movie.Runtime, movie.Genre,
		movie.Writer, movie.Actors, movie.Plot, movie.Language, movie.Country,
		movie.Awards, movie.Poster, movie.Metascore, movie.ImdbRating, movie.ImdbVotes,
		movie.ImdbID, movie.Type, movie.BoxOffice,
		movie.Response, time.Now(), time.Now(),
	).Scan(&movie.Id)

	if err != nil {
		return fmt.Errorf("failed to execute SQL query: %v", err)
	}

	return nil

}

func (r *MovieRepository) GetFilteredMovies(genre, actor, year string) ([]dto.Movie, error) {

	query := `
        SELECT * FROM movies where (genre =$1) AND (actors=$2) AND (year=$3)
    `

	rows, err := r.db.Query(query, sql.Named("genre", genre), sql.Named("actors", actor), sql.Named("year", year))

	if err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}
	defer rows.Close()

	var movies []dto.Movie

	for rows.Next() {
		var movie dto.Movie
		fmt.Println("Execution.......", err)
		// Scan the database row into the MovieDTO struct
		if err := rows.Scan(
			&movie.Id, &movie.Title, &movie.Year, &movie.Rated, &movie.Released, &movie.Runtime, &movie.Genre,
			&movie.Writer, &movie.Actors, &movie.Plot, &movie.Language, &movie.Country,
			&movie.Awards, &movie.Poster, &movie.Metascore, &movie.ImdbRating, &movie.ImdbVotes,
			&movie.ImdbID, &movie.Type, &movie.BoxOffice,
			&movie.Response, &movie.CreatedAt, &movie.UpdatedAt,
		); err != nil {
			fmt.Println("Error in db query after scan", err)
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		fmt.Printf("Found movie: %+v\n", movie)

		movies = append(movies, movie)
	}

	return movies, nil
}

func (r *MovieRepository) GetMovieDetailsByTitle(title string) (dto.Movie, error) {
	query := `SELECT * FROM movies WHERE title = $1`

	var movie dto.Movie
	err := r.db.QueryRow(query, title).Scan(
		&movie.Id, &movie.Title, &movie.Year, &movie.Rated, &movie.Released, &movie.Runtime, &movie.Genre,
		&movie.Writer, &movie.Actors, &movie.Plot, &movie.Language, &movie.Country,
		&movie.Awards, &movie.Poster, &movie.Metascore, &movie.ImdbRating, &movie.ImdbVotes,
		&movie.ImdbID, &movie.Type, &movie.BoxOffice,
		&movie.Response, &movie.CreatedAt, &movie.UpdatedAt,
	)
	if err != nil {
		return dto.Movie{}, err
	}
	return movie, nil
}

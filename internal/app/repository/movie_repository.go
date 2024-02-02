package repository

import (
	"database/sql"
	"fmt"
	"github.com/Sakshi1997/GOLANGPROJECT/internal/app/dto"
	"github.com/goccy/go-json"
	"log"
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

// CreateMovie inserts a new movie into the database
func (r *MovieRepository) SaveMovie(movie dto.Movie) error {
	query := `
        INSERT INTO movies (
            title, year, rated, released, runtime, genre, director, writer, actors, plot,
            language, country, awards, poster, ratings, metascore, imdb_rating, imdb_votes,
            imdb_id, type, dvd, box_office, production, website, response, created_at, updated_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26)
        RETURNING id
    `

	// Convert JSONB data to PostgreSQL JSONB format
	ratingsJSONB, err := json.Marshal(movie.Ratings)
	if err != nil {
		return fmt.Errorf("failed to marshal ratings to JSONB: %v", err)
	}

	// Execute the SQL query
	err = r.db.QueryRow(
		query,
		movie.Title, movie.Year, movie.Rated, movie.Released, movie.Runtime, movie.Genre,
		movie.Director, movie.Writer, movie.Actors, movie.Plot, movie.Language, movie.Country,
		movie.Awards, movie.Poster, ratingsJSONB, movie.Metascore, movie.ImdbRating, &movie.Id, &movie.ImdbID, movie.ImdbVotes,
		movie.Type, movie.DVD, movie.BoxOffice, movie.Production, movie.Website,
		movie.Response, time.Now(), time.Now(),
	).Scan(movie.Id)

	if err != nil {
		return fmt.Errorf("failed to execute SQL query: %v", err)
	}

	return nil
}

// GetMovies returns a list of all movies from the database
func (r *MovieRepository) GetMovies() ([]dto.Movie, error) {
	query := "SELECT * FROM movies"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var movie dto.Movie
		// Scan row data into MovieDTO fields
		if err := rows.Scan(
			&movie.Title, &movie.Year, &movie.Rated, &movie.Released, &movie.Runtime,
			&movie.Genre, &movie.Director, &movie.Writer, &movie.Actors, &movie.Plot,
			&movie.Language, &movie.Country, &movie.Awards, &movie.Poster, &movie.Ratings,
			&movie.Metascore, &movie.ImdbRating, &movie.ImdbVotes, &movie.ImdbID, &movie.Type,
			&movie.DVD, &movie.BoxOffice, &movie.Production, &movie.Id, &movie.Website, &movie.Response); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
	}
	return nil, nil
}

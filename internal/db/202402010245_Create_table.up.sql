CREATE TABLE IF NOT EXISTS movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    year VARCHAR(10),
    rated VARCHAR(50),
    released DATE,
    runtime VARCHAR(10),
    genre VARCHAR(255),
    writer VARCHAR(255),
    actors VARCHAR(255),
    plot TEXT,
    language VARCHAR(50),
    country VARCHAR(50),
    awards VARCHAR(255),
    poster VARCHAR(255),
    metascore VARCHAR(10),
    imdb_rating VARCHAR(20),
    imdb_votes VARCHAR(255),
    imdb_id VARCHAR(255),
    type VARCHAR(255),
    box_office VARCHAR(255),
    response VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
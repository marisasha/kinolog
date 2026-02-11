CREATE TYPE media_type AS ENUM ('film', 'serial');
CREATE TYPE watch_status AS ENUM ('planned','watched');
CREATE TYPE role_type AS ENUM ('actor', 'director');


-- CREATE TABLE users (
--     id SERIAL PRIMARY KEY,
--     email VARCHAR(255) UNIQUE NOT NULL,
--     password_hash VARCHAR(255) NOT NULL,
--     first_name VARCHAR(100),
--     last_name VARCHAR(100),
--     avatar_url VARCHAR(500),
--     registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
-- );

CREATE TABLE movie (
    id SERIAL PRIMARY KEY,
    type media_type NOT NULL, 
    title VARCHAR(255) NOT NULL,       
    year INTEGER,                             
    description TEXT,                      
    poster_url VARCHAR(500),              
    
);

CREATE TABLE user_movie(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    movie_id INTEGER NOT NULL REFERENCES movie(id) ON DELETE CASCADE,
    status watch_status NOT NULL DEFAULT 'planned',
    mark INTEGER CHECK (mark >= 1 AND mark <= 10),
    review TEXT,

    UNIQUE(user_id, movie_id),
);

CREATE TABLE movie_actors (
    id SERIAL PRIMARY KEY ,
    movie_id INTEGER NOT NULL REFERENCES movie(id) ON DELETE CASCADE,
    role role_type NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    bio_url VARCHAR(500),

);

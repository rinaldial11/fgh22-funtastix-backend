CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    profile_id INT REFERENCES profiles (id),
    email VARCHAR(255),
    password VARCHAR(255),
    role VARCHAR(10),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

CREATE TABLE profiles (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    phone_number VARCHAR(16),
    point VARCHAR(255),
    picture VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255),
    image VARCHAR(255),
    banner VARCHAR(255),
    release_date DATE,
    author VARCHAR(255),
    duration VARCHAR,
    synopsis TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

CREATE TABLE movie_genre (
    id SERIAL PRIMARY KEY,
    movie_id INT REFERENCES movies (id),
    genre_name VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

CREATE TABLE movie_cast (
    id SERIAL PRIMARY KEY,
    movie_id INT REFERENCES movies (id),
    cast_name VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    profile_id INT REFERENCES profiles (id),
    movie_id INT REFERENCES movies (id),
    date DATE,
    time VARCHAR(255),
    location VARCHAR(255),
    cinema VARCHAR(255),
    seat VARCHAR(255),
    payment_method VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);
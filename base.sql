CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    profile_id INT REFERENCES profiles (id),
    email VARCHAR(255) UNIQUE,
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
    uploaded_by int REFERENCES users (id),
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
    user_id INT REFERENCES users (id),
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

CREATE TABLE show_dates (
    id serial PRIMARY KEY,
    date date,
    created_at timestamp,
    updated_at timestamp
);

CREATE TABLE show_times (
    id serial PRIMARY KEY,
    time varchar,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp
);

CREATE TABLE show_locations (
    id serial PRIMARY KEY,
    location varchar,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp
);

CREATE TABLE show_cinemas (
    id serial PRIMARY KEY,
    cinema varchar,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp
);

CREATE TABLE seats (
    id serial PRIMARY KEY,
    seat varchar,
    price int,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp
);

CREATE TABLE payment_methods (
    id serial PRIMARY KEY,
    method varchar,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp
);

CREATE TABLE orders (
    id serial PRIMARY KEY,
    user_id int REFERENCES users (id),
    movie_id int REFERENCES movies (id),
    date_id int REFERENCES show_dates (id),
    time_id int REFERENCES show_times (id),
    location_id int REFERENCES show_locations (id),
    cinema_id int REFERENCES show_cinemas (id),
    seat_id int REFERENCES seats (id),
    payment_method_id int REFERENCES payment_methods (id),
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp
);
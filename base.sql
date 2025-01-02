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

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    profile_id INT REFERENCES profiles (id),
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
    role VARCHAR(10),
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
    duration VARCHAR(255),
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

CREATE TABLE show_dates (
    id serial PRIMARY KEY,
    date date,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp
);

CREATE TABLE show_times (
    id serial PRIMARY KEY,
    time varchar(255),
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp
);

CREATE TABLE show_locations (
    id serial PRIMARY KEY,
    location varchar(255),
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp
);

CREATE TABLE show_cinemas (
    id serial PRIMARY KEY,
    cinema varchar(255),
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp
);

CREATE TABLE seats (
    id serial PRIMARY KEY,
    seat varchar(255),
    price int,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp
);

CREATE TABLE payment_methods (
    id serial PRIMARY KEY,
    method varchar(255),
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
    payment_method_id int REFERENCES payment_methods (id),
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp
);

CREATE TABLE seats_order (
    id SERIAL PRIMARY KEY,
    seat_id int REFERENCES seats (id),
    order_id int REFERENCES orders (id),
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp
)

INSERT INTO
    movies (
        title,
        image,
        banner,
        release_date,
        author,
        duration,
        synopsis,
        uploaded_by
    )
VALUES (
        'Avengers: Endgame',
        'https://example.com/image1.jpg',
        'https://example.com/banner1.jpg',
        '2019-04-26',
        'Anthony Russo, Joe Russo',
        '03:02:00',
        'The Avengers assemble to reverse the damage caused by Thanos in Avengers: Infinity War.',
        1
    ),
    (
        'The Lion King',
        'https://example.com/image2.jpg',
        'https://example.com/banner2.jpg',
        '2019-07-19',
        'Jon Favreau',
        '01:58:00',
        'A young lion prince flees his kingdom only to learn the true meaning of responsibility and bravery.',
        1
    ),
    (
        'Inception',
        'https://example.com/image3.jpg',
        'https://example.com/banner3.jpg',
        '2010-07-16',
        'Christopher Nolan',
        '02:28:00',
        'A thief who enters the dreams of others to steal secrets from their subconscious is given a chance to have his criminal record erased.',
        1
    ),
    (
        'Titanic',
        'https://example.com/image4.jpg',
        'https://example.com/banner4.jpg',
        '1997-12-19',
        'James Cameron',
        '03:14:00',
        'A seventeen-year-old aristocrat falls in love with a kind but poor artist aboard the luxurious, ill-fated R.M.S. Titanic.',
        1
    ),
    (
        'The Matrix',
        'https://example.com/image5.jpg',
        'https://example.com/banner5.jpg',
        '1999-03-31',
        'Lana Wachowski, Lilly Wachowski',
        '02:16:00',
        'A computer hacker learns from mysterious rebels about the true nature of his reality and his role in the war against its controllers.',
        1
    ),
    (
        'The Dark Knight',
        'https://example.com/image6.jpg',
        'https://example.com/banner6.jpg',
        '2008-07-18',
        'Christopher Nolan',
        '02:32:00',
        'When the menace known as The Joker emerges from his mysterious past, he wreaks havoc and chaos on the people of Gotham.',
        1
    ),
    (
        'Interstellar',
        'https://example.com/image7.jpg',
        'https://example.com/banner7.jpg',
        '2014-11-07',
        'Christopher Nolan',
        '02:49:00',
        'A team of explorers travel through a wormhole in space in an attempt to ensure humanity survival.',
        1
    ),
    (
        'Jurassic Park',
        'https://example.com/image8.jpg',
        'https://example.com/banner8.jpg',
        '1993-06-11',
        'Steven Spielberg',
        '02:07:00',
        'During a preview tour, a theme park suffers a major power breakdown that allows its cloned dinosaur exhibits to run amok.',
        1
    ),
    (
        'Guardians of the Galaxy',
        'https://example.com/image9.jpg',
        'https://example.com/banner9.jpg',
        '2014-08-01',
        'James Gunn',
        '02:01:00',
        'A group of intergalactic criminals are forced to work together to stop a fanatical warrior from taking control of the universe.',
        1
    ),
    (
        'Forrest Gump',
        'https://example.com/image10.jpg',
        'https://example.com/banner10.jpg',
        '1994-07-06',
        'Robert Zemeckis',
        '02:22:00',
        'The presidencies of Kennedy and Johnson, the Vietnam War, the Watergate scandal and other historical events unfold from the perspective of an Alabama man with an extraordinary life story.',
        1
    ),
    (
        'Spider-Man: No Way Home',
        'https://example.com/image11.jpg',
        'https://example.com/banner11.jpg',
        '2021-12-17',
        'Jon Watts',
        '02:28:00',
        'Spider-Man seeks help from Doctor Strange to make his identity as Peter Parker a secret again, but the spell goes wrong.',
        1
    ),
    (
        'Frozen II',
        'https://example.com/image12.jpg',
        'https://example.com/banner12.jpg',
        '2019-11-22',
        'Chris Buck, Jennifer Lee',
        '01:43:00',
        'Anna, Elsa, Kristoff, Olaf and Sven embark on a dangerous but remarkable journey to discover the truth behind Elsa powers.',
        1
    ),
    (
        'Shrek',
        'https://example.com/image13.jpg',
        'https://example.com/banner13.jpg',
        '2001-04-22',
        'Andrew Adamson, Vicky Jenson',
        '01:30:00',
        'An ogre quiet life is disrupted when a host of fairy tale characters are exiled to his swamp by a corrupt lord.',
        1
    ),
    (
        'The Godfather',
        'https://example.com/image14.jpg',
        'https://example.com/banner14.jpg',
        '1972-03-24',
        'Francis Ford Coppola',
        '02:55:00',
        'The aging patriarch of an organized crime dynasty transfers control of his clandestine empire to his reluctant son.',
        1
    ),
    (
        'Star Wars: A New Hope',
        'https://example.com/image15.jpg',
        'https://example.com/banner15.jpg',
        '1977-05-25',
        'George Lucas',
        '02:01:00',
        'Luke Skywalker joins forces with a Jedi knight, a cocky pilot, a Wookiee and two droids to rescue Princess Leia from an evil galactic empire.',
        1
    ),
    (
        'Pulp Fiction',
        'https://example.com/image16.jpg',
        'https://example.com/banner16.jpg',
        '1994-10-14',
        'Quentin Tarantino',
        '02:34:00',
        'The lives of two mob hitmen, a boxer, a gangster wife, and a pair of diner bandits intertwine in four tales of violence and redemption.',
        1
    ),
    (
        'The Shawshank Redemption',
        'https://example.com/image17.jpg',
        'https://example.com/banner17.jpg',
        '1994-09-22',
        'Frank Darabont',
        '02:22:00',
        'Two imprisoned men bond over a number of years, finding solace and eventual redemption through acts of common decency.',
        1
    ),
    (
        'Mad Max: Fury Road',
        'https://example.com/image18.jpg',
        'https://example.com/banner18.jpg',
        '2015-05-15',
        'George Miller',
        '02:00:00',
        'In a post-apocalyptic wasteland, Max teams up with a mysterious woman to escape a tyrannical warlord.',
        1
    ),
    (
        'The Lord of the Rings: The Fellowship of the Ring',
        'https://example.com/image19.jpg',
        'https://example.com/banner19.jpg',
        '2001-12-19',
        'Peter Jackson',
        '02:58:00',
        'A young hobbit, Frodo Baggins, is tasked with the dangerous journey of destroying the One Ring to ensure the destruction of its master, Sauron.',
        1
    ),
    (
        'Jaws',
        'https://example.com/image20.jpg',
        'https://example.com/banner20.jpg',
        '1975-06-20',
        'Steven Spielberg',
        '02:04:00',
        'A great white shark terrorizes a small island community, prompting a police chief, a marine biologist, and a professional shark hunter to catch it.',
        1
    );
    
INSERT INTO
    movie_genre (movie_id, genre_name)
VALUES (1, 'Action'),
    (1, 'Adventure'),
    (1, 'Sci-Fi'),
    (2, 'Animation'),
    (2, 'Adventure'),
    (3, 'Action'),
    (3, 'Sci-Fi'),
    (4, 'Drama'),
    (4, 'Romance'),
    (5, 'Action'),
    (5, 'Sci-Fi'),
    (6, 'Action'),
    (6, 'Crime'),
    (7, 'Adventure'),
    (7, 'Sci-Fi'),
    (8, 'Action'),
    (8, 'Sci-Fi'),
    (9, 'Action'),
    (9, 'Adventure'),
    (10, 'Drama'),
    (10, 'Romance'),
    (11, 'Crime'),
    (11, 'Drama'),
    (12, 'Action'),
    (12, 'Sci-Fi'),
    (13, 'Comedy'),
    (13, 'Animation'),
    (14, 'Drama'),
    (14, 'Thriller'),
    (15, 'Horror'),
    (15, 'Adventure'),
    (16, 'Action'),
    (16, 'Crime'),
    (17, 'Drama'),
    (17, 'Crime'),
    (18, 'Action'),
    (18, 'Adventure'),
    (19, 'Fantasy'),
    (19, 'Adventure'),
    (20, 'Horror'),
    (20, 'Thriller');

INSERT INTO
    movie_cast (movie_id, cast_name)
VALUES (1, 'Robert Downey Jr.'),
    (1, 'Chris Evans'),
    (2, 'Donald Glover'),
    (2, 'Beyoncé'),
    (3, 'Leonardo DiCaprio'),
    (3, 'Joseph Gordon-Levitt'),
    (4, 'Leonardo DiCaprio'),
    (4, 'Kate Winslet'),
    (5, 'Keanu Reeves'),
    (5, 'Laurence Fishburne'),
    (6, 'Christian Bale'),
    (6, 'Heath Ledger'),
    (7, 'Matthew McConaughey'),
    (7, 'Anne Hathaway'),
    (8, 'Sam Neill'),
    (8, 'Laura Dern'),
    (9, 'Chris Pratt'),
    (9, 'Zoe Saldana'),
    (10, 'Tom Hanks'),
    (10, 'Robin Wright'),
    (11, 'Al Pacino'),
    (11, 'Robert De Niro'),
    (12, 'Matthew McConaughey'),
    (12, 'Anne Hathaway'),
    (13, 'Mike Myers'),
    (13, 'Eddie Murphy'),
    (14, 'Marlon Brando'),
    (14, 'Al Pacino'),
    (15, 'Mark Hamill'),
    (15, 'Carrie Fisher'),
    (16, 'John Travolta'),
    (16, 'Uma Thurman'),
    (17, 'Tim Robbins'),
    (17, 'Morgan Freeman'),
    (18, 'Tom Hardy'),
    (18, 'Charlize Theron'),
    (19, 'Elijah Wood'),
    (19, 'Ian McKellen'),
    (20, 'Roy Scheider'),
    (20, 'Robert Shaw');

INSERT INTO
    show_dates ("date")
VALUES ('2025-01-03'),
    ('2025-01-04'),
    ('2025-01-05'),
    ('2025-01-06'),
    ('2025-01-07'),
    ('2025-01-08'),
    ('2025-01-09'),
    ('2025-01-10'),
    ('2025-01-11'),
    ('2025-01-12'),
    ('2025-01-13'),
    ('2025-01-014'),
    ('2025-01-15'),
    ('2025-01-16'),
    ('2025-01-17'),
    ('2025-01-18'),
    ('2025-01-19'),
    ('2025-01-20');

INSERT INTO
    show_times ("time")
VALUES ('11:30'),
    ('12:00'),
    ('12:30'),
    ('13:00'),
    ('13:30'),
    ('14:00'),
    ('14:30'),
    ('15:00'),
    ('15:30'),
    ('16:00'),
    ('16:30'),
    ('17:00'),
    ('17:30'),
    ('18:30'),
    ('19:00'),
    ('19:30'),
    ('20:00'),
    ('20:30'),
    ('21:00'),
    ('21:30');

INSERT INTO
    show_locations (location)
VALUES ('Jakarta'),
    ('Tangerang'),
    ('Bekasi'),
    ('Depok'),
    ('Bogor'),
    ('Bandung'),
    ('Cimahi'),
    ('Surabaya'),
    ('Semarang'),
    ('Padang'),
    ('Medan');

INSERT INTO
    show_cinemas (cinema)
VALUES ('ebv.id'),
    ('CineOne21'),
    ('hiflix');

INSERT INTO
    seats (seat, price)
VALUES ('A1', 35000),
    ('A2', 35000),
    ('A3', 35000),
    ('A4', 35000),
    ('A5', 35000),
    ('A6', 35000),
    ('A7', 35000),
    ('A8', 35000),
    ('A9', 35000),
    ('A10', 35000),
    ('A11', 35000),
    ('A12', 35000),
    ('A13', 35000),
    ('A14', 35000),
    ('B1', 35000),
    ('B2', 35000),
    ('B3', 35000),
    ('B4', 35000),
    ('B5', 35000),
    ('B6', 35000),
    ('B7', 35000),
    ('B8', 35000),
    ('B9', 35000),
    ('B10', 35000),
    ('B11', 35000),
    ('B12', 35000),
    ('B13', 35000),
    ('B14', 35000),
    ('C1', 35000),
    ('C2', 35000),
    ('C3', 35000),
    ('C4', 35000),
    ('C5', 35000),
    ('C6', 35000),
    ('C7', 35000),
    ('C8', 35000),
    ('C9', 35000),
    ('C10', 35000),
    ('C11', 35000),
    ('C12', 35000),
    ('C13', 35000),
    ('C14', 35000),
    ('D1', 35000),
    ('D2', 35000),
    ('D3', 35000),
    ('D4', 35000),
    ('D5', 35000),
    ('D6', 35000),
    ('D7', 35000),
    ('D8', 35000),
    ('D9', 35000),
    ('D10', 35000),
    ('D11', 35000),
    ('D12', 35000),
    ('D13', 35000),
    ('D14', 35000),
    ('E1', 35000),
    ('E2', 35000),
    ('E3', 35000),
    ('E4', 35000),
    ('E5', 35000),
    ('E6', 35000),
    ('E7', 35000),
    ('E8', 35000),
    ('E9', 35000),
    ('E10', 35000),
    ('E11', 35000),
    ('E12', 35000),
    ('E13', 35000),
    ('E14', 35000),
    ('F1', 35000),
    ('F2', 35000),
    ('F3', 35000),
    ('F4', 35000),
    ('F5', 35000),
    ('F6', 35000),
    ('F7', 35000),
    ('F8', 35000),
    ('F9', 35000),
    ('F10', 35000),
    ('F11', 35000),
    ('F12', 35000),
    ('F13', 35000),
    ('F14', 35000),
    ('G1', 35000),
    ('G2', 35000),
    ('G3', 35000),
    ('G4', 35000),
    ('G5', 35000),
    ('G6', 35000),
    ('G7', 35000),
    ('G8', 35000),
    ('G9', 35000),
    ('G10', 35000),
    ('G11', 35000),
    ('G12', 35000),
    ('G13', 35000),
    ('G14', 35000);

INSERT INTO
    payment_methods (method)
VALUES ('Google Pay'),
    ('Visa'),
    ('Gopay'),
    ('Paypall'),
    ('Dana'),
    ('BCA'),
    ('BRI'),
    ('Ovo');
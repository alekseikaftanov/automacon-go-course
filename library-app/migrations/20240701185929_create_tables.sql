-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied.

CREATE TABLE authors (
                         id SERIAL PRIMARY KEY,
                         name TEXT NOT NULL,
                         surname TEXT NOT NULL,
                         bio TEXT,
                         birthday DATE
);

CREATE TABLE books (
                       id SERIAL PRIMARY KEY,
                       title TEXT NOT NULL,
                       author_id INT NOT NULL,
                       year INT,
                       isbn TEXT,
                       FOREIGN KEY (author_id) REFERENCES authors(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back.

DROP TABLE books;
DROP TABLE authors;
package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"library-app/internal/models"
	"library-app/internal/repository"
)

type Postgres struct {
	db *sql.DB
}

// NewPostgres creates a new Postgres repository.
func NewPostgres(ctx context.Context, dsn string) (repository.Repository, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("postgres connect: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("postgres ping: %w", err)
	}
	return &Postgres{db: db}, nil
}

// GetAllBooks implements repository.Repository.
func (p *Postgres) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	query := `SELECT id, title, author_id, year, isbn FROM books`
	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("postgres get all books: %w", err)
	}
	defer rows.Close()

	var books []models.Book // Срез для хранения книг
	for rows.Next() {       // Перебор строк
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.AuthorID, &book.Year, &book.ISBN); err != nil {
			return nil, fmt.Errorf("postgres get all books: %w", err)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("postgres get all books: %w", err)
	}

	return books, nil
}

// GetBookByID implements repository.Repository.
func (p *Postgres) GetBookByID(ctx context.Context, id int) (models.Book, error) {
	query := `SELECT id, title, author_id, year, isbn FROM books WHERE id = $1`
	row := p.db.QueryRowContext(ctx, query, id)
	var book models.Book
	err := row.Scan(&book.ID, &book.Title, &book.AuthorID, &book.Year, &book.ISBN)
	if err != nil {
		return models.Book{}, fmt.Errorf("postgres get book by id: %w", err)
	}
	return book, nil
}

func (p *Postgres) CreateBook(ctx context.Context, book models.Book) (int, error) {
	query := `INSERT INTO books (title, author_id, year, isbn) VALUES ($1, $2, $3, $4) RETURNING id`
	row := p.db.QueryRowContext(ctx, query, book.Title, book.AuthorID, book.Year, book.ISBN)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("postgres create book: %w", err)
	}
	return id, nil
}

// Метод проверки на повторное создание книги

func (p *Postgres) GetBookByDetails(ctx context.Context, title string, authorID int, year int, isbn string) (models.Book, error) {
	query := `SELECT id, title, author_id, year, isbn FROM books WHERE title = $1 AND author_id = $2 AND year = $3 AND isbn = $4`
	row := p.db.QueryRowContext(ctx, query, title, authorID, year, isbn)
	var book models.Book
	err := row.Scan(&book.ID, &book.Title, &book.AuthorID, &book.Year, &book.ISBN)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Book{}, nil // Книга не найдена
		}
		return models.Book{}, fmt.Errorf("postgres get book by details: %w", err)
	}
	return book, nil
}

func (p *Postgres) GetAuthorIDByName(ctx context.Context, name, surname string) (int, error) {
	query := `SELECT id FROM authors WHERE name = $1 AND surname = $2`
	row := p.db.QueryRowContext(ctx, query, name, surname)
	var id int
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil // Автор не найден
		}
		return 0, fmt.Errorf("postgres get author by name: %w", err)
	}
	return id, nil
}

func (p *Postgres) UpdateBook(ctx context.Context, book models.Book) (int, error) {
	query := `UPDATE books SET title = $1, author_id = $2, year = $3, isbn = $4 WHERE id = $5 RETURNING id`
	row := p.db.QueryRowContext(ctx, query, book.Title, book.AuthorID, book.Year, book.ISBN, book.ID)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("postgres update book: %w", err)
	}
	return id, nil
}

func (p *Postgres) DeleteBook(ctx context.Context, id int) (int, error) {
	query := `DELETE FROM books WHERE id = $1 RETURNING id`
	row := p.db.QueryRowContext(ctx, query, id)
	var deletedID int
	err := row.Scan(&deletedID)
	if err != nil {
		return 0, fmt.Errorf("postgres delete book: %w", err)
	}
	return deletedID, nil
}

func (p *Postgres) CreateAuthor(ctx context.Context, author models.Author) (int, error) {
	query := `INSERT INTO authors (name, surname, bio, birthday) VALUES ($1, $2, $3, $4) RETURNING id`
	row := p.db.QueryRowContext(ctx, query, author.Name, author.Surname, author.Bio, author.Birthday)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("postgres create author: %w", err)
	}
	return id, nil
}

// GetAuthorByID implements repository.Repository.
func (p *Postgres) GetAuthorByID(ctx context.Context, id int) (models.Author, error) {
	query := `SELECT id, name, surname, bio, birthday FROM authors WHERE id = $1`
	row := p.db.QueryRowContext(ctx, query, id)
	var author models.Author
	err := row.Scan(&author.ID, &author.Name, &author.Surname, &author.Bio, &author.Birthday)
	if err != nil {
		return models.Author{}, fmt.Errorf("postgres get author by id: %w", err)
	}
	return author, nil
}

func (p *Postgres) GetAuthorByDetails(ctx context.Context, name, surname string) (models.Author, error) {
	query := `SELECT id, name, surname, bio, birthday FROM authors WHERE name = $1 AND surname = $2`
	row := p.db.QueryRowContext(ctx, query, name, surname)
	var author models.Author
	err := row.Scan(&author.ID, &author.Name, &author.Surname, &author.Bio, &author.Birthday)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Author{}, nil
		}
		return models.Author{}, fmt.Errorf("postgres get author by details: %w", err)
	}
	return author, nil
}

func (p *Postgres) GetAllAuthors(ctx context.Context) ([]models.Author, error) {
	query := `SELECT id, name, surname, bio, birthday FROM authors`
	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("postgres get all authors: %w", err)
	}
	defer rows.Close()

	var authors []models.Author
	for rows.Next() {
		var author models.Author
		if err := rows.Scan(&author.ID, &author.Name, &author.Surname, &author.Bio, &author.Birthday); err != nil {
			return nil, fmt.Errorf("postgres get all authors: %w", err)
		}
		authors = append(authors, author)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("postgres get all authors: %w", err)
	}

	return authors, nil
}

func (p *Postgres) UpdateAuthor(ctx context.Context, author models.Author) (int, error) {
	query := `UPDATE authors SET name = $1, surname = $2, bio = $3, birthday = $4 WHERE id = $5 RETURNING id`
	row := p.db.QueryRowContext(ctx, query, author.Name, author.Surname, author.Bio, author.Birthday, author.ID)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("postgres update author: %w", err)
	}
	return id, nil
}

// Добавление метода DeleteAuthor
func (p *Postgres) DeleteAuthor(ctx context.Context, id int) (int, error) {
	query := `DELETE FROM authors WHERE id = $1 RETURNING id`
	row := p.db.QueryRowContext(ctx, query, id)
	var deletedID int
	err := row.Scan(&deletedID)
	if err != nil {
		return 0, fmt.Errorf("postgres delete author: %w", err)
	}
	return deletedID, nil
}

// Транзакционное обновление данных автора и книги
func (p *Postgres) BeginTx(ctx context.Context) (*sql.Tx, error) {
	return p.db.BeginTx(ctx, nil)
}

func (p *Postgres) CommitTx(tx *sql.Tx) error {
	return tx.Commit()
}

func (p *Postgres) RollbackTx(tx *sql.Tx) error {
	return tx.Rollback()
}

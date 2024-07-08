package repository

import (
	"context"
	"database/sql"

	"library-app/internal/models"
)

type Repository interface {
	//Методы работы с книгами
	CreateBook(ctx context.Context, book models.Book) (int, error)
	GetBookByDetails(ctx context.Context, title string, authorID int, year int, isbn string) (models.Book, error) // Проверка на наличие книги
	GetAuthorIDByName(ctx context.Context, name, surname string) (int, error)                                     // Проверка наличия автора
	GetAllBooks(ctx context.Context) ([]models.Book, error)
	GetBookByID(ctx context.Context, id int) (models.Book, error)
	UpdateBook(ctx context.Context, book models.Book) (int, error)
	DeleteBook(ctx context.Context, id int) (int, error)

	// Методы работы с авторами
	CreateAuthor(ctx context.Context, author models.Author) (int, error)
	GetAuthorByDetails(ctx context.Context, name, surname string) (models.Author, error) // Проверка на наличие автора
	GetAllAuthors(ctx context.Context) ([]models.Author, error)
	GetAuthorByID(ctx context.Context, id int) (models.Author, error)
	UpdateAuthor(ctx context.Context, author models.Author) (int, error)
	DeleteAuthor(ctx context.Context, id int) (int, error)

	// Транзакционное обновление данных автора и книги
	BeginTx(ctx context.Context) (*sql.Tx, error)
	CommitTx(tx *sql.Tx) error
	RollbackTx(tx *sql.Tx) error
}

/*
Для книг:
• POST/books — Добавить новую книгу;
• GET /books — Получить все книги;
• GET /books/{id} — Получить книгу по ее идентификатору;
• PUT /books/{id} — обновить книгу по ее идентификатору;
• DELETE /books/{id} — Удалить книгу по ее идентификатору.
Для авторов:
• POST/authors — Добавить нового автора;
• GET /authors — Получить всех авторов;
• GET /authors/{id} — получить автора по его идентификатору;
• PUT /authors/{id} — обновить автора по его идентификатору;
• DELETE /authors/{id} — удалить автора по его идентификатору.
Транзакционное обновление:
• PUT /books/{book_id
*/

package usecase

import (
	"context"
	"fmt"

	"library-app/internal/models"
	"library-app/internal/repository"
)

type UseCase struct {
	repo repository.Repository
}

func NewUseCase(repo repository.Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) GetOrCreateAuthor(ctx context.Context, author models.Author) (int, error) {
	// Check if the author exists by name and surname
	authorID, err := uc.repo.GetAuthorIDByName(ctx, author.Name, author.Surname)
	if err != nil {
		return 0, fmt.Errorf("usecase get or create author: %w", err)
	}
	if authorID != 0 {
		return authorID, nil
	}

	// If author does not exist, create a new author
	if author.Birthday == "" {
		author.Birthday = "1970-01-01" // Пример даты по умолчанию
	}
	if author.Bio == "" {
		author.Bio = "Default" // Пример значения по умолчанию
	}

	newAuthorID, err := uc.repo.CreateAuthor(ctx, author)
	if err != nil {
		return 0, fmt.Errorf("usecase get or create author: %w", err)
	}
	return newAuthorID, nil
}

func (uc *UseCase) CreateBook(ctx context.Context, title string, authorID int, year int, isbn string) (int, error) {
	book := models.Book{ // Create a new book
		Title:    title,
		AuthorID: authorID,
		Year:     year,
		ISBN:     isbn,
	}
	bookID, err := uc.repo.CreateBook(ctx, book)
	if err != nil {
		return 0, fmt.Errorf("usecase create book: %w", err)
	}
	return bookID, nil
}

func (uc *UseCase) GetBookByDetails(ctx context.Context, title string, authorID int, year int, isbn string) (models.Book, error) {
	book, err := uc.repo.GetBookByDetails(ctx, title, authorID, year, isbn)
	if err != nil {
		return models.Book{}, fmt.Errorf("usecase get book by details: %w", err)
	}
	return book, nil
}

func (uc *UseCase) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	books, err := uc.repo.GetAllBooks(ctx)
	if err != nil {
		return nil, fmt.Errorf("usecase get all books: %w", err)
	}
	return books, nil
}

func (uc *UseCase) GetBookByID(ctx context.Context, id int) (models.Book, error) {
	book, err := uc.repo.GetBookByID(ctx, id)
	if err != nil {
		return models.Book{}, fmt.Errorf("usecase get book by id: %w", err)
	}
	return book, nil
}

func (uc *UseCase) UpdateBook(ctx context.Context, book models.Book) error {
	_, err := uc.repo.UpdateBook(ctx, book)
	if err != nil {
		return fmt.Errorf("usecase update book: %w", err)
	}
	return nil
}

func (uc *UseCase) DeleteBook(ctx context.Context, id int) error {
	_, err := uc.repo.DeleteBook(ctx, id)
	if err != nil {
		return fmt.Errorf("usecase delete book: %w", err)
	}
	return nil
}

// Authors userCases

func (uc *UseCase) CreateAuthor(ctx context.Context, author models.Author) (int, error) {
	authorID, err := uc.repo.CreateAuthor(ctx, author)
	if err != nil {
		return 0, fmt.Errorf("usecase create author: %w", err)
	}
	return authorID, nil
}

func (uc *UseCase) GetAuthorByID(ctx context.Context, id int) (models.Author, error) {
	author, err := uc.repo.GetAuthorByID(ctx, id)
	if err != nil {
		return models.Author{}, fmt.Errorf("usecase get author by id: %w", err)
	}
	return author, nil
}

func (uc *UseCase) GetAuthorByDetails(ctx context.Context, name, surname string) (models.Author, error) {
	author, err := uc.repo.GetAuthorByDetails(ctx, name, surname)
	if err != nil {
		return models.Author{}, fmt.Errorf("usecase get author by details: %w", err)
	}
	return author, nil
}

func (uc *UseCase) GetAllAuthors(ctx context.Context) ([]models.Author, error) {
	authors, err := uc.repo.GetAllAuthors(ctx)
	if err != nil {
		return nil, fmt.Errorf("usecase get all authors: %w", err)
	}
	return authors, nil
}

func (uc *UseCase) UpdateAuthor(ctx context.Context, author models.Author) error {
	_, err := uc.repo.UpdateAuthor(ctx, author)
	if err != nil {
		return fmt.Errorf("usecase update author: %w", err)
	}
	return nil
}

func (uc *UseCase) DeleteAuthor(ctx context.Context, id int) error {
	_, err := uc.repo.DeleteAuthor(ctx, id)
	if err != nil {
		return fmt.Errorf("usecase delete author: %w", err)
	}
	return nil
}

// Транзакционное обновление данных автора и книги
func (uc *UseCase) UpdateBookAndAuthor(ctx context.Context, book models.Book, author models.Author) error {
	tx, err := uc.repo.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("usecase begin transaction: %w", err)
	}

	// Использование контекста для транзакции
	ctxTx := context.WithValue(ctx, "tx", tx)

	if _, err := uc.repo.UpdateBook(ctxTx, book); err != nil {
		uc.repo.RollbackTx(tx)
		return fmt.Errorf("usecase update book: %w", err)
	}

	if _, err := uc.repo.UpdateAuthor(ctxTx, author); err != nil {
		uc.repo.RollbackTx(tx)
		return fmt.Errorf("usecase update author: %w", err)
	}

	if err := uc.repo.CommitTx(tx); err != nil {
		return fmt.Errorf("usecase commit transaction: %w", err)
	}

	return nil
}

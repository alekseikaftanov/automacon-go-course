package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"library-app/internal/models"
	"library-app/internal/usecase"

	"github.com/gorilla/mux"
)

func RegisterBookHandlers(r *mux.Router, uc *usecase.UseCase) {
	r.HandleFunc("/createBook", createBook(uc)).Methods("POST")
	r.HandleFunc("/getBook", getBooks(uc)).Methods("GET")
	r.HandleFunc("/getBookById/{id}", getBookByID(uc)).Methods("GET")
	r.HandleFunc("/updateBook/{id}", updateBook(uc)).Methods("PUT")
	r.HandleFunc("/deleteBook/{id}", deleteBook(uc)).Methods("DELETE")
	r.HandleFunc("/createAuthor", createAuthor(uc)).Methods("POST")
	r.HandleFunc("/getAuthors", getAllAuthors(uc)).Methods("GET")
	r.HandleFunc("/getAuthorById/{id}", getAuthorByID(uc)).Methods("GET")
	r.HandleFunc("/updateAuthor/{id}", updateAuthor(uc)).Methods("PUT")
	r.HandleFunc("/deleteAuthor/{id}", deleteAuthor(uc)).Methods("DELETE")
	r.HandleFunc("/books/{book_id}/authors/{author_id}", updateBookAndAuthor(uc)).Methods("PUT")
}

func createBook(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Title   string `json:"title"`
			Author  string `json:"author"`
			Surname string `json:"surname"`
			Year    int    `json:"year"`
			ISBN    string `json:"isbn"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		authorID, err := uc.GetOrCreateAuthor(r.Context(), models.Author{Name: req.Author, Surname: req.Surname, Bio: "", Birthday: ""})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Проверка на существование книги с такими же деталями
		existingBook, err := uc.GetBookByDetails(r.Context(), req.Title, authorID, req.Year, req.ISBN)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if existingBook.ID != 0 {
			http.Error(w, "Book with the same details already exists", http.StatusConflict)
			return
		}

		bookID, err := uc.CreateBook(r.Context(), req.Title, authorID, req.Year, req.ISBN)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"id": bookID}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func getBooks(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := uc.GetAllBooks(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(books); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func getBookByID(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid book ID", http.StatusBadRequest)
			return
		}

		book, err := uc.GetBookByID(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(book); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func updateBook(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid book ID", http.StatusBadRequest)
			return
		}

		var req struct {
			Title   string `json:"title"`
			Author  string `json:"author"`
			Surname string `json:"surname"`
			Year    int    `json:"year"`
			ISBN    string `json:"isbn"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		book := models.Book{
			ID:       id,
			Title:    req.Title,
			AuthorID: 6, // Замените на подходящий идентификатор автора
			Year:     req.Year,
			ISBN:     req.ISBN,
		}

		err = uc.UpdateBook(r.Context(), book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"id": id}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func deleteBook(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid book ID", http.StatusBadRequest)
			return
		}

		err = uc.DeleteBook(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"deleted": id}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func createAuthor(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Name     string `json:"name"`
			Surname  string `json:"surname"`
			Bio      string `json:"bio"`
			Birthday string `json:"birthday"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Проверка на существование автора с такими же деталями
		existingAuthor, err := uc.GetAuthorByDetails(r.Context(), req.Name, req.Surname)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if existingAuthor.ID != 0 {
			http.Error(w, "Author with the same details already exists", http.StatusConflict)
			return
		}

		authorID, err := uc.CreateAuthor(r.Context(), models.Author{Name: req.Name, Surname: req.Surname, Bio: req.Bio, Birthday: req.Birthday})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"id": authorID}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func getAllAuthors(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authors, err := uc.GetAllAuthors(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(authors); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func getAuthorByID(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid author ID", http.StatusBadRequest)
			return
		}

		author, err := uc.GetAuthorByID(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(author); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func updateAuthor(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid author ID", http.StatusBadRequest)
			return
		}

		var req struct {
			Name     string `json:"name"`
			Surname  string `json:"surname"`
			Bio      string `json:"bio"`
			Birthday string `json:"birthday"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		author := models.Author{
			ID:       id,
			Name:     req.Name,
			Surname:  req.Surname,
			Bio:      req.Bio,
			Birthday: req.Birthday,
		}

		err = uc.UpdateAuthor(r.Context(), author)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"id": id}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("Author updated", author)
	}
}

func deleteAuthor(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid author ID", http.StatusBadRequest)
			return
		}

		err = uc.DeleteAuthor(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(map[string]interface{}{"deleted": id}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// Транзакционное обновление данных автора и книги
func updateBookAndAuthor(uc *usecase.UseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bookID, err := strconv.Atoi(vars["book_id"])
		if err != nil {
			http.Error(w, "Invalid book ID", http.StatusBadRequest)
			return
		}
		authorID, err := strconv.Atoi(vars["author_id"])
		if err != nil {
			http.Error(w, "Invalid author ID", http.StatusBadRequest)
			return
		}

		var req struct {
			BookTitle      string `json:"book_title"`
			BookYear       int    `json:"book_year"`
			BookISBN       string `json:"book_isbn"`
			AuthorName     string `json:"author_name"`
			AuthorSurname  string `json:"author_surname"`
			AuthorBio      string `json:"author_bio"`
			AuthorBirthday string `json:"author_birthday"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		book := models.Book{
			ID:       bookID,
			Title:    req.BookTitle,
			AuthorID: authorID,
			Year:     req.BookYear,
			ISBN:     req.BookISBN,
		}
		author := models.Author{
			ID:       authorID,
			Name:     req.AuthorName,
			Surname:  req.AuthorSurname,
			Bio:      req.AuthorBio,
			Birthday: req.AuthorBirthday,
		}

		if err := uc.UpdateBookAndAuthor(r.Context(), book, author); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("Book and author updated successfully")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

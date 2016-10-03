package repository

import (
	"database/sql"

	"github.com/orderbynull/book-store/core"
	"github.com/orderbynull/book-store/models"
)

type BookRepository struct {
	db *sql.DB
}

//NewBookRepository инициализирует репозиторий
func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

//ByBookshelf получает книги из PSQL по библиотеке
func (s *BookRepository) ByBookshelf(id int) (models.Books, error) {

	//Проверка на правильность параметров
	if id == 0 {
		return nil, core.ErrIncorrectParams
	}

	//Выборка книг по соотв. библиотеке
	rows, err := s.db.Query("SELECT book_id id, title, a.id AS author_id, a.name AS author_name FROM books b JOIN books_bookshelfs bb ON b.id = bb.book_id JOIN authors a ON a.id = b.author_id WHERE bookshelf_id = $1 AND trash = false", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//Формирование результатов в модель
	results := models.Books{}
	for rows.Next() {
		var b models.Book
		var a models.Author

		err = rows.Scan(&b.ID, &b.Title, &a.ID, &a.Name)

		b.Author = a

		results = append(results, b)
	}

	return results, nil
}

//New создает новую книгу в PSQL
func (s *BookRepository) New(title string, authorID, bookshelfID int) (int, error) {

	//Проверка на правильность параметров
	if len(title) == 0 || authorID == 0 || bookshelfID == 0 {
		return 0, core.ErrIncorrectParams
	}

	var lastInsertID int

	//Начало транзакции
	tx, err := s.db.Begin()
	if err != nil {
		return 0, err
	}

	err = tx.QueryRow("INSERT INTO books (title, author_id) VALUES ($1, $2) RETURNING id", title, authorID).Scan(&lastInsertID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	_, err = tx.Exec("INSERT INTO books_bookshelfs (book_id, bookshelf_id) VALUES ($1, $2)", lastInsertID, bookshelfID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	//Конец транзакции
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return 0, err
	}

	return lastInsertID, err

}

//Delete удаляет(выставляет флаг trash=true) книгу в PSQL
func (s *BookRepository) Delete(id int) error {

	//Проверка на правильность параметров
	if id == 0 {
		return core.ErrIncorrectParams
	}

	//Обновление флага trash
	_, err := s.db.Exec("UPDATE books SET trash = true WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

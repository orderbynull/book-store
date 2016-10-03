package repository

import (
	"database/sql"

	"github.com/orderbynull/book-store/models"
)

//BookshelfRepository представляет репозиторий книг
type BookshelfRepository struct {
	db *sql.DB
}

//NewBookshelfRepository инициализирует репозиторий
func NewBookshelfRepository(db *sql.DB) *BookshelfRepository {
	return &BookshelfRepository{db: db}
}

//All возвращает все библиотеки из PSQL
func (s *BookshelfRepository) All() (models.Bookshelfs, error) {
	rows, err := s.db.Query("SELECT id, title FROM bookshelfs")
	if err != nil {
		return nil, err
	}

	results := models.Bookshelfs{}
	for rows.Next() {
		var sh models.Bookshelf

		err = rows.Scan(&sh.ID, &sh.Title)

		results = append(results, sh)
	}

	return results, nil
}

//One возвращает одну конкретную библиотеки из PSQL
func (s *BookshelfRepository) One(id int) (models.Bookshelf, error) {
	row := s.db.QueryRow("SELECT id, title FROM bookshelfs WHERE id = $1", id)

	result := models.Bookshelf{}

	if err := row.Scan(&result.ID, &result.Title); err != nil {
		return result, err
	}

	return result, nil
}

//New создает библиотеку в PSQL
func (s *BookshelfRepository) New(title string) (int, error) {
	if len(title) > 0 {
		var lastInsertID int

		err := s.db.QueryRow("INSERT INTO bookshelfs (title) VALUES ($1) RETURNING id", title).Scan(&lastInsertID)

		return lastInsertID, err
	}

	return 0, sql.ErrNoRows
}

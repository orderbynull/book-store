package repository

import (
	"database/sql"

	"github.com/orderbynull/book-store/models"
)

//AuthorRepository представляет репозиторий авторов
type AuthorRepository struct {
	db *sql.DB
}

//NewAuthorRepository инициализирует репозиторий
func NewAuthorRepository(db *sql.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

//All получает всех авторов из PSQL
func (s *AuthorRepository) All() (models.Authors, error) {
	rows, err := s.db.Query("SELECT id, name FROM authors")
	if err != nil {
		return nil, err
	}

	//Формирование результатов в модель
	results := models.Authors{}
	for rows.Next() {
		var sh models.Author

		err = rows.Scan(&sh.ID, &sh.Name)

		results = append(results, sh)
	}

	return results, nil
}

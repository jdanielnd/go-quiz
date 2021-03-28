package answer

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type UseCase interface {
	GetAll() ([]*Answer, error)
	Store(a *Answer) error
}

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetAll() ([]*Answer, error) {
	var result []*Answer

	rows, err := s.DB.Query("select id, user_id, alternative_id from answers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var a Answer
		err = rows.Scan(&a.ID, &a.UserID, &a.AlternativeID)
		if err != nil {
			return nil, err
		}
		result = append(result, &a)
	}
	return result, nil
}

func (s *Service) Store(a *Answer) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into answers(user_id, alternative_id) values (?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(a.UserID, a.AlternativeID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

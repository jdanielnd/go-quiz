package alternative

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type UseCase interface {
	GetAllFromQuestion(QuestionID int64) ([]*Alternative, error)
	Get(ID int64) (*Alternative, error)
	Store(b *Alternative) error
	Update(b *Alternative) error
	Remove(ID int64) error
}

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetAllFromQuestion(QuestionID int64) ([]*Alternative, error) {
	var result []*Alternative

	stmt, err := s.DB.Prepare("select id, text from alternatives where question_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(QuestionID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a Alternative
		err = rows.Scan(&a.ID, &a.Text)
		if err != nil {
			return nil, err
		}
		result = append(result, &a)
	}
	return result, nil
}

func (s *Service) Get(ID int64) (*Alternative, error) {
	var a Alternative

	stmt, err := s.DB.Prepare("select id, question_id, text from alternatives where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(ID).Scan(&a.ID, &a.QuestionID, &a.Text)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (s *Service) Store(a *Alternative) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into alternatives(question_id, text) values (?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(a.QuestionID, a.Text)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *Service) Update(a *Alternative) error {
	if a.ID == 0 {
		return fmt.Errorf("invalid ID")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("update alternatives set text=?, question_id=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(a.QuestionID, a.Text, a.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *Service) Remove(ID int64) error {
	if ID == 0 {
		return fmt.Errorf("invalid ID")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from alternatives where id=?", ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

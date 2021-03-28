package question

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type UseCase interface {
	GetAll() ([]*Question, error)
	Get(ID int64) (*Question, error)
	Store(b *Question) error
	Update(b *Question) error
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

func (s *Service) GetAll() ([]*Question, error) {
	var result []*Question

	rows, err := s.DB.Query("select id, text, type from questions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var q Question
		err = rows.Scan(&q.ID, &q.Text, &q.Type)
		if err != nil {
			return nil, err
		}
		result = append(result, &q)
	}
	return result, nil
}

func (s *Service) Get(ID int64) (*Question, error) {
	var q Question

	stmt, err := s.DB.Prepare("select id, text, type from questions where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(ID).Scan(&q.ID, &q.Text, &q.Type)
	if err != nil {
		return nil, err
	}
	return &q, nil
}

func (s *Service) Store(q *Question) error {
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into questions(id, text, type) values (?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(q.ID, q.Text, q.Type)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (s *Service) Update(q *Question) error {
	if q.ID == 0 {
		return fmt.Errorf("invalid ID")
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("update questions set text=?, type=? where id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(q.Text, q.Type, q.ID)
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
	_, err = tx.Exec("delete from questions where id=?", ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

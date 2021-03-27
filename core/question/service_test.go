package question_test

import (
	"database/sql"
	"testing"

	"github.com/jdanielnd/go-quiz/core/question"
)

func TestStore(t *testing.T) {
	q := &question.Question{
		ID:   1,
		Text: "What is the package required for running a Go application?",
		Type: question.TypeSingleChoice,
	}
	db, err := sql.Open("sqlite3", "../../data/quiz_test.db")
	if err != nil {
		t.Fatalf("Error connecting to database %s", err.Error())
	}
	err = clearDB(db)
	if err != nil {
		t.Fatalf("Error cleaning database: %s", err.Error())
	}
	defer db.Close()
	service := question.NewService(db)
	err = service.Store(q)
	if err != nil {
		t.Fatalf("Error inserting to database: %s", err.Error())
	}
	saved, err := service.Get(1)
	if err != nil {
		t.Fatalf("Error fetching database: %s", err.Error())
	}
	if saved.ID != 1 {
		t.Fatalf("Invalid data. Expecting %d, returning %d", 1, saved.ID)
	}
}

func clearDB(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from questions")
	tx.Commit()
	return err
}

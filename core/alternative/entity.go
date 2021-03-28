package alternative

type Alternative struct {
	ID         int64  `json:"id"`
	QuestionID int64  `json:"-"`
	Text       string `json:"text"`
}

/*

CREATE TABLE alternatives (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   question_id INTEGER NOT NULL,
   text text NOT NULL,
	 FOREIGN KEY (question_id) REFERENCES questions(id)
);

*/

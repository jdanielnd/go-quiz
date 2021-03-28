package alternative

type Alternative struct {
	ID         int64  `json:"id"`
	QuestionID int64  `json:"-"`
	Text       string `json:"text"`
	Correct    bool   `json:"correct"`
}

/*

CREATE TABLE alternatives (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   question_id INTEGER NOT NULL,
   text text NOT NULL,
	 correct boolean NOT NULL,
	 FOREIGN KEY (question_id) REFERENCES questions(id)
);

*/

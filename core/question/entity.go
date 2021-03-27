package question

type Question struct {
	ID   int64        `json:"id"`
	Text string       `json:"text"`
	Type QuestionType `json:"type"`
}

/*

CREATE TABLE questions (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   text text NOT NULL,
   type integer NOT NULL
);

*/

type QuestionType int

const (
	TypeSingleChoice   = 1
	TypeMultipleChoice = 2
	TypeTrueFalse      = 3
)

func (t QuestionType) String() string {
	switch t {
	case TypeSingleChoice:
		return "Single Choice"
	case TypeMultipleChoice:
		return "Multiple Choice"
	case TypeTrueFalse:
		return "True or False"
	}
	return "Unknown"
}

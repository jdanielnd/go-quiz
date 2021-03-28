package answer

type Answer struct {
	ID            int64 `json:"id"`
	UserID        int64 `json:"user_id"`
	AlternativeID int64 `json:"alternative_id"`
}

/*

CREATE TABLE answers (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   user_id INTEGER NOT NULL,
	 alternative_id INTEGER NOT NULL,
	 FOREIGN KEY (alternative_id) REFERENCES alternatives(id)
);

*/

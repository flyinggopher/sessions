package sessions

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	filepath string
	base     *sql.DB
}

func (s *Storage) Connect() error {
	database, err := sql.Open("sqlite3", s.filepath)

	if err != nil {
		return err
	}

	s.base = database

	_, err = s.base.Exec("CREATE TABLE IF NOT EXISTS sessions (sessionid UNSIGNED BIG INTEGER PRIMARY KEY, userid INTEGER, createdat varchar(255), deadline DATETIME);")

	if err != nil {
		return err
	}

	return nil
}

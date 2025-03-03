package gateway

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func initDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		useraddress TEXT NOT NULL,
		username TEXT NOT NULL,
		cid TEXT NOT NULL,
		filename TEXT NOT NULL
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	return db, nil
}

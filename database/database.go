package database

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type LiveEvent struct {
	ID          int
	Title       string
	Description string
	StartTime   time.Time
	EndTime     *time.Time
	Rewards     string
}

type DB struct {
	conn *sql.DB
}

// InitDB initializes the SQLite database if it doesn't exist,
func InitDB(filepath string) (DB, error) {
	conn, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return DB{}, err
	}

	// SQL statement to create the table if it doesn't exist
	sqlStmt := `
		CREATE TABLE IF NOT EXISTS live_events (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			description TEXT,
			start_time INTEGER NOT NULL,
			end_time INTEGER,
			rewards TEXT
		);`

	_, err = conn.Exec(sqlStmt)
	if err != nil {
		return DB{}, err
	}

	return DB{conn: conn}, nil
}

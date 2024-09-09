package db

import "database/sql"

type Database struct {
	db *sql.DB
}

func NewDb() *Database {
	db, err := initDb()
	if err != nil {
		panic(err)
	}

	return &Database{db}
}

func initDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./webapp.db")
	if err != nil {
		return nil, err
	}

	// Create the users table if it doesn't exist
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS shorten_urls (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            shorten_url TEXT NOT NULL,
            original_url TEXT NOT NULL,
            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
        )
    `)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *Database) Close() {
	db.db.Close()
}

func (db *Database) GetShortenUrl(shortenUrl string) (*ShortenUrl, error) {
	var shortenUrls []ShortenUrl
	err := db.db.QueryRow("SELECT * FROM shorten_urls WHERE shorten_url = ?", shortenUrl).Scan(&shortenUrls)
	if err != nil {
		return nil, err
	}
	return &shortenUrls[0], nil
}

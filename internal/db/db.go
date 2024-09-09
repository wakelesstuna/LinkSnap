package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

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
	db, err := sql.Open("sqlite3", "./app.db")
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

	rows, err := db.db.Query("SELECT * FROM shorten_urls WHERE shorten_url = ?", shortenUrl)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shortenUrls []ShortenUrl

	for rows.Next() {
		var su ShortenUrl
		err := rows.Scan(&su.Id, &su.ShortenUrl, &su.OriginalUrl, &su.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		shortenUrls = append(shortenUrls, su)
	}
	return &shortenUrls[0], nil
}

func (db *Database) AddShortenUrl(shortenUrl string, originalUrl string) (*ShortenUrl, error) {
	query := `INSERT INTO shorten_urls VALUES (null, ?, ?, ?)`
	r, err := db.db.Exec(query, shortenUrl, originalUrl, time.Now())
	if err != nil {
		return nil, err
	}
	lastId, err := r.LastInsertId()
	if err != nil {
		return nil, err
	}
	var sh ShortenUrl
	err = db.db.QueryRow("SELECT * FROM shorten_urls WHERE id = ?", lastId).Scan(
		&sh.Id,
		&sh.ShortenUrl,
		&sh.OriginalUrl,
		&sh.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &sh, nil
}

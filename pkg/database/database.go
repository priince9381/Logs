package database

import (
	"assisment/internal/models"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Database interface {
	storeLogInDatabase(log *models.Log) error
	Close() error
}

type DatabaseImpl struct {
	db *sql.DB
}

func InitializeDB() (*DatabaseImpl, error) {
	maxRetries := 10
	retryInterval := 3 * time.Second
	// Todo To create a secure method to fetch secrets
	connStr := "postgres://newuser:password@database:5432/log?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Ping the database to ensure a successful connection
	for i := 0; i < maxRetries; i++ {
		err = db.Ping()
		if err == nil {
			// Connection successful, return the db instance
			return &DatabaseImpl{db: db}, nil
		}

		log.Printf("Failed to connect to the database. Retrying in %v...", retryInterval)
		time.Sleep(retryInterval)
	}

	return nil, fmt.Errorf("failed to connect to the database after multiple retries")
}

// Todo To Create a generic Method to handle database query
func (d *DatabaseImpl) StoreLogInDatabase(log *models.Log) error {
	// Insert the log into the database
	_, err := d.db.Exec(`
		INSERT INTO logs (request_id, unix_ts, user_id, event_name)
		VALUES ($1, $2, $3, $4)
	`, log.GetId(), log.GetUnixTs(), log.GetUserId(), log.GetEventName())
	if err != nil {
		return err
	}

	return nil
}

func (d *DatabaseImpl) CreateLogTable() error {
	_, err := d.db.Exec(`
		CREATE TABLE IF NOT EXISTS logs (
		    id SERIAL PRIMARY KEY,
		    request_id INT,
			user_id INT,
			event_name TEXT,
			unix_ts TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}

	return nil
}

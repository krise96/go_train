package repositories

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
	db     *sql.DB
}

// Constructor
func CreateStore(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// Public
func (store *Store) Open() error {
	db, err := sql.Open("postgres", store.config.DBURL)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	store.db = db

	return nil
}

func (store *Store) Close() error {
	if err := store.db.Close(); err != nil {
		return err
	}

	return nil
}

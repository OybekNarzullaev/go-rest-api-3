package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// store ...
type Store struct {
	config *Config
	db *sql.DB
	userRepositry *UserRepositry
}

// constructor
func New(config *Config) *Store {
	return &Store{config: config}
}

// open...
func (s *Store) Open() error{
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// close...
func (s *Store) Close(){
	s.db.Close()
}

// user...
func (s *Store) User() *UserRepositry {
	if s.userRepositry != nil {
		return s.userRepositry
	}
	s.userRepositry = &UserRepositry{
		store: s,
	}

	return s.userRepositry
}
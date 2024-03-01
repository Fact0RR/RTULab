package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	Connection      *string
	DB              *sql.DB
	TokenRefreshMap map[string]string
	TokenAccessMap  map[string]string
}

func New(connection string) *Store {
	tokenRefreshMap :=  make(map[string]string)
	tokenAccessMap :=  make(map[string]string)

	store := Store{
		Connection: &connection,
		TokenRefreshMap: tokenRefreshMap,
		TokenAccessMap: tokenAccessMap,
	}
	return &store
}

func (s *Store) Open() error {

	db, err := sql.Open("postgres", *s.Connection)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}

	s.DB = db

	return nil
}

func (s *Store) Close() {
	s.DB.Close()
}

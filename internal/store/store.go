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
	tokenRefreshMap := make(map[string]string)
	tokenAccessMap := make(map[string]string)

	store := Store{
		Connection:      &connection,
		TokenRefreshMap: tokenRefreshMap,
		TokenAccessMap:  tokenAccessMap,
	}
	return &store
}

func (s *Store) Open(k, j int) error {

	db, err := sql.Open("postgres", *s.Connection)
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE IF EXISTS constants")
	if err != nil {
		return err
	}
	_, err = db.Exec("create table IF NOT EXISTS constants(id varchar(15) primary key, value int)")
	if err != nil {
		return err
	}

	_, err = db.Exec("insert into constants values($1,$2),($3,$4);","k",k,"j",j)
	if err != nil {
		return err
	}

	s.DB = db

	return nil
}

func (s *Store) Close() {
	s.DB.Close()
}

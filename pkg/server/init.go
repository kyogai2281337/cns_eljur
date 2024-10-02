package server

import (
	"database/sql"

	"github.com/kyogai2281337/cns_eljur/pkg/sql/store/sqlstore"
)

func Start(config *Config) error {
	db, err := NewDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()
	store := sqlstore.New(db)
	s := NewServer(store, config.BrokerURL)

	return s.ServeHTTP(config.BindAddr)
}

func NewDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

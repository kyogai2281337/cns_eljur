package apiserver

import (
	"database/sql"
	"net/http"
	"github.com/kyogai2281337/cns_eljur/internal/auth/store/sqlstore"

	"github.com/gorilla/sessions"
)

func Start(config *Config) error {
	db, err := NewDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	store := sqlstore.New(db)
	sessStore := sessions.NewCookieStore([]byte(config.SessionKey))
	s := NewServer(store, sessStore)

	return http.ListenAndServe(config.BindAddr, s)
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

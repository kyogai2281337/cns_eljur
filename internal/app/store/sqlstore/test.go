package sqlstore

import (
	"database/sql"
	"sync"
	"testing"
)

// Testing DB
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
	return db, func(tables ...string) {
		if len(tables) > 0 {
			var wg sync.WaitGroup
			for _, el := range tables {
				wg.Add(1)
				go func(table string) {
					defer wg.Done()
					_, err := db.Exec("TRUNCATE TABLE " + table)
					if err != nil {
						t.Errorf("Error truncating table %s: %v", table, err)
					}
				}(el)
			}
			wg.Wait()
		}
		db.Close()
	}
}

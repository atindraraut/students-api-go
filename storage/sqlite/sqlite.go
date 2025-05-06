package sqlite

import (
	"database/sql"
	"log/slog"
	_ "github.com/mattn/go-sqlite3"
	"github.com/atindraraut/crudgo/internal/config"
)


type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	slog.Info("Connecting to database...", slog.String("path", cfg.Storagepath))
	db,err:=sql.Open("sqlite3", cfg.Storagepath)
	if err != nil {
		return nil, err	
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER NOT NULL,
		email TEXT NOT NULL UNIQUE
	)`)
	if err != nil {
		return nil, err
	}
	return &Sqlite{Db: db}, nil
}
package db

import (
	"database/sql"
	"embed"
	"os"
	"path/filepath"
	"strings"

	_ "modernc.org/sqlite"
)

const appName = "life"
const dbFile = "life.db"

//go:embed migrations/*.sql
var migrations embed.FS

func runMigrations(db *sql.DB) error {
	entries, err := migrations.ReadDir("migrations")
	if err != nil {
		return err
	}

	for _, entry := range entries {
		content, err := migrations.ReadFile("migrations/" + entry.Name())
		if err != nil {
			return err
		}

		// Execute statements individually and ignore "duplicate column" / "already exists" errors
		// to make migrations idempotent when re-running on the same DB.
		stmts := strings.Split(string(content), ";")
		for _, s := range stmts {
			stmt := strings.TrimSpace(s)
			if stmt == "" {
				continue
			}
			if _, err := db.Exec(stmt); err != nil {
				if err != nil && (strings.Contains(err.Error(), "duplicate column name") || strings.Contains(err.Error(), "already exists") || strings.Contains(err.Error(), "table .* already exists")) {
					// ignore and continue
					continue
				}
				return err
			}
		}
	}

	return nil
}

func dbPath() (string, error) {
	dataDir, ok := os.LookupEnv("XDG_DATA_HOME")
	if !ok {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		dataDir = filepath.Join(home, ".local", "share")
	}
	return filepath.Join(dataDir, appName, dbFile), nil
}

func OpenConn() (*sql.DB, error) {
	path, err := dbPath()
	if err != nil {
		return nil, err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := runMigrations(db); err != nil {
		return nil, err
	}

	return db, nil
}

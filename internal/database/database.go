package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// DB adalah tipe untuk representasi koneksi database
type DB struct {
	*sql.DB
}

// InitDatabase membuka koneksi ke database dan mengembalikan objek DB
func InitDatabase(dbPath string) (*sql.DB, error) {
	if dbPath == "" {
		return nil, fmt.Errorf("database path is empty")
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	return db, nil
}

// GetSuffixesByPrefix mencari suffix berdasarkan prefix
func GetSuffixesByPrefix(db *DB, prefix string) ([]string, error) {
	query := `SELECT suffix FROM hibp WHERE prefix = ?`
	rows, err := db.Query(query, prefix)
	if err != nil {
		return nil, fmt.Errorf("failed to query suffixes: %v", err)
	}
	defer rows.Close()

	var suffixes []string
	for rows.Next() {
		var suffix string
		if err := rows.Scan(&suffix); err != nil {
			return nil, fmt.Errorf("failed to scan suffix: %v", err)
		}
		suffixes = append(suffixes, suffix)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	return suffixes, nil
}

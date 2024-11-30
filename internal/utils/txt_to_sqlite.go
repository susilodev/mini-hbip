package main

import (
	"bufio"
	"database/sql"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Load konfigurasi dari environment
	// txtFile := os.Getenv("TXT_FILE_PATH")
	// dbFile := os.Getenv("DATABASE_PATH")

	// Buka file .txt
	file, err := os.Open("../../data/hibp_example.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Buka koneksi ke SQLite
	db, err := sql.Open("sqlite3", "../../data/hbip.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Buat tabel jika belum ada
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS pwned_passwords (
		prefix TEXT NOT NULL,
		suffix TEXT NOT NULL,
		occurrences INTEGER NOT NULL,
		PRIMARY KEY (prefix, suffix)
	)`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// Masukkan data ke dalam tabel
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Failed to begin transaction: %v", err)
	}
	stmt, err := tx.Prepare("INSERT OR IGNORE INTO pwned_passwords (prefix, suffix, occurrences) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatalf("Failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		hash := parts[0]
		count := parts[1]

		prefix := hash[:5]
		suffix := hash[5:]

		_, err := stmt.Exec(prefix, suffix, count)
		if err != nil {
			log.Printf("Failed to insert record: %v", err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}

	log.Println("Data conversion completed.")
}

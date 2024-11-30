package handler

import (
	"database/sql"
	"fmt"
	"net/http"
)

// query handler
func CheckHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// fetch prefix from url query
		prefix := r.URL.Query().Get("prefix")
		if len(prefix) != 5 {
			http.Error(w, "Prefix must be exactly 5 characters", http.StatusBadRequest)
			return
		}

		rows, err := db.Query("SELECT suffix, COUNT(*) FROM pwned_passwords WHERE prefix = ? GROUP BY suffix", prefix)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		defer rows.Close()

		var result string
		for rows.Next() {
			var suffix string
			var occurrences int
			if err := rows.Scan(&suffix, &occurrences); err != nil {
				http.Error(w, "Error scanning result", http.StatusInternalServerError)
				return
			}

			result += fmt.Sprintf("%s:%d\n", suffix, occurrences)
		}

		if result == "" {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "No suffix found")
			return
		}

		// returning text format as a response
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, result)
	}
}

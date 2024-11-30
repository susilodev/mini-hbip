package config

import "os"

// LoadConfig loads configuration variables from environment
func LoadConfig() map[string]string {
	return map[string]string{
		"DB_HOST": os.Getenv("DB_HOST"),
		"DB_PORT": os.Getenv("DB_PORT"),
	}
}

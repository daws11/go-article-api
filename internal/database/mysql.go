package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url" // Untuk parsing URL
	"os"
	"strings" // Untuk DSN formatting jika perlu

	_ "github.com/go-sql-driver/mysql" // Driver MySQL
)

// Fungsi untuk membuat DSN dari URL atau variabel individual
func getDSN() (string, error) {
	// Prioritaskan DATABASE_URL dari Heroku Add-on
	dbURL := os.Getenv("DATABASE_URL") // Atau CLEARDB_DATABASE_URL, JAWSDB_URL, dll.
	if dbURL != "" {
		log.Println("Using DATABASE_URL environment variable.")
		// Heroku add-on biasanya formatnya: mysql://user:pass@host:port/dbname?options
		// Driver Go MySQL butuh format: user:pass@tcp(host:port)/dbname?options

		parsedURL, err := url.Parse(dbURL)
		if err != nil {
			return "", fmt.Errorf("failed to parse DATABASE_URL: %w", err)
		}

		user := parsedURL.User.Username()
		pass, _ := parsedURL.User.Password()
		host := parsedURL.Hostname()
		port := parsedURL.Port()
		dbName := strings.TrimPrefix(parsedURL.Path, "/") // Hapus '/' di awal path

		// Tambahkan port default jika tidak ada
		if port == "" {
			port = "3306"
		}

		// Gabungkan query options dari URL asli
		queryOptions := parsedURL.RawQuery
		if queryOptions != "" {
			queryOptions = "?" + queryOptions
		}
		// Tambahkan parseTime=true jika belum ada di options
		if !strings.Contains(queryOptions, "parseTime=true") {
			if queryOptions == "" {
				queryOptions = "?parseTime=true"
			} else {
				queryOptions += "&parseTime=true"
			}
		}

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", user, pass, host, port, dbName, queryOptions)
		return dsn, nil

	} else {
		// Fallback ke variabel individual (untuk lokal)
		log.Println("DATABASE_URL not found, using individual DB environment variables (for local).")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		if dbUser == "" || dbHost == "" || dbPort == "" || dbName == "" {
			return "", fmt.Errorf("required DB environment variables (DB_USER, DB_HOST, DB_PORT, DB_NAME) are missing")
		}

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
		return dsn, nil
	}
}

func InitDB() (*sql.DB, error) {
	// Hapus load .env jika tidak dipakai di Heroku
	// err := godotenv.Load()
	// if err != nil { ... }

	dsn, err := getDSN()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection established")
	return db, nil
}

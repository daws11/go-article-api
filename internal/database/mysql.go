package database

import (
	"database/sql"
	"fmt"
	"log"
	"os" // Atau ganti dengan viper/envconfig jika lebih kompleks

	_ "github.com/go-sql-driver/mysql" // Blank import for driver registration
	"github.com/joho/godotenv"         // Jika pakai godotenv
)

func InitDB() (*sql.DB, error) {
	// Load .env file (jika pakai godotenv)
	err := godotenv.Load() // Cari file .env di root project
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables directly")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Format: user:password@tcp(host:port)/dbname?parseTime=true
	// parseTime=true penting agar bisa scan tipe TIMESTAMP/DATETIME ke time.Time
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Cek koneksi
	err = db.Ping()
	if err != nil {
		db.Close() // Tutup jika ping gagal
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection established")
	return db, nil
}

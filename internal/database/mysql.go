package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url" // Untuk parsing URL
	"os"
	"strings" // Untuk manipulasi string dasar

	_ "github.com/go-sql-driver/mysql" // Driver MySQL
)

// Fungsi untuk membuat DSN dari URL atau variabel individual
func getDSN() (string, error) {
	// Prioritaskan DATABASE_URL dari Heroku Add-on
	dbURL := os.Getenv("DATABASE_URL")
	log.Printf("DEBUG: Read DATABASE_URL: [%s]\n", dbURL) // Log nilai mentah

	if dbURL != "" {
		log.Println("DEBUG: Using DATABASE_URL environment variable.")

		parsedURL, err := url.Parse(dbURL)
		if err != nil {
			log.Printf("ERROR: Failed to parse DATABASE_URL: %v\n", err)
			return "", fmt.Errorf("failed to parse DATABASE_URL: %w", err)
		}

		// Ekstrak komponen URL
		user := parsedURL.User.Username()
		pass, _ := parsedURL.User.Password()
		host := parsedURL.Hostname()
		port := parsedURL.Port()
		dbName := strings.TrimPrefix(parsedURL.Path, "/")

		if port == "" {
			port = "3306" // Default MySQL port
		}

		// --- Penanganan Query Parameters (Termasuk SSL/TLS) ---
		// Ambil query parameters asli
		originalParams := parsedURL.Query()
		// Buat map baru untuk parameter DSN akhir
		dsnParams := make(url.Values)

		// 1. Selalu tambahkan parseTime=true
		dsnParams.Set("parseTime", "true")

		// 2. Cek parameter SSL/TLS dari URL asli
		useSSL := originalParams.Get("useSSL") == "true"
		requireSSL := originalParams.Get("requireSSL") == "true"

		// 3. Set parameter 'tls' untuk driver Go berdasarkan parameter asli
		if useSSL || requireSSL {
			// Gunakan tls=true. Jika ada masalah verifikasi sertifikat di Stackhero,
			// Anda MUNGKIN perlu mencoba 'skip-verify' TAPI ini kurang aman.
			// Coba 'true' dulu.
			dsnParams.Set("tls", "true")
			log.Println("DEBUG: Setting DSN parameter 'tls=true' based on original URL options.")
		} else {
			// Jika tidak ada useSSL/requireSSL, mungkin tidak perlu parameter tls
			// atau Anda bisa set ke false jika driver memerlukannya
			// dsnParams.Set("tls", "false") // Biasanya tidak perlu
			log.Println("DEBUG: No SSL/TLS parameters detected in original URL.")
		}

		// 4. (Opsional) Salin parameter lain yang relevan jika ada
		// for key, values := range originalParams {
		//     if key != "useSSL" && key != "requireSSL" && key != "parseTime" && key != "tls" {
		//          // Salin parameter lain jika perlu
		//          for _, value := range values {
		//              dsnParams.Add(key, value)
		//          }
		//     }
		// }

		// Encode parameter menjadi query string
		finalQueryString := dsnParams.Encode()
		if finalQueryString != "" {
			finalQueryString = "?" + finalQueryString
		}
		// ------------------------------------------------------

		// Buat DSN final dalam format yang benar
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", user, pass, host, port, dbName, finalQueryString)
		log.Printf("DEBUG: Generated DSN: %s\n", dsn) // Log DSN final
		return dsn, nil

	} else {
		// Fallback ke variabel individual (untuk lokal) - Seharusnya tidak terjadi di Heroku
		log.Println("WARNING: DATABASE_URL not found, using individual DB environment variables (for local). This should not happen on Heroku!")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		if dbUser == "" || dbHost == "" || dbPort == "" || dbName == "" {
			return "", fmt.Errorf("DATABASE_URL not found and required local DB environment variables (DB_USER, DB_HOST, DB_PORT, DB_NAME) are missing")
		}

		// DSN untuk lokal (asumsi tanpa SSL/TLS, tambahkan jika perlu)
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
		log.Printf("DEBUG: Generated DSN from local vars: %s\n", dsn)
		return dsn, nil
	}
}

func InitDB() (*sql.DB, error) {
	// Hapus atau komentari load .env jika tidak dipakai/dibutuhkan lagi
	// err := godotenv.Load()
	// if err != nil { log.Println("Warning: .env file not found") }

	dsn, err := getDSN()
	if err != nil {
		// Log error dari getDSN sebelum return
		log.Printf("ERROR: Failed to get DSN: %v\n", err)
		return nil, fmt.Errorf("failed to configure database DSN: %w", err)
	}

	// Coba buka koneksi
	log.Println("Attempting to open database connection...")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("ERROR: sql.Open failed: %v\n", err)
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Coba ping database
	log.Println("Pinging database...")
	err = db.Ping()
	if err != nil {
		log.Printf("ERROR: db.Ping failed: %v\n", err)
		db.Close() // Tutup koneksi jika ping gagal
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection established successfully!")
	return db, nil
}

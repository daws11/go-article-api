// File: cmd/server/main.go

package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	// ===> 1. IMPORT Paket CORS <===
	"github.com/gin-contrib/cors"

	// Import internal packages (sesuaikan path jika beda)
	"github.com/daws11/go-article-api/internal/article"
	"github.com/daws11/go-article-api/internal/database"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// 1. Init Database Connection
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Could not initialize database connection: %v", err)
	}
	defer db.Close() // Pastikan koneksi ditutup saat main selesai

	// 2. Init Dependencies (Repo, Service, Handler)
	articleRepo := article.NewMySQLRepository(db)
	articleService := article.NewService(articleRepo)
	articleHandler := article.NewHandler(articleService)

	// 3. Init Gin Router
	router := gin.Default() // Logger dan Recovery sudah termasuk

	// ===> 2. KONFIGURASI dan GUNAKAN Middleware CORS <===
	// Konfigurasi CORS dasar untuk pengembangan lokal
	// Izinkan origin frontend Anda (misal: http://localhost:5173)
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                             // Origin frontend React Anda
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},  // Metode HTTP yang diizinkan
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"}, // Header yang diizinkan
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Izinkan jika Anda perlu mengirim cookie/auth header
		MaxAge:           12 * time.Hour,
	}
	// Gunakan middleware CORS secara global
	router.Use(cors.New(config))
	// ======================================================

	// Middleware lain jika perlu (sudah ada Logger & Recovery dari gin.Default())
	// router.Use(gin.Logger())   // Sudah ada di gin.Default()
	// router.Use(gin.Recovery()) // Sudah ada di gin.Default()

	// 4. Define Routes
	api := router.Group("/article") // Base URL sudah di set di frontend axios config
	{
		api.POST("/", articleHandler.CreateArticleHandler)
		// GUNAKAN ROUTE YANG SESUAI DENGAN API CLIENT ANDA SEKARANG
		// Jika client pakai /:limit/:offset
		// api.GET("/:limit/:offset", articleHandler.GetArticlesHandler)
		// Jika client sudah diubah ke query param ?limit=..&offset=..
		api.GET("/", articleHandler.GetArticlesHandler)

		api.GET("/:id", articleHandler.GetArticleByIDHandler)
		api.PUT("/:id", articleHandler.UpdateArticleHandler)
		api.PATCH("/:id", articleHandler.UpdateArticleHandler) // Handle PATCH juga (opsional)
		api.DELETE("/:id", articleHandler.DeleteArticleHandler)
	}

	// Route tambahan untuk health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	// 5. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port backend
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  15 * time.Second, // Naikkan sedikit jika perlu
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Server starting on port %s", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", port, err)
	}
}

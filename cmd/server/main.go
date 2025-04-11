package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/daws11/go-article-api/internal/article"
	"github.com/daws11/go-article-api/internal/database"
	"github.com/gin-gonic/gin"
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
	router := gin.Default()
	// Tambahkan middleware jika perlu (CORS, logging, recovery, etc)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// 4. Define Routes
	api := router.Group("/article")
	{
		api.POST("/", articleHandler.CreateArticleHandler)
		api.GET("/", articleHandler.GetArticlesHandler)
		api.GET("/:id", articleHandler.GetArticleByIDHandler)
		api.PUT("/:id", articleHandler.UpdateArticleHandler)
		api.PATCH("/:id", articleHandler.UpdateArticleHandler)
		api.DELETE("/:id", articleHandler.DeleteArticleHandler)
	}

	// Route tambahan untuk health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	// 5. Start Server
	port := os.Getenv("PORT") // Ambil port dari env, atau default
	if port == "" {
		port = "8080" // Default port
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Printf("Server starting on port %s", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", port, err)
	}
}

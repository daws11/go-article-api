package article

import (
	"database/sql"
	"errors" // Untuk cek error spesifik jika perlu
	"fmt"
	"log"
	"net/http"
	"strconv" // Untuk konversi string ID dari URL ke int

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// POST /article/
func (h *Handler) CreateArticleHandler(c *gin.Context) {
	var input ArticleInput
	// Bind JSON body ke struct input
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	// Panggil service
	createdArticle, err := h.service.CreateArticle(input)
	if err != nil {
		log.Printf("Error creating article: %v", err)
		// Cek jenis error untuk response yang lebih baik
		// Contoh: Jika error validasi
		// if validationErrors, ok := err.(validator.ValidationErrors); ok {
		//     c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": validationErrors.Error()})
		//     return
		// }
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article: " + err.Error()})
		return
	}

	// Sukses (Sesuai dokumen response {}, tapi lebih baik kembalikan object yg dibuat)
	// c.JSON(http.StatusCreated, gin.H{})
	c.JSON(http.StatusCreated, createdArticle) // Mengembalikan objek yang baru dibuat
}

// GET /article/:limit/:offset
func (h *Handler) GetArticlesHandler(c *gin.Context) {
	limitStr := c.Param("limit")
	offsetStr := c.Param("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10 // Default limit jika parsing gagal atau invalid
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0 // Default offset jika parsing gagal atau invalid
	}

	articles, err := h.service.GetArticles(limit, offset)
	if err != nil {
		log.Printf("Error getting articles: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve articles: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, articles)
}

// GET /article/:id
func (h *Handler) GetArticleByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64) // Parse ke int64
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID format"})
		return
	}

	article, err := h.service.GetArticleByID(id)
	if err != nil {
		log.Printf("Error getting article by ID %d: %v", id, err)
		// Cek jika error karena tidak ditemukan
		if errors.Is(err, sql.ErrNoRows) || err.Error() == fmt.Sprintf("article with id %d not found", id) { // Cek error spesifik repo
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Article with ID %d not found", id)})
			return
		}
		if err.Error() == fmt.Sprintf("invalid article ID: %d", id) { // Cek error service
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve article: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, article)
}

// PUT atau PATCH /article/:id (Menggunakan PUT sebagai contoh)
func (h *Handler) UpdateArticleHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID format"})
		return
	}

	var input ArticleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Error binding JSON for update: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	updatedArticle, err := h.service.UpdateArticle(id, input)
	if err != nil {
		log.Printf("Error updating article %d: %v", id, err)
		// Handle not found error
		if err.Error() == fmt.Sprintf("article with id %d not found for update", id) || err.Error() == fmt.Sprintf("article with id %d not found", id) {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Article with ID %d not found", id)})
			return
		}
		// Handle validation error
		// if validationErrors, ok := err.(validator.ValidationErrors); ok { ... }
		if err.Error() == fmt.Sprintf("invalid article ID: %d", id) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update article: " + err.Error()})
		return
	}

	// Sesuai dokumen {}
	// c.JSON(http.StatusOK, gin.H{})
	c.JSON(http.StatusOK, updatedArticle) // Kembalikan yang sudah diupdate
}

// DELETE /article/:id
func (h *Handler) DeleteArticleHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID format"})
		return
	}

	err = h.service.DeleteArticle(id)
	if err != nil {
		log.Printf("Error deleting article %d: %v", id, err)
		// Handle not found error
		if err.Error() == fmt.Sprintf("article with id %d not found for delete", id) {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Article with ID %d not found", id)})
			return
		}
		if err.Error() == fmt.Sprintf("invalid article ID: %d", id) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article ID"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete article: " + err.Error()})
		return
	}

	// Sukses - No Content atau pesan sukses
	// c.Status(http.StatusNoContent)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Article with ID %d deleted successfully", id)})
}

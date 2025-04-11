package article

import (
	"database/sql"
	"fmt"
	"log"
)

// Interface untuk decoupling
type Repository interface {
	CreateArticle(article ArticleInput) (int64, error)
	// GetArticles(limit, offset int) ([]Article, error) // pagination server side
	GetArticles() ([]Article, error)
	GetArticleByID(id int64) (Article, error)
	UpdateArticle(id int64, article ArticleInput) error
	DeleteArticle(id int64) error
}

type mysqlRepository struct {
	db *sql.DB
}

// Constructor untuk repository
func NewMySQLRepository(db *sql.DB) Repository {
	return &mysqlRepository{db: db}
}

func (r *mysqlRepository) CreateArticle(article ArticleInput) (int64, error) {
	query := "INSERT INTO posts (title, content, category, status) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, article.Title, article.Content, article.Category, article.Status)
	if err != nil {
		return 0, fmt.Errorf("error creating article: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert ID: %w", err)
	}
	return id, nil
}

func (r *mysqlRepository) GetArticles() ([]Article, error) { // Hapus limit, offset dari signature
	// Query tanpa LIMIT dan OFFSET. Pertimbangkan ORDER BY untuk konsistensi.
	query := "SELECT id, title, content, category, created_date, updated_date, status FROM posts ORDER BY created_date DESC"

	// Eksekusi query tanpa parameter limit/offset
	rows, err := r.db.Query(query)
	if err != nil {
		// Kembalikan error jika query gagal
		return nil, fmt.Errorf("error querying all articles: %w", err)
	}
	// Pastikan rows ditutup setelah selesai
	defer rows.Close()

	articles := []Article{} // Inisialisasi slice kosong
	for rows.Next() {
		var art Article
		// Scan data dari baris ke struct Article
		errScan := rows.Scan(&art.ID, &art.Title, &art.Content, &art.Category, &art.CreatedDate, &art.UpdatedDate, &art.Status)
		if errScan != nil {
			// Log error scan tapi lanjutkan ke baris berikutnya (atau hentikan sesuai kebutuhan)
			log.Printf("Warning: error scanning article row: %v", errScan)
			// Jika ingin berhenti saat ada error scan:
			// return nil, fmt.Errorf("error scanning article row: %w", errScan)
			continue // Lanjutkan ke baris berikutnya
		}
		// Tambahkan artikel yang berhasil di-scan ke slice
		articles = append(articles, art)
	}

	// Periksa error yang mungkin terjadi selama iterasi rows
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating article rows: %w", err)
	}

	// Kembalikan slice (bisa kosong jika tabel kosong atau semua baris error scan)
	return articles, nil
}

func (r *mysqlRepository) GetArticleByID(id int64) (Article, error) {
	query := "SELECT id, title, content, category, created_date, updated_date, status FROM posts WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var art Article
	err := row.Scan(&art.ID, &art.Title, &art.Content, &art.Category, &art.CreatedDate, &art.UpdatedDate, &art.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return Article{}, fmt.Errorf("article with id %d not found", id) // Specific error for not found
		}
		return Article{}, fmt.Errorf("error scanning article: %w", err)
	}
	return art, nil
}

func (r *mysqlRepository) UpdateArticle(id int64, article ArticleInput) error {
	query := "UPDATE posts SET title = ?, content = ?, category = ?, status = ? WHERE id = ?"
	result, err := r.db.Exec(query, article.Title, article.Content, article.Category, article.Status, id)
	if err != nil {
		return fmt.Errorf("error updating article: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("article with id %d not found for update", id) // Or handle as success if idempotent PUT
	}
	return nil
}

func (r *mysqlRepository) DeleteArticle(id int64) error {
	query := "DELETE FROM posts WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting article: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("article with id %d not found for delete", id)
	}
	return nil
}

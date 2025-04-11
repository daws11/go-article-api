package article

import (
	"fmt"
	"log" // Tambahkan import log jika belum ada

	"github.com/go-playground/validator/v10"
)

// Interface Service (Sudah benar di kode Anda)
type Service interface {
	CreateArticle(input ArticleInput) (Article, error)
	GetArticles() ([]Article, error)
	GetArticleByID(id int64) (Article, error)
	UpdateArticle(id int64, input ArticleInput) (Article, error)
	DeleteArticle(id int64) error
}

// Struct service (Sudah benar di kode Anda)
type service struct {
	repo      Repository
	validator *validator.Validate
}

// Fungsi NewService (Sudah benar di kode Anda)
func NewService(repo Repository) Service {
	return &service{
		repo:      repo,
		validator: validator.New(),
	}
}

// Fungsi CreateArticle (Sudah benar di kode Anda)
func (s *service) CreateArticle(input ArticleInput) (Article, error) {
	// 1. Validasi Input
	err := s.validator.Struct(input)
	if err != nil {
		// Bisa format error validasi agar lebih user-friendly
		return Article{}, fmt.Errorf("validation error: %w", err)
	}

	// 2. Panggil Repository
	newID, err := s.repo.CreateArticle(input)
	if err != nil {
		return Article{}, err // Error sudah di-wrap di repo
	}

	// 3. Ambil data yang baru dibuat untuk dikembalikan (best practice)
	createdArticle, err := s.repo.GetArticleByID(newID)
	if err != nil {
		// Mungkin log error, tapi kembalikan ID saja jika Get gagal setelah Create sukses
		log.Printf("Warning: article created (id: %d) but failed to retrieve details: %v", newID, err)
		// Kembalikan struct dasar jika gagal retrieve, agar tidak error total
		return Article{ID: newID, Title: input.Title, Content: input.Content, Category: input.Category, Status: input.Status}, nil
		// Atau: return Article{ID: newID}, fmt.Errorf(...) // Jika ingin tetap mengembalikan error
	}

	return createdArticle, nil
}

// --- PERBAIKAN IMPLEMENTASI GetArticles ---
// Menghapus parameter limit, offset dan logika defaultnya
func (s *service) GetArticles() ([]Article, error) {
	// Langsung panggil repository tanpa limit/offset
	articles, err := s.repo.GetArticles()
	if err != nil {
		// Mungkin wrap error di sini jika perlu konteks tambahan
		return nil, fmt.Errorf("service error fetching articles: %w", err)
	}
	// Kembalikan slice kosong jika repo mengembalikan nil (meski sebaiknya repo mengembalikan slice kosong)
	if articles == nil {
		return []Article{}, nil
	}
	return articles, nil
}

// ---------------------------------------

// Fungsi GetArticleByID (Sudah benar di kode Anda)
func (s *service) GetArticleByID(id int64) (Article, error) {
	if id <= 0 {
		return Article{}, fmt.Errorf("invalid article ID: %d", id)
	}
	// Langsung panggil repository
	article, err := s.repo.GetArticleByID(id)
	if err != nil {
		// Wrap error jika perlu
		return Article{}, fmt.Errorf("service error fetching article by ID %d: %w", id, err)
	}
	return article, nil
}

// Fungsi UpdateArticle (Sudah benar di kode Anda)
func (s *service) UpdateArticle(id int64, input ArticleInput) (Article, error) {
	if id <= 0 {
		return Article{}, fmt.Errorf("invalid article ID: %d", id)
	}
	// 1. Validasi Input
	err := s.validator.Struct(input)
	if err != nil {
		return Article{}, fmt.Errorf("validation error: %w", err)
	}

	// 2. Panggil Repository untuk update
	err = s.repo.UpdateArticle(id, input)
	if err != nil {
		return Article{}, fmt.Errorf("service error updating article %d: %w", id, err)
	}

	// 3. Ambil data yang sudah diupdate untuk dikembalikan
	updatedArticle, err := s.repo.GetArticleByID(id)
	if err != nil {
		log.Printf("Warning: article updated (id: %d) but failed to retrieve updated details: %v", id, err)
		// Kembalikan data input sebagai fallback jika retrieve gagal
		return Article{ID: id, Title: input.Title, Content: input.Content, Category: input.Category, Status: input.Status}, nil
		// Atau: return Article{}, fmt.Errorf(...)
	}
	return updatedArticle, nil
}

// Fungsi DeleteArticle (Sudah benar di kode Anda)
func (s *service) DeleteArticle(id int64) error {
	if id <= 0 {
		return fmt.Errorf("invalid article ID: %d", id)
	}
	err := s.repo.DeleteArticle(id)
	if err != nil {
		return fmt.Errorf("service error deleting article %d: %w", id, err)
	}
	return nil
}

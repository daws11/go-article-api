package article

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	CreateArticle(input ArticleInput) (Article, error) // Return Article agar bisa lihat hasil create
	GetArticles(limit, offset int) ([]Article, error)
	GetArticleByID(id int64) (Article, error)
	UpdateArticle(id int64, input ArticleInput) (Article, error) // Return Article hasil update
	DeleteArticle(id int64) error
}

type service struct {
	repo      Repository
	validator *validator.Validate
}

func NewService(repo Repository) Service {
	return &service{
		repo:      repo,
		validator: validator.New(), // Inisialisasi validator
	}
}

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
		return Article{ID: newID}, fmt.Errorf("article created (id: %d) but failed to retrieve details: %w", newID, err)
	}

	return createdArticle, nil
}

func (s *service) GetArticles(limit, offset int) ([]Article, error) {
	if limit <= 0 {
		limit = 10 // Default limit
	}
	if offset < 0 {
		offset = 0 // Default offset
	}
	return s.repo.GetArticles(limit, offset)
}

func (s *service) GetArticleByID(id int64) (Article, error) {
	if id <= 0 {
		return Article{}, fmt.Errorf("invalid article ID: %d", id)
	}
	return s.repo.GetArticleByID(id)
}

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
		return Article{}, err
	}

	// 3. Ambil data yang sudah diupdate untuk dikembalikan
	updatedArticle, err := s.repo.GetArticleByID(id)
	if err != nil {
		return Article{}, fmt.Errorf("article updated (id: %d) but failed to retrieve updated details: %w", id, err)
	}
	return updatedArticle, nil
}

func (s *service) DeleteArticle(id int64) error {
	if id <= 0 {
		return fmt.Errorf("invalid article ID: %d", id)
	}
	return s.repo.DeleteArticle(id)
}

package article

import "time"

type Article struct {
	ID          int64     `json:"id"` // Biasanya ID dikirim sebagai int/number di JSON
	Title       string    `json:"title" validate:"required"`
	Content     string    `json:"content" validate:"required"`
	Category    string    `json:"category" validate:"required,min=3"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
	Status      string    `json:"status" validate:"required,oneof=Publish Draft Thrash"` // oneof untuk enum
}

// Input struct terpisah untuk create/update, tanpa ID dan timestamp generated DB
type ArticleInput struct {
	Title    string `json:"title" validate:"required"`
	Content  string `json:"content" validate:"required"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status" validate:"required,oneof=Publish Draft Thrash"`
}

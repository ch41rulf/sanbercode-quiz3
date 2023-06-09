package structs

import "time"

type Book struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	ReleaseYear int64     `json:"release_year"`
	Price       string    `json:"price"`
	TotalPage   int64     `json:"total_page"`
	Thickness   string    `json:"thickness"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CategoryID  int64     `json:"category_id"`
}

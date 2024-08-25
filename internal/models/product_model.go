package models

import "time"

type ProductRequest struct {
	Name          string `json:"name" validate:"required"`
	Description   string `json:"description"`
	CurrentPrice  int    `json:"current_price"`
	OriginalPrice int    `json:"original_price"`
	Currency      string `json:"currency"`
}

type ProductDetailResponse struct {
	ID             uint    `json:"id"`
	Name           string  `json:"name"`
	Rating         float64 `json:"rating"`
	ReviewsCount   int     `json:"reviews_count"`
	ReviewersCount int     `json:"reviewers_count"`
	Description    string  `json:"description"`
	CurrentPrice   int     `json:"current_price"`
	OriginalPrice  int     `json:"original_price"`
	Currency       string  `json:"currency"`
	Colors         []Color
	Sizes          []Size
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Image 		   string
}

type ProductCreateResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Color struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Hex  string `json:"hex"`
}

type Size struct {
	ID    uint `json:"id"`
	Value int  `json:"value"`
}

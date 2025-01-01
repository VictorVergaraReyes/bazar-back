package entities

import "time"

type Product struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Dimensions  Dimensions `json:"dimensions"`
	Price       float64    `json:"price"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Image       string     `json:"image"`
	Stock       int        `json:"stock"`
}

type Dimensions struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
	Depth  float64 `json:"depth"`
}

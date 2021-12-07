package models

type Category struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Good struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Manufacturer string `json:"manufacturer"`
	Brand string `json:"brand"`
	Material string `json:"material"`
	Weight int `json:"weight"`
	SIZE    string `json:"size"`
	Color  string `json:"color"`
	Price float64 `json:"price"`
	CategoryID string `json:"category_id" db:"category_id"`
}

type Client struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Surname string `json:"surname"`

	Email string `json:"email"`
	Password string `json:"password"`
}


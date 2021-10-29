package models

type Bag struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	Brand string `json:"brand"`
	Material string `json:"material"`
	Weight int `json:"weight"`
	SIZE    string `json:"size"`
	Color  string `json:"color"`
	Price float64 `json:"price"`
}

type Bracelet struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Brand string `json:"brand"`
	Material string `json:"material"`
	SIZE    int `json:"size"`
	Color  string `json:"color"`
	Price float64 `json:"price"`
}

type Necklace struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Brand string `json:"brand"`
	Material string `json:"material"`
	SIZE    int `json:"size"`
	Color  string `json:"color"`
	Price float64 `json:"price"`
}

type Client struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email string `json:"email"`
}
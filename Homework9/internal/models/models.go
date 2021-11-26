package models


type Accesory struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Manufacturer string `json:"manufacturer" db:"manufacturer"`
	Brand string `json:"brand" db:"brand"`
	Material string `json:"material" db:"material"`
	Weight int `json:"weight" db:"weight"`
	SIZE    string `json:"size" db:"size"`
	Color  string `json:"color" db:"color"`
	Price float64 `json:"price" db:"price"`
	CategoryID string `json:"category_id" db:"category_id"`
}


type AccesoryFilter struct {
	Query *string `json:"query"`
}


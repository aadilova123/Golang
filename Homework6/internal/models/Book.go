package models

type Book struct {
	ID      int `json:"id"`
	Title   string `json:"title"`
	Author *Author `json:"author"`
	Year 	int  `json:"year"`
	isAvailable  bool    `json:"isAvailable"`
	Description  string  `json:"description"`
	Price 		 float64 `json:"price"`
}
  
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

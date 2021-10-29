package models

type Romance struct {
	ID      int `json:"id"`
	Title   string `json:"title"`
	Author *Author `json:"author"`
	Year 	int  `json:"year"`
	Description  string  `json:"description"`
	Category    string  `json:"category"`
	ImageUrl    string  `json:"image"`
}
  
type Author struct {
	ID      int `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

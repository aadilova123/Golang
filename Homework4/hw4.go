package main

type Students struct {
    Students []Student `json:"students"`
}


type Student struct {
	Id string `json:"id"`
    FullName   string `json:"full_name"`
    Faculty   string `json:"faculty"`
	Gender   string `json:"gender"`
	Gpa   string `json:"gpa"`
    Age    int    `json:"Age"`
    Social Social `json:"social"`
	YearofStudy int `json:"year"`
}

type Social struct {
    Facebook string `json:"facebook"`
    Twitter  string `json:"twitter"`
}


func addYearofStudy(a int ) int {
	return a + 1
}




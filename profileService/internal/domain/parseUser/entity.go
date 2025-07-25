package parseuser

import "time"

type Person struct {
	Name              string    `json:"Name"`
	RussianName       string    `json:"RusName"`
	Position          string    `json:"Position"`
	StudentType       string    `json:"Student_Type"`
	Rates             []float64 `json:"Rates"`
	NumberOfLanguages int       `json:"Langs_Number"`
	Mode              string    `json:"Mode"`
	Maxload           float64   `json:"Max_Load"`
	Institute         string    `json:"Institute"`
	Degree            bool      `json:"Degree"`
	Email             string    `json:"Email"`
	Alias             string    `json:"Alias"`
	EmploymentType    string    `json:"Employment_Type"`
	StartTime         time.Time `json:"Start_Time"`
	EndTime           time.Time `json:"End_Time"`
}

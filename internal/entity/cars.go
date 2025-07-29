package entity

type Cars struct {
	CarsName  string  `json:"cars_name"`
	DayRate   float64 `json:"day_rate"`
	MonthRate float64 `json:"month_rate"`
	Image     string  `json:"image"`
}

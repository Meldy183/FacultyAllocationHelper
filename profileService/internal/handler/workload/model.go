package workload

type Stats struct {
	T1    Classes `json:"t1"`
	T2    Classes `json:"t2"`
	T3    Classes `json:"t3"`
	Total Classes `json:"total"`
}

type Classes struct {
	Lec  int64   `json:"lec_hours"`
	Tut  int64   `json:"tut_hours"`
	Lab  int64   `json:"lab_hours"`
	Elec int64   `json:"elective_hours"`
	Rate float64 `json:"rate"`
}

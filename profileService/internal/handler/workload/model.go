package workload

type WorkloadStats struct {
	T1    Classes `json:"t1"`
	T2    Classes `json:"t2"`
	T3    Classes `json:"t3"`
	Total Classes `json:"total"`
}

type Classes struct {
	Lec  int     `json:"lec_hours"`
	Tut  int     `json:"tut_hours"`
	Lab  int     `json:"lab_hours"`
	Elec int     `json:"elective_hours"`
	Rate float64 `json:"rate"`
}
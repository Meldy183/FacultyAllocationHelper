package entities

const (
	Onsite = iota
	Mixed
	Remote
)
const (
	ALL = iota
	DS
	GAMEDEV
	HUM
	RO
	SDE
	SNE
)
const (
	Core = iota
	Tech
	Hum
)
const (
	FirstBlock = iota
	SecondBlock
	Full
)

type Course struct {
	Name          string   `json:"name"`
	OfficialName  string   `json:"OfName"`
	Year          int      `json:"Year"`
	Semester      string   `json:"Semester"`
	Type          string   `json:"Type"`
	AcademicYear  string   `json:"AcademicYear"`
	Programms     []string `json:"Programms"`
	Tracks        []string `json:"Tracks"`
	Form          string   `json:"Form"`
	LectureFormat string   `json:"LecFormat"`
	LabFormat     string   `json:"LabFormat"`
	PI            string   `json:"PI"`
	TI            string   `json:"TI"`
	TA            []string `json:"TA"`
	LecHours      int      `json:"Lec_Hours"`
	LabHours      int      `json:"Lab_Hours"`
	InstituteId   string   `json:"Institute_Id"`
}

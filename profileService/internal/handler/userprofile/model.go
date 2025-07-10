package userprofile

type AddProfileRequest struct {
	NameEnglish      string `json:"name_english"`
	Email            string `json:"email"`
	Alias            string `json:"alias"`
	InstituteID      int    `json:"institute_id"`
	Position         string `json:"position"`
	IsRepresentative bool   `json:"isRepr"`
}

type GetProfileResponse struct {
	ProfileID      int64         `json:"profileID"`
	NameEnglish    string        `json:"nameEng"`
	NameRussian    string        `json:"nameRu"`
	Alias          string        `json:"alias"`
	Email          string        `json:"email"`
	Position       string        `json:"position"`
	Institute      string        `json:"institute"`
	InstituteID    int           `json:"instituteID"`
	Workload       float64       `json:"workload"`
	StudentType    string        `json:"studentType"`
	Degree         bool          `json:"degree"`
	Fsro           string        `json:"FSRO,omitempty"`
	Languages      []Lang        `json:"languages"`
	Courses        []Course      `json:"courses"`
	EmploymentType string        `json:"employmentType"`
	HiringStatus   string        `json:"hiringStatus,omitempty"`
	Mode           string        `json:"mode"`
	MaxLoad        int           `json:"maxLoad"`
	FrontalHours   int           `json:"frontalHours,omitempty"`
	ExtraActivity  float64       `json:"extraActivity,omitempty"`
	WorkloadStats  WorkloadStats `json:"workloadStats,omitempty"`
}

type Lang struct {
	Language string `json:"language"`
}

type Course struct {
	CourseInstanceID int64 `json:"id"`
}

type WorkloadStats struct {
	UnitStat []UnitStat `json:"unitStat"`
	Total    Total      `json:"total"`
}

type UnitStat struct {
	Trimester string  `json:"trimester"`
	Classes   Classes `json:"classes"`
}
type Total struct {
	Lectures   int `json:"totalLec"`
	TotalTut   int `json:"totalTut"`
	TotalLab   int `json:"totalLab"`
	TotalElect int `json:"totalElect"`
	TotalRate  int `json:"totalRate"`
}

type Classes struct {
	Lec  int `json:"lec"`
	Tut  int `json:"tut"`
	Lab  int `json:"lab"`
	Elec int `json:"elec"`
	Rate int `json:"rate"`
}

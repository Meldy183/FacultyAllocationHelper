package userprofile

type AddProfileRequest struct {
	NameEnglish      string `json:"name_eng"`
	Email            string `json:"email"`
	Alias            string `json:"alias"`
	InstituteID      int    `json:"institute_id"`
	Position         string `json:"position"`
	IsRepresentative bool   `json:"is_repr"`
}

type GetProfileResponse struct {
	ProfileID      int64          `json:"profile_id"`
	NameEnglish    string         `json:"name_eng"`
	NameRussian    *string        `json:"name_ru"`
	Alias          string         `json:"alias"`
	Email          string         `json:"email"`
	Position       string         `json:"position"`
	Institute      string         `json:"institute"`
	InstituteID    int            `json:"institute_id"`
	Workload       *float64       `json:"workload"`
	StudentType    *string        `json:"student_type"`
	Degree         *bool          `json:"degree"`
	Fsro           *string        `json:"fsro,omitempty"`
	Languages      *[]Lang        `json:"languages"`
	Courses        *[]Course      `json:"courses"`
	EmploymentType *string        `json:"employment_type"`
	HiringStatus   *string        `json:"hiring_status,omitempty"`
	Mode           *string        `json:"mode"`
	MaxLoad        *int           `json:"max_load"`
	FrontalHours   *int           `json:"frontal_hours,omitempty"`
	ExtraActivity  *float64       `json:"extra_activity,omitempty"`
	WorkloadStats  *WorkloadStats `json:"workload_stats,omitempty"`
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

type GetAllFacultiesResponse struct {
}

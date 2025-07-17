package facultyProfile

type AddProfileRequest struct {
	NameEnglish      string `json:"name_eng"`
	Email            string `json:"email"`
	Alias            string `json:"alias"`
	InstituteID      []int  `json:"institute_ids"`
	PositionID       int    `json:"position_id"`
	IsRepresentative bool   `json:"is_repr"`
	Year             int    `json:"year"`
}

type GetProfileResponse struct {
	ProfileID      int64          `json:"profile_id"`
	NameEnglish    string         `json:"name_eng"`
	NameRussian    *string        `json:"name_ru"`
	Alias          string         `json:"alias"`
	Email          string         `json:"email"`
	PositionName   string         `json:"position_name"`
	InstituteNames []string       `json:"institute_names"`
	Workload       *float64       `json:"workload"`
	StudentType    *string        `json:"student_type"`
	Degree         *bool          `json:"degree"`
	Fsro           *string        `json:"fsro"`
	LanguageCodes  *[]Lang        `json:"languages"`
	Courses        *[]Course      `json:"courses"`
	EmploymentType *string        `json:"employment_type"`
	HiringStatus   *string        `json:"hiring_status"`
	Mode           *string        `json:"mode"`
	MaxLoad        *int           `json:"max_load"`
	FrontalHours   *int           `json:"frontal_hours"`
	ExtraActivity  *float64       `json:"extra_activity"`
	WorkloadStats  *WorkloadStats `json:"workload_stats"`
}

type Lang struct {
	Language string `json:"language_code"`
}

type Course struct {
	CourseID             *int64     `json:"course_id"`
	BriefName            *string    `json:"brief_name"`
	OfficialName         *string    `json:"official_name"`
	AcademicYearName     *string    `json:"academic_year_name"`
	SemesterName         *string    `json:"semester_name"`
	StudyPrograms        *[]string  `json:"study_program_names"`
	InstituteName        *string    `json:"responsible_institute_name"`
	Tracks               *[]string  `json:"track_names"`
	IsAllocationFinished *bool      `json:"allocation_finished"`
	Mode                 *string    `json:"mode"`
	StudyYear            *int       `json:"study_year"`
	Form                 *string    `json:"form"`
	LectureHours         *int       `json:"lecture_hours"`
	LabHours             *int       `json:"lab_hours"`
	GroupsNeeded         *int       `json:"groups_needed"`
	GroupsTaken          *int       `json:"groups_taken"`
	PI                   PI         `json:"pi"`
	TI                   PI         `json:"ti"`
	TAs                  *[]Faculty `json:"tas"`
}

type PI struct {
	AllocationStatus *string  `json:"allocation_status"`
	ProfileData      *Faculty `json:"profile_data"`
}
type Faculty struct {
	ProfileID      *int64   `json:"profile_id"`
	NameEng        *string  `json:"name_eng"`
	Alias          *string  `json:"alias"`
	Email          *string  `json:"email"`
	PositionName   *string  `json:"position_name"`
	InstituteNames []string `json:"institute_names"`
	Workload       *float64 `json:"workload"`
	Classes        []string `json:"classes"`
}
type WorkloadStats struct {
	T1    []Classes `json:"t1"`
	T2    []Classes `json:"t2"`
	T3    []Classes `json:"t3"`
	Total []Classes `json:"total"`
}

type Classes struct {
	Lec  int `json:"lec_hours"`
	Tut  int `json:"tut_hours"`
	Lab  int `json:"lab_hours"`
	Elec int `json:"elective_hours"`
	Rate int `json:"rate"`
}

type GetAllFacultiesResponse struct {
	Profiles []ShortProfile `json:"profiles"`
}

type ShortProfile struct {
	ProfileID   int64  `json:"profile_id"`
	NameEnglish string `json:"name_eng"`
	Alias       string `json:"alias"`
	Email       string `json:"email"`
	Position    string `json:"position"`
	Institute   string `json:"institute"`
}
type InstituteObj struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PositionObj struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetFacultyFiltersResponse struct {
	InstituteFilters []InstituteObj `json:"institute"`
	PositionFilters  []PositionObj  `json:"position"`
}

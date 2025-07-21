package facultyProfile

import (
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/workload"
)

type AddProfileRequest struct {
	NameEnglish  string `json:"name_eng"`
	Email        string `json:"email"`
	Alias        string `json:"alias"`
	InstituteIDs []int  `json:"institute_id"`
	PositionID   int    `json:"position_id"`
	Year         int    `json:"year"`
}

type AddProfileResponse struct {
	ProfileVersionID int64    `json:"profile_id"`
	NameEnglish      string   `json:"name_eng"`
	Email            string   `json:"email"`
	Alias            string   `json:"alias"`
	InstituteNames   []string `json:"institute_names"`
	PositionName     string   `json:"position_name"`
	Year             int      `json:"year"`
}

type GetProfileResponse struct {
	ProfileVersionID int64           `json:"profile_id"`
	Year             int             `json:"year"`
	NameEnglish      string          `json:"name_eng"`
	NameRussian      *string         `json:"name_ru"`
	Alias            string          `json:"alias"`
	Email            string          `json:"email"`
	PositionName     string          `json:"position_name"`
	InstituteNames   []string        `json:"institute_names"`
	StudentType      *string         `json:"student_type"`
	Degree           *bool           `json:"degree"`
	Fsro             *string         `json:"fsro"`
	LanguageCodes    *[]Lang         `json:"languages"`
	EmploymentType   *string         `json:"employment_type"`
	HiringStatus     *string         `json:"hiring_status"`
	Mode             *string         `json:"mode"`
	MaxLoad          *int            `json:"max_load"`
	FrontalHours     *int            `json:"frontal_hours"`
	ExtraActivity    *float64        `json:"extra_activities"`
	WorkloadStats    *workload.Stats `json:"workload_stats"`
}

type Lang struct {
	Language string `json:"language_code"`
}
type GetAllFacultiesResponse struct {
	Profiles []ShortProfile `json:"profiles"`
}

type ShortProfile struct {
	ProfileVersionID int64    `json:"profile_id"`
	NameEnglish      string   `json:"name_eng"`
	Alias            string   `json:"alias"`
	Email            string   `json:"email"`
	Position         string   `json:"position_name"`
	Institutes       []string `json:"institute_names"`
}

type EditProfileRequest struct {
	ProfileVersionID int64   `json:"profile_id"`
	Year             int     `json:"year"`
	NameEng          string  `json:"name_eng"`
	NameRu           string  `json:"name_ru"`
	Alias            string  `json:"alias"`
	Email            string  `json:"email"`
	PositionID       int     `json:"position_id"`
	InstituteIDs     *[]int  `json:"institute_id"`
	StudentType      *string `json:"student_type"`
	Degree           *bool   `json:"degree"`
	Languages        *[]Lang `json:"languages"`
	EmploymentType   *string `json:"employment_type"`
	HiringStatus     *string `json:"hiring_status"`
	FSRO             *string `json:"fsro"`
	Mode             *string `json:"mode"`
}

type EditProfileResponse struct {
	ProfileVersionID int64     `json:"profile_id"`
	Year             int       `json:"year"`
	NameEng          string    `json:"name_eng"`
	NameRu           string    `json:"name_ru"`
	Alias            string    `json:"alias"`
	Email            string    `json:"email"`
	PositionName     *string   `json:"position_name"`
	InstituteNames   *[]string `json:"institute_names"`
	StudentType      *string   `json:"student_type"`
	Degree           *bool     `json:"degree"`
	Languages        *[]Lang   `json:"languages"`
	EmploymentType   *string   `json:"employment_type"`
	HiringStatus     *string   `json:"hiring_status"`
	FSRO             *string   `json:"fsro"`
	Mode             *string   `json:"mode"`
}

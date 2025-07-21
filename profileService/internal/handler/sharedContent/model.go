package sharedContent

type Course struct {
	InstanceID           *int64    `json:"course_id"`
	BriefName            *string   `json:"brief_name"`
	OfficialName         *string   `json:"official_name"`
	AcademicYearName     *string   `json:"academic_year_name"`
	SemesterName         *string   `json:"semester_name"`
	StudyPrograms        []*string `json:"study_program_names"`
	InstituteName        *string   `json:"responsible_institute_name"`
	Tracks               []*string `json:"track_names"`
	IsAllocationFinished *bool     `json:"allocation_finished"`
	Mode                 *string   `json:"mode"`
	Year                 *int      `json:"year"`
	Form                 *string   `json:"form"`
	LectureHours         *int      `json:"lecture_hours"`
	LabHours             *int      `json:"lab_hours"`
	GroupsNeeded         *int      `json:"groups_needed"`
	GroupsTaken          *int      `json:"groups_taken"`
	PI                   PI        `json:"pi"`
	TI                   PI        `json:"ti"`
	TAs                  []Faculty `json:"tas"`
}

type PI struct {
	AllocationStatus *string  `json:"allocation_status"`
	ProfileData      *Faculty `json:"profile_data"`
}
type Faculty struct {
	ProfileVersionID int64    `json:"profile_id"`
	NameEng          *string  `json:"name_eng"`
	Alias            *string  `json:"alias"`
	Email            *string  `json:"email"`
	PositionName     *string  `json:"position_name"`
	InstituteNames   []string `json:"institute_names"`
	Classes          []string `json:"classes"`
	IsConfirmed      bool     `json:"is_confirmed"`
}

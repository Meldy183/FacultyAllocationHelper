package filters

type FilterObj struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetFacultyFiltersResponse struct {
	InstituteFilters []FilterObj `json:"institute"`
	PositionFilters  []FilterObj `json:"position"`
}

type GetCourseFiltersResponse struct {
	AllocationStatus bool        `json:"allocaion_not_finished"`
	YearOfStudy      []FilterObj `json:"academic_year"`
	Semester         []FilterObj `json:"semester"`
	StudyProgram     []FilterObj `json:"position"`
	InstituteFilters []FilterObj `json:"institute"`
}

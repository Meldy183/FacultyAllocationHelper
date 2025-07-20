package courses

import (
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/sharedContent"
)

type GetCourseListResponse struct {
	Courses []sharedContent.Course `json:"courses"`
}

type AddNewCourseRequest struct {
	BriefName              string `json:"name"`
	AcademicYearID         int    `json:"academic_year_id"`
	SemesterID             int    `json:"semester_id"`
	Year                   int    `json:"year"`
	ProgramIDs             []int  `json:"program_ids"`
	TrackIDs               []int  `json:"track_ids"`
	ResponsibleInstituteID int64  `json:"responsible_institute_id"`
	GroupsNeeded           int    `json:"groups_needed"`
	IsElective             bool   `json:"is_elective"`
}

type AddNewCourseResponse struct {
	CourseInstanceID         int64                   `json:"course_id"`
	BriefName                string                  `json:"name"`
	AcademicYearName         string                  `json:"academic_year_name"`
	SemesterName             string                  `json:"semester_name"`
	ProgramNames             []string                `json:"program_names"`
	TrackNames               []string                `json:"track_names"`
	ResponsibleInstituteName string                  `json:"responsible_institute_name"`
	GroupsNeeded             int                     `json:"groups_needed"`
	Pi                       sharedContent.PI        `json:"pi"`  // empty
	Ti                       sharedContent.PI        `json:"ti"`  // empty
	TAs                      []sharedContent.Faculty `json:"tas"` // empty
}

type GetCourseResponse struct {
	Course sharedContent.Course
}

type EditCourseRequest struct {
	CourseInstanceID       int     `json:"course_id"`
	BriefName              string  `json:"name"`
	OfficialName           *string `json:"official_name"`
	ResponsibleInstituteID int64   `json:"responsible_institute_id"`
	StudyProgramIDS        []int   `json:"study_program_ids"`
	TrackIDs               []int   `json:"track_ids"`
	Mode                   *string `json:"mode"`
	AcademicYearID         int     `json:"academic_year_id"`
	SemesterID             int     `json:"semester_id"`
	Form                   *string `json:"form"`
	LectureHours           *int    `json:"lecture_hours"`
	LabHours               *int    `json:"lab_hours"`
}

type EditCourseResponse struct {
	CourseInstanceID         int      `json:"course_id"`
	BriefName                string   `json:"name"`
	OfficialName             *string  `json:"official_name"`
	ResponsibleInstituteName string   `json:"responsible_institute_name"`
	StudyProgramNames        []string `json:"study_program_names"`
	TrackNames               []string `json:"track_names"`
	Mode                     *string  `json:"mode"`
	AcademicYearName         string   `json:"academic_year_name"`
	SemesterName             string   `json:"semester_name"`
	Form                     *string  `json:"form"`
	LectureHours             *int     `json:"lecture_hours"`
	LabHours                 *int     `json:"lab_hours"`
}

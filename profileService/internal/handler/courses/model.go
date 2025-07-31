package courses

import (
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/sharedContent"
)

type GetCourseListResponse struct {
	Courses []sharedContent.Course `json:"courses"`
}

type AddNewCourseRequest struct {
	BriefName              string  `json:"brief_name"`
	OfficialName           string  `json:"official_name"`
	AcademicYearID         int64   `json:"academic_year_id"`
	SemesterID             int64   `json:"semester_id"`
	Year                   int64   `json:"year"`
	ProgramIDs             []int64 `json:"program_ids"`
	TrackIDs               []int64 `json:"track_ids"`
	ResponsibleInstituteID int64   `json:"responsible_institute_id"`
	GroupsNeeded           int64   `json:"groups_needed"`
	IsElective             bool    `json:"is_elective"`
}

type AddNewCourseResponse struct {
	CourseInstanceID         int64                   `json:"course_id"`
	BriefName                string                  `json:"brief_name"`
	OfficialName             string                  `json:"official_name"`
	AcademicYearName         string                  `json:"academic_year_name"`
	SemesterName             string                  `json:"semester_name"`
	ProgramNames             []string                `json:"program_names"`
	TrackNames               []string                `json:"track_names"`
	ResponsibleInstituteName string                  `json:"responsible_institute_name"`
	GroupsNeeded             int64                   `json:"groups_needed"`
	Pi                       sharedContent.PI        `json:"pi"`
	Ti                       sharedContent.PI        `json:"ti"`
	TAs                      []sharedContent.Faculty `json:"tas"`
}

type GetCourseResponse struct {
	Course sharedContent.Course
}

type EditCourseRequest struct {
	CourseInstanceID       int64   `json:"course_id"`
	BriefName              string  `json:"brief_name"`
	OfficialName           *string `json:"official_name"`
	ResponsibleInstituteID int64   `json:"responsible_institute_id"`
	StudyProgramIDS        []int   `json:"study_program_ids"`
	TrackIDs               []int   `json:"track_ids"`
	Mode                   *string `json:"mode"`
	AcademicYearID         int64   `json:"academic_year_id"`
	SemesterID             int64   `json:"semester_id"`
	Form                   *string `json:"form"`
	LectureHours           *int64  `json:"lecture_hours"`
	LabHours               *int64  `json:"lab_hours"`
}

type EditCourseResponse struct {
	CourseInstanceID         int64    `json:"course_id"`
	BriefName                string   `json:"brief_name"`
	OfficialName             *string  `json:"official_name"`
	ResponsibleInstituteName string   `json:"responsible_institute_name"`
	StudyProgramNames        []string `json:"study_program_names"`
	TrackNames               []string `json:"track_names"`
	Mode                     *string  `json:"mode"`
	AcademicYearName         string   `json:"academic_year_name"`
	SemesterName             string   `json:"semester_name"`
	Form                     *string  `json:"form"`
	LectureHours             *int64   `json:"lecture_hours"`
	LabHours                 *int64   `json:"lab_hours"`
}

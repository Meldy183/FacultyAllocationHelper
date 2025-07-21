package semester

import "context"

type Service interface {
	GetSemesterNameByID(ctx context.Context, semesterID int64) (*string, error)
	GetAllSemesters(ctx context.Context) ([]Semester, error)
}

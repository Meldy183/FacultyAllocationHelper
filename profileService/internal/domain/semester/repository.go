package semester

import "context"

type Repository interface {
	GetSemesterNameByID(ctx context.Context, semesterID int64) (*string, error)
}

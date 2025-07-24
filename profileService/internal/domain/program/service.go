package program

import "context"

type Service interface {
	GetProgramNamesByInstanceID(ctx context.Context, courseID int64) ([]*string, error)
	GetProgramNameByID(ctx context.Context, id int64) (*string, error)
	GetAllPrograms(ctx context.Context) ([]*Program, error)
}

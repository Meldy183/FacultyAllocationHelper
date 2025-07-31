package program

import "context"

type Service interface {
	GetProgramNamesByInstanceID(ctx context.Context, courseID int64) ([]*string, error)
	GetProgramNameByID(ctx context.Context, id int64) (*string, error)
	GetProgramIDByName(ctx context.Context, name string) (*int64, error)
	GetAllPrograms(ctx context.Context) ([]*Program, error)
}

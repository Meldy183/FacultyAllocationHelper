package program

import "context"

type Repository interface {
	GetAllPrograms(ctx context.Context) ([]*Program, error)
	GetProgramNameByID(ctx context.Context, id int64) (*string, error)
	GetProgramIDByName(ctx context.Context, name string) (*int64, error)
}

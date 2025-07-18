package program

import "context"

type Repository interface {
	GetAllPrograms(ctx context.Context) ([]*Program, error)
	GetProgramsByID(ctx context.Context, code string) (*Program, error)
}

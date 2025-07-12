package lab

import "context"

type Repository interface {
	GetAllLabs(ctx context.Context) ([]*Lab, error)
	GetLabsByInstituteID(ctx context.Context, instituteID int64) ([]*Lab, error)
}

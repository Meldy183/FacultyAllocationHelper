package lab

import "context"

type Repository interface {
	GetAllLabs(ctx context.Context) ([]int64, error)
	GetLabsByInstituteID(ctx context.Context, instituteID int64) ([]int64, error)
	GetLabByID(ctx context.Context, labID int64) (*Lab, error)
}

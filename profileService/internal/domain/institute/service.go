package institute

import "context"

type Service interface {
	GetInstituteByID(ctx context.Context, instituteID int64) (*Institute, error)
	GetAllInstitutes(ctx context.Context) ([]*Institute, error)
	GetInstituteIDByName(ctx context.Context, instituteName string) (*int64, error)
}

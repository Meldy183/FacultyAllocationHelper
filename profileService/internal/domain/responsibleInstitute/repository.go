package responsibleInstitute

import "context"

type Repository interface {
	GetResponsibleInstituteNameByID(ctx context.Context, responsibleInstituteID int64) (*string, error)
	GetAllInstitutes(ctx context.Context) ([]ResponsibleInstitute, error)
	GetResponsibleInstituteIDByName(ctx context.Context, responsubleInstituteName string) (*int64, error)
}

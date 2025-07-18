package responsibleInstitute

import "context"

type Repository interface {
	GetResponsibleInstituteNameByID(ctx context.Context, responsibleInstituteID int64) (*string, error)
}

package responsibleInstitute

import "context"

type Service interface {
	GetResponsibleInstituteNameByID(ctx context.Context, instituteID int64) (*string, error)
}

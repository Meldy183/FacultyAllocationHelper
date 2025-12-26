package responsibleInstitute

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	GetResponsibleInstituteNameByID(ctx context.Context, responsibleInstituteID uuid.UUID) (*string, error)
	GetAllInstitutes(ctx context.Context) ([]ResponsibleInstitute, error)
	GetResponsibleInstituteIDByName(ctx context.Context, responsubleInstituteName string) (*uuid.UUID, error)
}

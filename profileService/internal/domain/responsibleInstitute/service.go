package responsibleInstitute

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	GetResponsibleInstituteNameByID(ctx context.Context, instituteID uuid.UUID) (*string, error)
	GetResponsibleInstituteIDByName(ctx context.Context, responsubleInstituteName string) (*uuid.UUID, error)
}

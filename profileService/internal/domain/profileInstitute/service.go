package profileInstitute

import (
	"context"

	"github.com/google/uuid"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/institute"
)

type Service interface {
	GetUserInstitutesByProfileID(ctx context.Context, profileID uuid.UUID) ([]*institute.Institute, error)
	AddUserInstitute(ctx context.Context, userInstitute *UserInstitute) error
}

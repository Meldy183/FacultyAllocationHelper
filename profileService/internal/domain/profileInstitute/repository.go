package profileInstitute

import (
	"context"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/institute"
)

type Repository interface {
	GetUserInstitutesByProfileID(ctx context.Context, profileID int64) ([]*institute.Institute, error)
	AddUserInstitute(ctx context.Context, userInstitute *UserInstitute) error
}

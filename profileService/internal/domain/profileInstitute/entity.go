package profileInstitute

import "github.com/google/uuid"

type UserInstitute struct {
	UserInstituteID uuid.UUID
	InstituteID     uuid.UUID
	ProfileID       uuid.UUID
}

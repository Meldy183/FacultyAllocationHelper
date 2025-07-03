package userprofile

type Repository interface {
	Create(profile *UserProfile) error
	GetByProfileID(profileID int64) (*UserProfile, error)
	GetByUserID(userID string) (*UserProfile, error)
	Update(profile *UserProfile) error
}

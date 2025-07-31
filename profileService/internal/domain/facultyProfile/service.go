package facultyProfile

import "context"

type Service interface {
	AddProfile(ctx context.Context, profile *UserProfile) error
	GetProfileByID(ctx context.Context, profileID int64) (*UserProfile, error)
	UpdateProfileByID(ctx context.Context, profile *UserProfile) error
	GetProfilesByFilters(ctx context.Context, institutes []int64, positions []int64) ([]int64, error)
}

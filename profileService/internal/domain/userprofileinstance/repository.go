package userprofileversion

import "context"

type Repository interface {
	AddProfile(ctx context.Context, profile *UserProfileVersion) error
	GetProfileByID(ctx context.Context, profileID int64) (*UserProfileVersion, error)
	UpdateProfileByID(ctx context.Context, profile *UserProfileVersion) error
	GetProfilesByFilter(ctx context.Context, institutes []int, positions []int) ([]int64, error)
}

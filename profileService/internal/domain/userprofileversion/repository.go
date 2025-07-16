package userprofileversion

import "context"

type Repository interface {
	AddProfileVersion(ctx context.Context, profile *UserProfileVersion) error
	GetProfileVersionByID(ctx context.Context, profileID int64) (*UserProfileVersion, error)
	UpdateProfileVersionByID(ctx context.Context, profile *UserProfileVersion) error
	GetProfileVersionsByFilter(ctx context.Context, institutes []int, positions []int) ([]int64, error)
}

package userprofile

import "context"

type Repository interface {
	AddProfile(ctx context.Context, profile *UserProfile) error
	GetProfileByID(ctx context.Context, profileID int64) (*UserProfile, error)
	UpdateProfileByID(ctx context.Context, profile *UserProfile) error
}

package userprofile

import "context"

type Repository interface {
	Create(ctx context.Context, profile *UserProfile) error
	GetByProfileID(ctx context.Context, profileID int64) (*UserProfile, error)
	Update(ctx context.Context, profile *UserProfile) error
	GetProfilesByFilter(ctx context.Context, institutes []int, positions []int) ([]int64, error)
}

package profileVersion

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	AddProfileVersion(ctx context.Context, tx *pgx.Tx, profile *ProfileVersion) error
	GetVersionByProfileID(ctx context.Context, profileID int64, year int64) (*ProfileVersion, error)
	GetVersionByVersionID(ctx context.Context, profileID int64) (*ProfileVersion, error)
}

//TODO: UpdateProfileVersionByID(ctx context.Context, profile *ProfileVersion) error

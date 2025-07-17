package profileVersion

import "context"

type Service interface {
	GetVersionByProfileID(ctx context.Context, profileID int64) (*ProfileVersion, error)
}

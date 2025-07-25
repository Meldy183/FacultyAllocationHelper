package parseuser

import "context"

type Repository interface {
	ParseUsers(ctx context.Context, users [][]string) (*[]Person, error)
}

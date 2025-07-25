package parseuser

import (
	"context"
)

type Service interface {
	ProcessUsers(ctx context.Context, users [][]string, studyyear int) error
}

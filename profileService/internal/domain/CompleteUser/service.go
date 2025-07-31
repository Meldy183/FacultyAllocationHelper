package CompleteUser

import "context"

type Service interface {
	AddFullUser(ctx context.Context, fulluser *FullUser) error
}

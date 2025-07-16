package logpage

import "context"

type Repository interface {
	GetLogpages(ctx context.Context, last_id string, limit string) ([]*LogPage, error)
	AddLogpage(ctx context.Context, logPage *LogPage) error
}

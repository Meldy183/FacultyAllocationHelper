package parsing

import (
	"context"
	"mime/multipart"
)

type Service interface {
	Parse(ctx context.Context, file *multipart.File) error
}

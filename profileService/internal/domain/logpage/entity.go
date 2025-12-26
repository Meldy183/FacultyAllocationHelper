package logpage

import (
	"time"

	"github.com/google/uuid"
)

type LogPage struct {
	LogPageID uuid.UUID
	UserID    string
	Action    string
	SubjectID uuid.UUID
	Timestamp time.Time
}

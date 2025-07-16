package logpage

import "time"

type LogPage struct {
	LogID     int64
	UserID    string
	Action    string
	SubjectID string
	Timestamp time.Time
}

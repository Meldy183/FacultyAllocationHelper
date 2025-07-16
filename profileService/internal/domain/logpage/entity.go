package logpage

import "time"

type LogPage struct {
	LogID     int64
	UserID    string
	Action    string
	SubjectID int64
	Timestamp time.Time
}

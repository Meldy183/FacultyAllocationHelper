package logpage

import "time"

type LogPage struct {
	LogPageID int
	UserID    string
	Action    string
	SubjectID int64
	Timestamp time.Time
}

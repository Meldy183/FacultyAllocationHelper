package track

import "github.com/google/uuid"

type Track struct {
	TrackID   uuid.UUID
	Name      string
	ProgramID uuid.UUID
}

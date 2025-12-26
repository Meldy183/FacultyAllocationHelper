package trackcourseinstance

import "github.com/google/uuid"

type TrackToCourseInstance struct {
	TrackCourseInstanceID uuid.UUID
	TrackID               uuid.UUID
	CourseInstanceID      uuid.UUID
}

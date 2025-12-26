package profileCourseInstance

import "github.com/google/uuid"

type ProfileVersionCourseInstance struct {
	ProfileCourseID  uuid.UUID
	ProfileVersionID uuid.UUID
	CourseInstanceID uuid.UUID
}

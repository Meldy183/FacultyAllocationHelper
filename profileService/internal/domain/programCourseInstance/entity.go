package programcourseinstance

type TrackCourseInstance struct {
	ProgramCourseID  int
	ProgramID        int
	CourseInstanceID int
}

func NewTrackCourseInstance(
	ProgramCourseID, trackID, instanceID int,
) *TrackCourseInstance {
	return &TrackCourseInstance{

		ProgramID:        trackID,
		CourseInstanceID: instanceID,
	}
}

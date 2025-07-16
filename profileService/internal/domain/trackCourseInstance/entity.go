package trackcourseinstance

type TrackCourseInstance struct {
	TrackID    int
	InstanceID int
}

func NewTrackCourseInstance(
	trackID, instanceID int,
) *TrackCourseInstance {
	return &TrackCourseInstance{
		TrackID:    trackID,
		InstanceID: instanceID,
	}
}

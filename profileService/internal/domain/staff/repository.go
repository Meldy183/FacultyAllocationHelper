package staff

import "context"

type Repository interface {
	GetCourseStaffInstanceByProfileCourseIDs(ctx context.Context, profileId int64, courseInstanceId int64) (*Staff, error)
}

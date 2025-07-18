package courseStaff

import "context"

type Repository interface {
	GetCourseStaffInstanceByProfileCourseIDs(ctx context.Context, profile_id int64, course_instance_id int64) (*CourseStaff, error)
}

package allocationHandler

import "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/staff"

type AllocateFacultyRequest struct {
	CourseID       int64  `json:"course_id"`
	ProfileID      int64  `json:"profile_id"`
	Position_type  string `json:"position_type"`
	GroupsAssigned *int64 `json:"groups_assigned"`
}
type DeallocateFacultyRequest struct {
	CourseID       int64  `json:"course_id"`
	ProfileID      int64  `json:"profile_id"`
	Position_type  string `json:"position_type"`
	GroupsAssigned *int64 `json:"groups_assigned"`
}
type AllocateFacultyResponse struct {
	Staff staff.Staff `json:"staff"`
}

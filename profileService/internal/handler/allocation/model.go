package allocationHandler

type AllocateFacultyRequest struct {
	CourseInstanceID int64  `json:"instance_id"`
	ProfileID        int64  `json:"profile_id"`
	Position_type    string `json:"position_type"`
}

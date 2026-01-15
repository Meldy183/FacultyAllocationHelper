package staff

import "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/sharedContent"

type Staff struct {
	AssignmentID     int64   `json:"assignment_id"`
	InstanceID       int64   `json:"instance_id"`
	ProfileVersionID int64   `json:"profile_version_id"`
	PositionType     *string `json:"position_type"`
	GroupsAssigned   *int64  `json:"groups_assigned"`
	IsConfirmed      bool    `json:"is_confirmed"`
	LecturesCount    *int64  `json:"lectures_count"`
	TutorialsCount   *int64  `json:"tutorials_count"`
	LabsCount        *int64  `json:"labs_count"`
}

func NewStaff(instanceID int64, profileVersionID int64, positionType *string, groupsAssigned *int64) *Staff {
	lectureCount := int64(0)
	tutorialCount := int64(0)
	labCount := int64(0)
	switch *positionType {
	case "PI":
		lectureCount = 15
	case "TI":
		tutorialCount = 15
	case "TA":
		labCount = 15 * *groupsAssigned
	}
	return &Staff{
		InstanceID:       instanceID,
		ProfileVersionID: profileVersionID,
		PositionType:     positionType,
		GroupsAssigned:   groupsAssigned,
		IsConfirmed:      false,
		LecturesCount:    sharedContent.Ptr(lectureCount),
		TutorialsCount:   sharedContent.Ptr(tutorialCount),
		LabsCount:        sharedContent.Ptr(labCount),
	}
}

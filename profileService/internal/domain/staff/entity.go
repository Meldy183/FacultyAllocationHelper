package staff

import "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/sharedContent"

type Staff struct {
	AssignmentID     int64
	InstanceID       int64
	ProfileVersionID int64
	PositionType     *string
	GroupsAssigned   *int64
	IsConfirmed      bool
	LecturesCount    *int64
	TutorialsCount   *int64
	LabsCount        *int64
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

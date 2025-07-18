package semesterworkload

type SemesterWorkload struct {
	SemesterWorkloadID int64
	ProfileVersionID   int64
	SemesterID         int
	LecturesCount      int
	TutorialsCount     int
	LabsCount          int
	ElectivesCount     int
}

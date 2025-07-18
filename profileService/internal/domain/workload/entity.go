package workload

type Workload struct {
	WorkloadID       int64
	ProfileVersionID int64
	SemesterID       int
	LecturesCount    int
	TutorialsCount   int
	LabsCount        int
	ElectivesCount   int
}

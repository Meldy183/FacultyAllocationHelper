package workload

type Workload struct {
	WorkloadID       int64
	ProfileVersionID int64
	SemesterID       int64
	LecturesCount    int64
	TutorialsCount   int64
	LabsCount        int64
	ElectivesCount   int64
	Rate             float64
}

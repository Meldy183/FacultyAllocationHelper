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

func NewWorkload(profileVersionID int64, semesterID int64, lecturesCount int64, tutorialsCount int64, labsCount int64) *Workload {
	return &Workload{
		ProfileVersionID: profileVersionID,
		SemesterID:       semesterID,
		LecturesCount:    lecturesCount,
		TutorialsCount:   tutorialsCount,
		LabsCount:        labsCount,
		ElectivesCount:   0,
		Rate:             0.0,
	}
}

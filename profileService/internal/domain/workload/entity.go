package workload

import "github.com/google/uuid"

type Workload struct {
	WorkloadID       uuid.UUID
	ProfileVersionID uuid.UUID
	SemesterID       uuid.UUID
	LecturesCount    int64
	TutorialsCount   int64
	LabsCount        int64
	ElectivesCount   int64
	Rate             float64
}

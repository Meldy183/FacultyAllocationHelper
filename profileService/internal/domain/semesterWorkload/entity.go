package semesterworkload

type SemesterWorkload struct {
	SemesterWorkloadID int64
	SemesterID         int
	lecturesCount      int
	tutorialsCount     int
	labsCount          int
	electivesCount     int
}

func NewSemesterWorkload(
	semesterWorkloadID int64,
	SemesterID, lecturesCount, tutorialsCount, labsCount, electivesCount int,
) *SemesterWorkload {
	return &SemesterWorkload{
		SemesterWorkloadID: semesterWorkloadID,
		SemesterID:         SemesterID,
		lecturesCount:      lecturesCount,
		tutorialsCount:     tutorialsCount,
		labsCount:          labsCount,
		electivesCount:     electivesCount,
	}
}

package usercourseinstance

type Repository interface {
	GetByProfileID(ProfileID int64) ([]int64, error)
	AddCourseInstance(ProfileID int64, instance int64) error
}

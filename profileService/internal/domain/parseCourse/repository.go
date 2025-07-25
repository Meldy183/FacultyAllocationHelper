package parsecourse

import (
	"context"

	parseuser "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/parseUser"
)

type Repository interface {
	ParseCourses(courses [][]string, ctx context.Context, studyyear int, persons *[]parseuser.Person) (*[]Course, error)
	ParseElectives(electives [][]string, ctx context.Context, studyyear int, persons *[]parseuser.Person) (*[]Course, error)
}

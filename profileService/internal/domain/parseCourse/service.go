package parsecourse

import (
	"context"

	parseuser "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/parseUser"
)

type Service interface {
	ProcessCourse(courses [][]string, elecs [][]string, ctx context.Context, studyyear int, persons *[]parseuser.Person) error
}

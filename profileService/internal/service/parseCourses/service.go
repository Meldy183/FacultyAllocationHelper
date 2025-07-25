package parsecourses

import (
	"context"

	curs "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/course"
	instance "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/courseInstance"
	parsecourse "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/parseCourse"
	parseuser "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/parseUser"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/semester"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/course"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/courseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/program"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/track"
	"go.uber.org/zap"
)

var _ parsecourse.Service = (*Service)(nil)

type Service struct {
	logger          *zap.Logger
	instanceService *courseInstance.Service
	courseService   *course.Service
	trackService    *track.Service
	program         *program.Service
	semester        *semester.Service
	repo            parsecourse.Repository
}

func NewService(instance *courseInstance.Service, course *course.Service, track *track.Service, program *program.Service, logger *zap.Logger, repo parsecourse.Repository) *Service {
	return &Service{logger: logger, instanceService: instance, courseService: course, trackService: track, program: program, repo: repo}

}
func (s *Service) ProcessCourse(courses [][]string, elecs [][]string, ctx context.Context, studyyear int, persons *[]parseuser.Person) error {
	courseEntites, err := s.repo.ParseCourses(courses, ctx, studyyear, persons)
	if err != nil {
		return err
	}
	electiveEntities, err := s.repo.ParseElectives(elecs, ctx, studyyear, persons)
	if err != nil {
		return err
	}
	IsElective := false
	for _, cours := range *courseEntites {
		entity := curs.Course{Name: cours.Name, OfficialName: &cours.OfficialName, IsElective: &IsElective, LecHours: &cours.LecHours,
			LabHours: &cours.LabHours}
		err := s.courseService.AddCourse(ctx, &entity)
		if err != nil {
			return err
		}
		var mode string
		if cours.LabFormat != cours.LectureFormat {
			mode = "partially remote"
		} else if cours.LectureFormat == "remote" {
			mode = "remote"
		} else {
			mode = "onsite"
		}
		groupsTaken := int64(len(cours.TA))
		instance := instance.CourseInstance{Year: int64(studyyear), Form: (*instance.Form)(&cours.Form), Mode: (*instance.Mode)(&mode),
			GroupsTaken: }
	}
	for _, electiv := range *electiveEntities {

	}
}

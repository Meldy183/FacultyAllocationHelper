package CompleteCourse

import (
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/course"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/courseInstance"
)

type FullCourse struct {
	course.Course
	courseInstance.CourseInstance
	StudyPrograms []*string
	Tracks        []*string
}

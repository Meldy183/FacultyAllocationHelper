package Parsing

import (
	"context"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/CompleteCourse"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/CompleteUser"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/courseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/parsing"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/position"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/responsibleInstitute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ parsing.Service = (*Service)(nil)

type Service struct {
	logger                *zap.Logger
	completeUserService   CompleteUser.Service
	completeCourseService CompleteCourse.Service
	positionService       position.Service
	respInstituteService  responsibleInstitute.Service
}

func NewService(logger *zap.Logger,
	completeCourseService CompleteCourse.Service,
	completeUserService CompleteUser.Service,
	positionService position.Service,
	respInstituteService responsibleInstitute.Service) *Service {
	return &Service{logger: logger,
		completeUserService:   completeUserService,
		completeCourseService: completeCourseService,
		positionService:       positionService,
		respInstituteService:  respInstituteService}
}
func (s *Service) Parse(ctx context.Context, file *multipart.File) error {
	f, err := excelize.OpenReader(*file)
	if err != nil {
		s.logger.Error("Unable to open excel file",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogParse),
			zap.Error(err))
		return err
	}
	sheets := f.GetSheetList()
	// rowsElectives, err := f.GetRows(sheets[0])
	// if err != nil {
	// 	s.logger.Error("Failed to read elective sheet",
	// 		zap.String("layer", logctx.LogServiceLayer),
	// 		zap.String("function", logctx.LogParse),
	// 		zap.Error(err))
	// 	return err
	// }
	rowsCourses, err := f.GetRows(sheets[1])
	if err != nil {
		s.logger.Error("Failed to read courses sheet",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogParse),
			zap.Error(err))
		return err
	}
	rowsPersons, err := f.GetRows(sheets[2])
	if err != nil {
		s.logger.Error("Failed to read person sheet",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogParse),
			zap.Error(err))
		return err
	}
	sheetname := strings.ReplaceAll(sheets[1], ".", "/")
	sheetname = strings.ReplaceAll(sheetname, "-", "/")
	years := strings.Split(sheetname, "/")
	yearT1num, err := strconv.Atoi(years[1])
	if err != nil {
		s.logger.Error("unable to parse year of T1",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogParse),
			zap.Error(err))
		return err
	}
	err = s.parseUsers(ctx, rowsPersons)
	if err != nil {
		s.logger.Error("unable to parse users",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogParseCourse),
			zap.Error(err))
		return err
	}
	err = s.parseCourses(rowsCourses, ctx, yearT1num)
	if err != nil {
		s.logger.Error("unable to parse courses",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogParseCourse),
			zap.Error(err))
		return err
	}
	return nil
}

func (s *Service) parseCourses(courses [][]string, ctx context.Context, studyyear int) error {
	var semester int
	var year int
	for ind, rows := range courses {
		var course CompleteCourse.FullCourse
		if ind == 0 || ind == 1 || len(rows) == 0 {
			continue
		} else {
			rows[0] = strings.TrimSpace(rows[0])
			if rows[0] != "факультатив" && rows[0] != "BS" && rows[0] != "MS" && rows[0] != "PhD" && rows[0] != "" {
				yearsem := strings.Split(rows[0], "-")
				var err error
				yearString := strings.Trim(yearsem[0], " ")
				year, err = strconv.Atoi(yearString[len(yearString)-1:])
				if err != nil && err != strconv.ErrSyntax {
					s.logger.Error("unable to parse academic year",
						zap.String("layer", logctx.LogServiceLayer),
						zap.String("function", logctx.LogParseCourse),
						zap.String("year string", yearString),
						zap.Int("row", ind),
						zap.String("fullcell", rows[0]),
						zap.Error(err))
					return err
				}
				semesterString := strings.Trim(yearsem[1], " ")
				semester, err = strconv.Atoi(semesterString[len(semesterString)-1:])
				if err != nil && err != strconv.ErrSyntax {
					s.logger.Error("unable to parse semester",
						zap.String("layer", logctx.LogServiceLayer),
						zap.String("function", logctx.LogParseCourse),
						zap.Error(err))
					return err
				}
				continue
			}
			for indj, cellValue := range rows {
				cells := cellValue
				switch indj {
				case 1:
					course.CourseInstance.AcademicYearID = int64(year)
					course.CourseInstance.SemesterID = int64(semester)
					course.Course.Name = cells
				case 2:
					course.Course.OfficialName = &cells
				case 3:
					st := strings.ReplaceAll(cells, ",", "")
					programms := strings.Split(st, " ")
					programPointers := make([]*string, len(programms))
					for i, program := range programms {
						programPointers[i] = &program
					}
					course.StudyPrograms = programPointers
				case 4:
					st := strings.ReplaceAll(cells, ",", "")
					tracks := strings.Split(st, " ")
					if tracks[0] == "all" {
						course.Tracks = append(course.StudyPrograms, &tracks[0])
					} else {
						trackPointers := make([]*string, len(tracks))
						for i, program := range tracks {
							trackPointers[i] = &program
						}
						course.Tracks = trackPointers
					}
				case 5:
					var form courseInstance.Form
					form = courseInstance.Form(cells)
					course.CourseInstance.Form = &form
				case 6:
					var mode courseInstance.Mode
					mode = courseInstance.Mode(cells)
					course.CourseInstance.Mode = &mode
				//TODO: allocation
				case 26:
					lechours, err := strconv.Atoi(cells)
					if err != nil && cells != "" {
						s.logger.Error("unable to parse lecture hours",
							zap.String("layer", logctx.LogServiceLayer),
							zap.String("function", logctx.LogParseCourse),
							zap.String("cell", cells),
							zap.Int("row index", ind),
							zap.Error(err))
						return err
					}
					lecturehours := int64(lechours)
					course.Course.LecHours = &lecturehours
				case 27:
					labhours, err := strconv.Atoi(cells)
					if err != nil && cells != "" {
						s.logger.Error("unable to parse labratory hours",
							zap.String("layer", logctx.LogServiceLayer),
							zap.String("function", logctx.LogParseCourse),
							zap.Error(err))
						return err
					}
					labratoryhours := int64(labhours)
					course.Course.LecHours = &labratoryhours
				case 28:
					st := strings.ReplaceAll(cells, ".", "/")
					intstituteID, err := s.respInstituteService.GetResponsibleInstituteIDByName(ctx, st)
					if err != nil || intstituteID == nil {
						s.logger.Error("unable to get responsible institute ID by name",
							zap.String("layer", logctx.LogServiceLayer),
							zap.String("function", logctx.LogParseCourse),
							zap.Error(err))
						continue
					}
					course.Course.ResponsibleInstituteID = *intstituteID
				default:
					//TODO: allocation
				}

			}
			if course.Course.OfficialName != nil {
				switch course.CourseInstance.SemesterID {
				case 2, 3:
					course.CourseInstance.Year = int64(studyyear + 1)
				case 1:
					course.CourseInstance.Year = int64(studyyear)
				}
				s.completeCourseService.AddFullCourse(ctx, &course)
			}
		}
	}
	return nil
}

func (s *Service) parseUsers(ctx context.Context, users [][]string) error {
	for ind, rows := range users {
		if ind == 0 || len(rows) < 2 {
			continue
		} else {
			var person CompleteUser.FullUser
			for indj, cellValue := range rows {
				cells := cellValue
				switch indj {
				case 0:
					person.UserProfile.EnglishName = cells
				case 1:
					person.UserProfile.RussianName = &cells
				case 2:
					positionID, err := s.positionService.GetPositionIDByName(ctx, cells)
					if err != nil || positionID == nil {
						s.logger.Error("unable to parse get ID of position by name",
							zap.String("layer", logctx.LogServiceLayer),
							zap.String("function", logctx.LogParseUser),
							zap.Error(err))
						continue
					}
					person.UserProfileVersion.PositionID = *positionID
				case 3:
					person.UserProfileVersion.StudentType = &cells
				// TODO: implement workload parsing
				case 7:
					person.Languages = append(person.Languages, &cells)
				case 8:
					person.UserProfileVersion.Mode = &cells
				case 9:
					mxload, err := strconv.Atoi(cells)
					if err != nil && cells != "" {
						s.logger.Error("unable to convert maxloadd to integer",
							zap.String("layer", logctx.LogServiceLayer),
							zap.String("function", logctx.LogParseUser),
							zap.Error(err))
						return err
					}
					maxload := int64(mxload)
					person.UserProfileVersion.MaxLoad = &maxload
				case 36:
					person.UserProfileVersion.Fsro = &cells
				case 37:
					person.Institutes = append(person.Institutes, &cells)
				case 38:
					ifDegree := (cells == "With")
					person.UserProfileVersion.Degree = &ifDegree
				case 39:
					person.UserProfile.Email = cells
				case 40:
					person.UserProfile.Alias = cells
				case 41:
					person.UserProfileVersion.EmploymentType = &cells
				case 42:
					t, err := time.Parse("01-02-06", cells)
					if err != nil {
						s.logger.Error("unable to parse start date",
							zap.String("layer", logctx.LogServiceLayer),
							zap.String("function", logctx.LogParseUser),
							zap.Error(err))
					}
					person.UserProfile.StartDate = &t
				case 43:
					t, err := time.Parse("01-02-06", cells)
					if err != nil {
						s.logger.Error("unable to parse end date",
							zap.String("layer", logctx.LogServiceLayer),
							zap.String("function", logctx.LogParseUser),
							zap.Error(err))
					}
					person.UserProfile.EndDate = &t
				default:
					continue
				}
			}
			if person.UserProfile.RussianName == nil {
				continue
			}
			err := s.completeUserService.AddFullUser(ctx, &person)
			if err != nil {
				s.logger.Error("error adding user",
					zap.String("layer", logctx.LogServiceLayer),
					zap.String("function", logctx.LogParseUser),
					zap.Error(err))
			}
		}
	}
	return nil
}
func IfSemester(cell string) bool {
	return (cell == "T1" || cell == "T2" || cell == "T3")
}

package postgres

import (
	"context"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	parsecourse "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/parseCourse"
	parseuser "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/parseUser"
	"go.uber.org/zap"
)

var _ parsecourse.Repository = (*ParseCourseRepo)(nil)

type ParseCourseRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewParseCourseRepo(pool *pgxpool.Pool, logger *zap.Logger) *ParseCourseRepo {
	return &ParseCourseRepo{pool: pool, logger: logger}
}
func IfSemester(cell string) bool {
	return (cell == "T1" || cell == "T2" || cell == "T3")
}
func (r *ParseCourseRepo) ParseElectives(electives [][]string, ctx context.Context, studyyear int, persons *[]parseuser.Person) (*[]parsecourse.Course, error) {
	elecs := make([]parsecourse.Course, 0)
	var semester string
	var course_type string
	var year []string
	for ind, rows := range electives {
		var elective1 parsecourse.Course
		var elective2 parsecourse.Course
		elective1.InstituteId = "Not Nedeed"
		elective2.InstituteId = "Not Needed"
		if ind == 0 || len(rows) == 0 {
			continue
		} else {
			for indj, cells := range rows {
				if indj == 0 && IfSemester(cells) {
					semester = cells
					break
				} else if indj == 0 && (cells == "Tech" || cells == "Hum") {
					course_type = cells
					break
				} else if indj == 1 && cells != "" && cells != "..." {
					year = strings.Split(cells, "-")
					if len(year) > 1 {
						year[1] = year[0][:len(year[0])-1] + year[1]
					}
					break
				}
				switch indj {
				case 2:
					if len(year) > 0 {
						elective1.AcademicYear = year[0]
					} else {
						elective1.AcademicYear = ""
					}
					elective1.Semester = semester
					elective1.Type = course_type
					if cells != "" {
						elective1.Name = cells
					}
				case 3:
					elective1.OfficialName = cells
				case 4:
					st := strings.ReplaceAll(cells, ",", "")
					programms := strings.Split(st, " ")
					elective1.Programms = programms
					elective1.Tracks = append(elective1.Tracks, "all")
				case 5:
					elective1.LectureFormat = cells
				case 6:
					elective1.LabFormat = cells
				case 7:
					elective1.PI = cells
				case 8:
					elective1.TI = "Not Nedeed"
					if cells != "" {
						for _, person := range *persons {
							if cells == person.Name {
								elective1.TA = append(elective1.TA, cells)
								break
							}
						}
					}
				case 9:
					if cells != "" {
						for _, person := range *persons {
							if cells == person.Name {
								elective1.TA = append(elective1.TA, cells)
								break
							}
						}
					}
				case 10:
					if cells != "" {
						for _, person := range *persons {
							if cells == person.Name {
								elective1.TA = append(elective1.TA, cells)
								break
							}
						}
					}
				case 11:
					if cells != "" {
						for _, person := range *persons {
							if cells == person.Name {
								elective1.TA = append(elective1.TA, cells)
								break
							}
						}
					}
				case 12:
					lechours, err := strconv.Atoi(cells)
					if err != nil {
						r.logger.Warn("error parsing lec hours",
							zap.Error(err),
							zap.String("lec hours", cells))
					}
					elective1.LecHours = lechours
				case 13:
					labhours, err := strconv.Atoi(cells)
					if err != nil {
						r.logger.Warn("error parsing lab hours",
							zap.Error(err),
							zap.String("lab hours", cells))
					}
					elective1.LabHours = labhours
				}
			}
		}
		if elective1.OfficialName != "" {
			switch elective1.Semester {
			case "T2", "T3":
				elective1.Year = studyyear + 1
			case "T1":
				elective1.Year = studyyear
			}
			elecs = append(elecs, elective1)
			if len(year) > 1 {
				elective2 = elective1
				elective2.AcademicYear = year[1]
				elecs = append(elecs, elective2)
			}
		}

	}
	return &elecs, nil
}

func (r *ParseCourseRepo) ParseCourses(courses [][]string, ctx context.Context, studyyear int, persons *[]parseuser.Person) (*[]parsecourse.Course, error) {
	cours := make([]parsecourse.Course, 0)
	var semester string
	var year string
	for ind, rows := range courses {
		var course parsecourse.Course
		if ind == 0 || ind == 1 || len(rows) == 0 {
			continue
		} else {
			rows[0] = strings.TrimSpace(rows[0])
			if rows[0] != "факультатив" && rows[0] != "BS" && rows[0] != "MS" && rows[0] != "PhD" && rows[0] != "" {
				yearsem := strings.Split(rows[0], "-")
				year = yearsem[0]
				semester = yearsem[1]
				continue
			}
			for indj, cells := range rows {
				switch indj {
				case 1:
					course.AcademicYear = year
					course.Semester = semester
					course.Name = cells
				case 2:
					course.OfficialName = cells
				case 3:
					st := strings.ReplaceAll(cells, ",", "")
					programms := strings.Split(st, " ")
					course.Programms = programms
				case 4:
					st := strings.ReplaceAll(cells, ",", "")
					tracks := strings.Split(st, " ")
					if tracks[0] == "all" {
						course.Tracks = append(course.Programms, "all")
					} else {
						course.Tracks = tracks
					}
				case 5:
					course.Form = cells
				case 6:
					course.LectureFormat = cells
					course.LabFormat = cells
				case 7:
					course.PI = cells
				case 8:
					course.TI = cells
				case 26:
					lechours, err := strconv.Atoi(cells)
					if err != nil {
						r.logger.Warn("error parsing lec hours",
							zap.Error(err),
							zap.String("lec hours", cells))
					}
					course.LecHours = lechours
				case 27:
					labhours, err := strconv.Atoi(cells)
					if err != nil {
						r.logger.Warn("error parsing lab hours",
							zap.Error(err),
							zap.String("lab hours", cells))
					}
					course.LabHours = labhours
				case 28:
					st := strings.ReplaceAll(cells, ".", "/")
					course.InstituteId = st
				default:
					if cells == "Not Needed" || cells == "" {
						continue
					} else {
						for _, person := range *persons {
							if cells == person.Name {
								course.TA = append(course.TA, cells)
								break
							}
						}
					}
				}

			}
			if course.OfficialName != "" {
				switch course.Semester {
				case "T2", "T3":
					course.Year = studyyear + 1
				case "T1":
					course.Year = studyyear
				}
				cours = append(cours, course)
			}
		}
	}
	return &cours, nil
}

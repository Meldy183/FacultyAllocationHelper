package application

import (
	"context"
	"strconv"
	"strings"
	"time"

	"gitlab.pg.innopolis.university/f.markin/fah/Fah-parsing_demo/internal/domain/entities"
	"gitlab.pg.innopolis.university/f.markin/fah/Fah-parsing_demo/internal/infrastructure/logger"
	"go.uber.org/zap"
)

func ParseUsers(users [][]string, ctx context.Context) (*[]entities.Person, error) {
	lg := logger.GetFromContext(ctx)
	persons := make([]entities.Person, 0)
	for ind, rows := range users {
		if ind == 0 || len(rows) < 2 {
			continue
		} else {
			var person entities.Person
			for indj, cells := range rows {
				switch indj {
				case 0:
					person.Name = cells
				case 1:
					person.RussianName = cells
				case 2:
					person.Position = cells
				case 3:
					person.StudentType = cells
				case 4:
					T1, err := strconv.ParseFloat(cells, 64)
					if err != nil {
						lg.Warn(ctx, "error parsing T1",
							zap.Error(err))
					}
					person.Rates = append(person.Rates, T1)
				case 5:
					T2, err := strconv.ParseFloat(cells, 64)
					if err != nil {
						lg.Warn(ctx, "error parsing T2",
							zap.Error(err))
					}
					person.Rates = append(person.Rates, T2)
				case 6:
					T3, err := strconv.ParseFloat(cells, 64)
					if err != nil {
						lg.Warn(ctx, "error parsing T3",
							zap.Error(err))
					}
					person.Rates = append(person.Rates, T3)
				case 7:
					langnum, err := strconv.Atoi(cells)
					if err != nil {
						lg.Warn(ctx, "error parsing language number",
							zap.Error(err))
					}
					person.NumberOfLanguages = langnum
				case 8:
					person.Mode = cells
				case 9:
					mxload, err := strconv.ParseFloat(cells, 64)
					if err != nil {
						lg.Warn(ctx, "error parsing max load",
							zap.Error(err))
					}
					person.Maxload = mxload
				case 37:
					person.Institute = cells
				case 38:
					person.Degree = (cells == "With")
				case 39:
					person.Email = cells
				case 40:
					person.Alias = cells
				case 41:
					lg.Info(ctx, "emp type", zap.String("type", cells), zap.Int("index", indj))
					person.EmploymentType = cells
				case 42:
					t, err := time.Parse("01-02-06", cells)
					if err != nil {
						lg.Warn(ctx, "error parsing start date",
							zap.Error(err),
							zap.String("date", cells))
					}
					person.StartTime = t
				case 43:
					t, err := time.Parse("01-02-06", cells)
					if err != nil {
						lg.Warn(ctx, "error parsing end date",
							zap.Error(err),
							zap.String("date", cells))
					}
					person.EndTime = t
				default:
					continue
				}
			}
			persons = append(persons, person)

		}
	}
	return &persons, nil
}
func IfSemester(cell string) bool {
	return (cell == "T1" || cell == "T2" || cell == "T3")
}
func ParseElectives(electives [][]string, ctx context.Context, studyyear int, persons *[]entities.Person) (*[]entities.Course, error) {
	lg := logger.GetFromContext(ctx)
	elecs := make([]entities.Course, 0)
	var semester string
	var course_type string
	var year []string
	for ind, rows := range electives {
		var elective1 entities.Course
		var elective2 entities.Course
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
						lg.Warn(ctx, "error parsing lec hours",
							zap.Error(err),
							zap.String("lec hours", cells))
					}
					elective1.LecHours = lechours
				case 13:
					labhours, err := strconv.Atoi(cells)
					if err != nil {
						lg.Warn(ctx, "error parsing lab hours",
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

func ParseCourses(courses [][]string, ctx context.Context, studyyear int, persons *[]entities.Person) (*[]entities.Course, error) {
	lg := logger.GetFromContext(ctx)
	cours := make([]entities.Course, 0)
	var semester string
	var year string
	for ind, rows := range courses {
		var course entities.Course
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
						lg.Warn(ctx, "error parsing lec hours",
							zap.Error(err),
							zap.String("lec hours", cells))
					}
					course.LecHours = lechours
				case 27:
					labhours, err := strconv.Atoi(cells)
					if err != nil {
						lg.Warn(ctx, "error parsing lab hours",
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

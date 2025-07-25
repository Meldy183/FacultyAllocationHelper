package postgres

import (
	"context"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	parseuser "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/parseUser"
	"go.uber.org/zap"
)

var _ parseuser.Repository = (*ParseUserRepo)(nil)

type ParseUserRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewParseUserRepo(pool *pgxpool.Pool, logger *zap.Logger) *ParseUserRepo {
	return &ParseUserRepo{pool: pool, logger: logger}
}
func (r *ParseUserRepo) ParseUsers(ctx context.Context, users [][]string) (*[]parseuser.Person, error) {
	persons := make([]parseuser.Person, 0)
	for ind, rows := range users {
		if ind == 0 || len(rows) < 2 {
			continue
		} else {
			var person parseuser.Person
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
						r.logger.Warn("error parsing T1",
							zap.Error(err))
					}
					person.Rates = append(person.Rates, T1)
				case 5:
					T2, err := strconv.ParseFloat(cells, 64)
					if err != nil {
						r.logger.Warn("error parsing T2",
							zap.Error(err))
					}
					person.Rates = append(person.Rates, T2)
				case 6:
					T3, err := strconv.ParseFloat(cells, 64)
					if err != nil {
						r.logger.Warn("error parsing T3",
							zap.Error(err))
					}
					person.Rates = append(person.Rates, T3)
				case 7:
					langnum, err := strconv.Atoi(cells)
					if err != nil {
						r.logger.Warn("error parsing language number",
							zap.Error(err))
					}
					person.NumberOfLanguages = langnum
				case 8:
					person.Mode = cells
				case 9:
					mxload, err := strconv.ParseFloat(cells, 64)
					if err != nil {
						r.logger.Warn("error parsing max load",
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
					r.logger.Info("emp type", zap.String("type", cells), zap.Int("index", indj))
					person.EmploymentType = cells
				case 42:
					t, err := time.Parse("01-02-06", cells)
					if err != nil {
						r.logger.Warn("error parsing start date",
							zap.Error(err),
							zap.String("date", cells))
					}
					person.StartTime = t
				case 43:
					t, err := time.Parse("01-02-06", cells)
					if err != nil {
						r.logger.Warn("error parsing end date",
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

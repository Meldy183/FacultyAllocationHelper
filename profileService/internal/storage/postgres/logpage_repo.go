package postgres

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/logpage"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ logpage.Repository = (*LogpageRepository)(nil)

type LogpageRepository struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

const (
	queryAddLogpage       = `INSERT INTO log (user_id, action, subject_id) VALUES ($1, $2, $3)`
	queryGetLogpagesFirst = `SELECT * FROM log ORDER BY log_id DESC LIMIT $1`
	queryGetLogpagesNext  = `SELECT * FROM log WHERE log_id < $1 ORDER BY log_id DESC LIMIT $2`
)

func NewLogpageRepository(pool *pgxpool.Pool, logger *zap.Logger) *LogpageRepository {
	return &LogpageRepository{pool: pool, logger: logger}
}
func (r *LogpageRepository) AddLogpage(ctx context.Context, logPage *logpage.LogPage) error {
	_, err := r.pool.Exec(ctx, queryAddLogpage, logPage.UserID, logPage.Action, logPage.SubjectID)
	if err != nil {
		r.logger.Error("Error adding logpage",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogAddLogpage),
			zap.Error(err))
		return err
	}
	r.logger.Info("Logpage added successfully",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogAddLogpage))
	return nil
}
func (r *LogpageRepository) GetLogpages(ctx context.Context, last_id string, limit string) ([]*logpage.LogPage, error) {
	if limit == "" {
		limit = "10"
	}
	if last_id == "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			r.logger.Error("Error converting limit to int",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetLogpages),
				zap.Error(err))
			return nil, err
		}
		rows, err := r.pool.Query(ctx, queryGetLogpagesFirst, limitInt)
		if err != nil {
			r.logger.Error("Error getting logpages",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetLogpages),
				zap.Error(err))
			return nil, err
		}
		defer rows.Close()
		var logPages []*logpage.LogPage
		for rows.Next() {
			logPage := &logpage.LogPage{}
			err := rows.Scan(
				&logPage.LogID,
				&logPage.UserID,
				&logPage.Action,
				&logPage.SubjectID,
				&logPage.Timestamp)
			if err != nil {
				r.logger.Error("Error scanning logpage",
					zap.String("layer", logctx.LogRepoLayer),
					zap.String("function", logctx.LogGetLogpages),
					zap.Error(err))
			}
			logPages = append(logPages, logPage)
		}
		if err := rows.Err(); err != nil {
			r.logger.Error("Error iterating rows",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetLogpages),
				zap.Error(err))
			return nil, err
		}
		r.logger.Info("Logpages retrieved successfully",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetLogpages))
		return logPages, nil
	} else {
		lastIDInt, err := strconv.Atoi(last_id)
		if err != nil {
			r.logger.Error("Error converting last_id to int",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetLogpages),
				zap.Error(err))
			return nil, err
		}
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			r.logger.Error("Error converting limit to int",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetLogpages),
				zap.Error(err))
			return nil, err
		}
		rows, err := r.pool.Query(ctx, queryGetLogpagesNext, lastIDInt, limitInt)
		if err != nil {
			r.logger.Error("Error getting logpages",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetLogpages),
				zap.Error(err))
			return nil, err
		}
		defer rows.Close()
		var logPages []*logpage.LogPage
		for rows.Next() {
			logpage := &logpage.LogPage{}
			err := rows.Scan(
				&logpage.LogID,
				&logpage.UserID,
				&logpage.Action,
				&logpage.SubjectID,
				&logpage.Timestamp)
			if err != nil {
				r.logger.Error("Error scanning logpage",
					zap.String("layer", logctx.LogRepoLayer),
					zap.String("function", logctx.LogGetLogpages),
					zap.Error(err))
				return nil, err
			}
			logPages = append(logPages, logpage)
		}
		if err := rows.Err(); err != nil {
			r.logger.Error("Error iterating rows",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetLogpages),
				zap.Error(err))
			return nil, err
		}
		r.logger.Info("Logpages retrieved successfully",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetLogpages))
		return logPages, nil
	}
}

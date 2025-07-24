package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/track"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ track.Repository = (*TrackRepo)(nil)

type TrackRepo struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewTrackRepo(pool *pgxpool.Pool, logger *zap.Logger) *TrackRepo {
	return &TrackRepo{pool: pool, logger: logger}
}

const (
	queryGetTracNameByID = `SELECT track_id, name FROM track WHERE track_id = $1`
	queryGetAllTracks    = `SELECT track_id, name FROM track`
)

func (r *TrackRepo) GetTrackNameByID(ctx context.Context, id int64) (*string, error) {
	row := r.pool.QueryRow(ctx, queryGetTracNameByID, id)
	var trackObj track.Track
	err := row.Scan(&trackObj.TrackID, &trackObj.Name)
	if err != nil {
		r.logger.Error("failed to Get Track Name By ID",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetTrackNameByID),
			zap.Int64("id", id),
			zap.Error(err),
		)
		return nil, fmt.Errorf("GetLanguageByCode: %w", err)
	}
	r.logger.Info("successfully Got TrackName By ID",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetTrackNameByID),
		zap.Int64("id", id),
		zap.String("name", trackObj.Name),
	)
	return &trackObj.Name, nil
}
func (r *TrackRepo) GetAllTracks(ctx context.Context) ([]*track.Track, error) {
	rows, err := r.pool.Query(ctx, queryGetAllTracks)
	if err != nil {
		r.logger.Error("failed to query all tracks",
			zap.String("layer", logctx.LogRepoLayer),
			zap.String("function", logctx.LogGetAllTracks),
			zap.Error(err),
		)
		return nil, fmt.Errorf("get all tracks: %w", err)
	}
	defer rows.Close()
	var tracks []*track.Track
	for rows.Next() {
		if rows.Err() != nil {
			r.logger.Error("failed to query all tracks",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetAllTracks),
				zap.Error(rows.Err()),
			)
			return nil, fmt.Errorf("get all tracks: %w", rows.Err())
		}
		var trackObj track.Track
		err := rows.Scan(
			&trackObj.TrackID,
			&trackObj.Name)
		if err != nil {
			r.logger.Error("failed to query all tracks",
				zap.String("layer", logctx.LogRepoLayer),
				zap.String("function", logctx.LogGetAllTracks),
				zap.Error(err),
			)
			return nil, fmt.Errorf("get all tracks: %w", err)
		}
		tracks = append(tracks, &trackObj)
	}
	r.logger.Info("all tracks returned",
		zap.String("layer", logctx.LogRepoLayer),
		zap.String("function", logctx.LogGetAllTracks),
		zap.Int("count", len(tracks)),
	)
	return tracks, nil
}

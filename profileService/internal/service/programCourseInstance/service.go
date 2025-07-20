package programCourseInstance

import (
	"context"
	"fmt"

	programcourseinstance "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/programCourseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ programcourseinstance.Service = (*Service)(nil)

type Service struct {
	logger                    *zap.Logger
	programCourseInstanceRepo programcourseinstance.Repository
}

func NewService(programCourseInstanceRepo programcourseinstance.Repository, logger *zap.Logger) *Service {
	return &Service{
		logger:                    logger,
		programCourseInstanceRepo: programCourseInstanceRepo,
	}
}

func (s *Service) GetProgramCourseInstancesByCourseID(ctx context.Context, instanceID int64) ([]*programcourseinstance.ProgramCourseInstance, error) {
	programCourseInstances, err := s.programCourseInstanceRepo.GetProgramCourseInstancesByCourseID(ctx, instanceID)
	if err != nil {
		s.logger.Error("Error getting program course instances by instance ids",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetProgramCourseByID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting program course instances by instance ids: %w", err)
	}
	s.logger.Info("successfully got program course instances by instance ids",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetProgramCourseByID),
	)
	return programCourseInstances, nil
}

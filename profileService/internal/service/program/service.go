package program

import (
	"context"
	"fmt"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/program"
	programcourseinstance "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/programCourseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ program.Service = (*Service)(nil)

type Service struct {
	logger            *zap.Logger
	programRepo       program.Repository
	programCourseRepo programcourseinstance.Repository
}

func (s *Service) GetAllPrograms(ctx context.Context) ([]*program.Program, error) {
	return s.programRepo.GetAllPrograms(ctx)
}

func (s *Service) GetProgramNameByID(ctx context.Context, id int) (*string, error) {
	return s.programRepo.GetProgramNameByID(ctx, id)
}

func NewService(programRepo program.Repository, programCourseInstance programcourseinstance.Repository, logger *zap.Logger) *Service {
	return &Service{
		programRepo:       programRepo,
		programCourseRepo: programCourseInstance,
		logger:            logger,
	}
}

func (s *Service) GetProgramNamesByInstanceID(ctx context.Context, instanceID int64) ([]*string, error) {
	instances, err := s.programCourseRepo.GetProgramCourseInstancesByCourseID(ctx, instanceID)
	if err != nil {
		s.logger.Error("failed to retrieve program course instances by instanceID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetProgramNamesByCourseID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to retrieve program course instances by instanceID: %w", err)
	}
	programNames := make([]*string, 0)
	for _, instance := range instances {
		name, err := s.programRepo.GetProgramNameByID(ctx, instance.ProgramID)
		if err != nil {
			s.logger.Error("failed to retrieve program name",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogGetProgramNamesByCourseID),
				zap.Error(err),
			)
			return nil, fmt.Errorf("failed to retrieve program name: %w", err)
		}
		programNames = append(programNames, name)
	}
	return programNames, nil
}

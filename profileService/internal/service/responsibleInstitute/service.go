package responsibleInstitute

import (
	"context"
	"errors"
	"fmt"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/responsibleInstitute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ responsibleInstitute.Service = (*Service)(nil)

type Service struct {
	respInstRepo responsibleInstitute.Repository
	logger       *zap.Logger
}

func NewService(respInstRepo responsibleInstitute.Repository, logger *zap.Logger) *Service {
	return &Service{respInstRepo: respInstRepo, logger: logger}
}
func (s *Service) GetResponsibleInstituteIDByName(ctx context.Context, name string) (*int64, error) {
	id, err := s.respInstRepo.GetResponsibleInstituteIDByName(ctx, name)
	if err != nil {
		s.logger.Error("Error getting responsible_institute name",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetResponsibleInstituteIDByName),
			zap.String("instituteName", name),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting responsible_institute name: %w", err)
	}
	return id, err
}

func (s *Service) GetResponsibleInstituteNameByID(ctx context.Context, instituteID int64) (*string, error) {
	if instituteID <= 0 || instituteID > 8 {
		s.logger.Error("Institute ID must be between 0 and 8",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetResponsibleInstituteNameByID),
			zap.Int64("instituteID", instituteID),
		)
		return nil, errors.New("institute ID must be between 0 and 8")
	}
	name, err := s.respInstRepo.GetResponsibleInstituteNameByID(ctx, instituteID)
	if err != nil {
		s.logger.Error("Error getting responsible_institute name",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetResponsibleInstituteNameByID),
			zap.Int64("instituteID", instituteID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting responsible_institute name: %w", err)
	}
	return name, nil
}

package responsibleInstitute

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
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

func (s *Service) GetResponsibleInstituteIDByName(ctx context.Context, name string) (*uuid.UUID, error) {
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

func (s *Service) GetResponsibleInstituteNameByID(ctx context.Context, instituteID uuid.UUID) (*string, error) {
	if instituteID == uuid.Nil {
		s.logger.Error("Institute ID is invalid",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetResponsibleInstituteNameByID),
			zap.String("instituteID", instituteID.String()),
		)
		return nil, errors.New("institute ID is invalid")
	}
	name, err := s.respInstRepo.GetResponsibleInstituteNameByID(ctx, instituteID)
	if err != nil {
		s.logger.Error("Error getting responsible_institute name",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetResponsibleInstituteNameByID),
			zap.String("instituteID", instituteID.String()),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting responsible_institute name: %w", err)
	}
	return name, nil
}

package responsibleInstitute

import (
	"context"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/responsibleInstitute"
	"go.uber.org/zap"
)

type Service struct {
	respInstRepo responsibleInstitute.Repository
	logger       *zap.Logger
}

func NewService(respInstRepo responsibleInstitute.Repository, logger *zap.Logger) *Service {
	return &Service{respInstRepo: respInstRepo, logger: logger}
}

func (s *Service) GetInstituteNameByID(ctx context.Context, instituteID int64) (*string, error) {
	
}

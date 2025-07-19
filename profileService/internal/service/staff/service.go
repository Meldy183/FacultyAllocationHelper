package staff

import (
	"context"
	"fmt"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/staff"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ staff.Service = (*Service)(nil)

type Service struct {
	repo   staff.Repository
	logger *zap.Logger
}

func NewStaffService(repo staff.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) GetAllStaffByInstanceID(ctx context.Context, instanceID int) ([]*staff.Staff, error) {
	//TODO: validations that important field are not nil
	if instanceID <= 0 {
		return nil, fmt.Errorf("invalid instance id: %d", instanceID)
	}
	return s.repo.GetAllStaffByInstanceID(ctx, instanceID)
}

func (s *Service) AddStaff(ctx context.Context, staff *staff.Staff) error {
	//TODO: validations that important field are not nil
	err := s.repo.AddStaff(ctx, staff)
	if err != nil {
		s.logger.Error("Error adding staff",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddStaff),
			zap.Error(err),
		)
		return fmt.Errorf("failed to add staff: %w", err)
	}
	return nil
}

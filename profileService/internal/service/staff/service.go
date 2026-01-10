package staff

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/staff"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ staff.Service = (*Service)(nil)

type Service struct {
	pool   *pgxpool.Pool // MOCK to avoid import error
	repo   staff.Repository
	logger *zap.Logger
}

func NewStaffService(repo staff.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) GetPI(staffs []*staff.Staff) *staff.Staff {
	var pi *staff.Staff
	for _, elem := range staffs {
		if *elem.PositionType == "PI" {
			pi = elem
			break
		}
	}
	return pi
}

func (s *Service) GetTI(staffs []*staff.Staff) *staff.Staff {
	var ti *staff.Staff
	for _, elem := range staffs {
		if *elem.PositionType == "TI" {
			ti = elem
			break
		}
	}
	return ti
}

func (s *Service) GetTAs(staffs []*staff.Staff) []*staff.Staff {
	var ti []*staff.Staff
	for _, elem := range staffs {
		if *elem.PositionType == "TA" {
			ti = append(ti, elem)
		}
	}
	return ti
}

func (s *Service) GetAllStaffByInstanceID(ctx context.Context, instanceID int64) ([]*staff.Staff, error) {
	//TODO: validations that important field are not nil
	if instanceID <= 0 {
		return nil, fmt.Errorf("invalid instance id: %d", instanceID)
	}
	return s.repo.GetAllStaffByInstanceID(ctx, instanceID)
}

func (s *Service) AddStaff(ctx context.Context, staff *staff.Staff) error {
	//TODO: validations that important field are not nil
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		s.logger.Error("error starting transaction",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddSemesterWorkload),
			zap.Error(err))
		return err
	}
	defer func() {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			s.logger.Error("error rolling back transaction",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAddSemesterWorkload),
				zap.Error(rollbackErr))
		}
	}()
	err = s.repo.AddStaff(ctx, &tx, staff)
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

package profileInstitute

import (
	"context"
	"fmt"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/institute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileInstitute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ profileInstitute.Repository = (*Service)(nil)

type Service struct {
	repo   profileInstitute.Repository
	logger *zap.Logger
}

func NewService(r profileInstitute.Repository, logger *zap.Logger) *Service {
	return &Service{repo: r, logger: logger}
}

func (s *Service) GetUserInstituteByID(ctx context.Context, userID int64) (*institute.Institute, error) {
	if userID <= 0 {
		s.logger.Error("userID must be positive",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetUserInstitute),
			zap.Int64("userID", userID))
		return nil, fmt.Errorf("userID must be positive: %d", userID)
	}
	userInst, err := s.repo.GetUserInstituteByID(ctx, userID)
	if err != nil {
		s.logger.Error("Error getting Institute by ProfileID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetUserInstitute),
			zap.Int64("ID", userID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting institute by ID: %w", err)
	}
	s.logger.Info("User Institute by ID found",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetUserInstitute),
		zap.Int64("ID", userID),
	)
	return userInst, nil
}

func (s *Service) AddUserInstitute(ctx context.Context, userInstitute *profileInstitute.UserInstitute) error {

	err := s.repo.AddUserInstitute(ctx, userInstitute)
	if err != nil {
		s.logger.Error("Error adding Institute to DB",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddUserInstitute),
			zap.Error(err))
		return fmt.Errorf("error adding Institute to DB: %w", err)
	}
	s.logger.Info("User Institute added to DB",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogAddUserInstitute),
	)
	return nil
}

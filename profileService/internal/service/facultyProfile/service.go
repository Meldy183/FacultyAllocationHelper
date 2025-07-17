package facultyProfile

import (
	"context"
	"fmt"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/facultyProfile"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
	"strings"
)

var _ facultyProfile.Repository = (*Service)(nil)

type Service struct {
	repo   facultyProfile.Repository
	logger *zap.Logger
}

func NewService(repo facultyProfile.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) AddProfile(ctx context.Context, profile *facultyProfile.UserProfile) error {
	if !isAliasValid(profile) {
		s.logger.Error(
			"Invalid Alias",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddProfile),
		)
		return fmt.Errorf("invalid alias: %v", profile.Alias)
	}
	err := s.repo.AddProfile(ctx, profile)
	if err != nil {
		s.logger.Error("error creating facultyProfile",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddProfile),
			zap.Error(err))
		return fmt.Errorf("error creaing userProfile %w", err)
	}
	s.logger.Info("user facultyProfile created",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogAddProfile),
	)
	return nil
}
func (s *Service) GetProfileByID(ctx context.Context, profileID int64) (*facultyProfile.UserProfile, error) {
	profile, err := s.repo.GetProfileByID(ctx, profileID)
	if err != nil {
		s.logger.Error("error getting facultyProfile",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetProfileByID),
			zap.Int64("profileID", profileID),
			zap.Error(err))
		return nil, fmt.Errorf("error getting facultyProfile %w", err)
	}
	if !isAliasValid(profile) {
		s.logger.Error(
			"Invalid Alias",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetProfileByID),
		)
		return nil, fmt.Errorf("invalid alias: %v", profile.Alias)
	}
	s.logger.Info("user facultyProfile found",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetProfileByID),
		zap.Int64("profileID", profileID),
		zap.Any("facultyProfile", profile),
	)

	return profile, nil
}
func (s *Service) UpdateProfileByID(ctx context.Context, profile *facultyProfile.UserProfile) error {
	err := s.repo.UpdateProfileByID(ctx, profile)
	if err != nil {
		s.logger.Error("error updating facultyProfile",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogUpdateFaculty),
			zap.Error(err))
		return fmt.Errorf("error updating facultyProfile %w", err)
	}
	if !isAliasValid(profile) {
		s.logger.Error(
			"Invalid Alias",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogUpdateFaculty),
		)
		return fmt.Errorf("invalid alias: %v", profile.Alias)
	}
	s.logger.Info("user facultyProfile updated",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogUpdateFaculty),
	)
	return nil
}

func (s *Service) GetProfilesByFilter(ctx context.Context, institutes []int, positions []int) ([]int64, error) {
	profilesByInst :=
}

func isAliasValid(req *facultyProfile.UserProfile) bool {
	if !strings.Contains(req.Alias, "@") || req.Alias == "" {
		return false
	}
	return true
}

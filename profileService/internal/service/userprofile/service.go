package userprofile

import (
	"context"
	"fmt"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/userprofile"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
	"strings"
)

var _ userprofile.Repository = (*Service)(nil)

type Service struct {
	repo   userprofile.Repository
	logger *zap.Logger
}

func NewService(repo userprofile.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) AddProfile(ctx context.Context, profile *userprofile.UserProfile) error {
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
		s.logger.Error("error creating userprofile",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddProfile),
			zap.Error(err))
		return fmt.Errorf("error creaing userProfile %w", err)
	}
	s.logger.Info("user profile created",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogAddProfile),
	)
	return nil
}
func (s *Service) GetProfileByID(ctx context.Context, profileID int64) (*userprofile.UserProfile, error) {
	profile, err := s.repo.GetProfileByID(ctx, profileID)
	if err != nil {
		s.logger.Error("error getting userprofile",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetProfile),
			zap.Int64("profileID", profileID),
			zap.Error(err))
		return nil, fmt.Errorf("error getting userprofile %w", err)
	}
	if !isAliasValid(profile) {
		s.logger.Error(
			"Invalid Alias",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetProfile),
		)
		return nil, fmt.Errorf("invalid alias: %v", profile.Alias)
	}
	s.logger.Info("user profile found",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetProfile),
		zap.Int64("profileID", profileID),
		zap.Any("profile", profile),
	)

	return profile, nil
}
func (s *Service) UpdateProfileByID(ctx context.Context, profile *userprofile.UserProfile) error {
	err := s.repo.UpdateProfileByID(ctx, profile)
	if err != nil {
		s.logger.Error("error updating userprofile",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogUpdateFaculty),
			zap.Error(err))
		return fmt.Errorf("error updating userprofile %w", err)
	}
	if !isAliasValid(profile) {
		s.logger.Error(
			"Invalid Alias",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogUpdateFaculty),
		)
		return fmt.Errorf("invalid alias: %v", profile.Alias)
	}
	s.logger.Info("user profile updated",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogUpdateFaculty),
	)
	return nil
}

func (s *Service) GetProfilesByFilter(ctx context.Context, institutes []int, positions []int) ([]int64, error) {
	profiles, err := s.repo.GetProfilesByFilter(ctx, institutes, positions)
	if err != nil {
		s.logger.Error("error getting userprofile by filter",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetProfilesByFilters),
			zap.Ints("institutes", institutes),
			zap.Ints("positions", positions),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting userprofile by filter %w", err)
	}
	return profiles, nil
}

func isAliasValid(req *userprofile.UserProfile) bool {
	if !strings.Contains(req.Alias, "@") || req.Alias == "" {
		return false
	}
	return true
}

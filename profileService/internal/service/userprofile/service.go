package userprofile

import (
	"context"
	"fmt"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/userprofile"
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

const (
	logLayer          = "service"
	logCreate         = "create"
	logUpdate         = "update"
	logGetByProfileID = "get_by_profile_id"
	logGetByUserID    = "get_by_user_id"
)

func (s *Service) Create(ctx context.Context, profile *userprofile.UserProfile) error {
	if !isAliasValid(profile) {
		s.logger.Error(
			"Invalid Alias",
			zap.String("layer", logLayer),
			zap.String("function", logCreate),
		)
		return fmt.Errorf("invalid alias: %v", profile.Alias)
	}
	err := s.repo.Create(ctx, profile)
	if err != nil {
		s.logger.Error("error creating userprofile",
			zap.String("layer", logLayer),
			zap.String("function", logCreate),
			zap.Error(err))
		return fmt.Errorf("error creaing userProfile %w", err)
	}
	s.logger.Info("user profile created",
		zap.String("layer", logLayer),
		zap.String("function", logCreate),
	)
	return nil
}
func (s *Service) GetByProfileID(ctx context.Context, profileID int64) (*userprofile.UserProfile, error) {
	profile, err := s.repo.GetByProfileID(ctx, profileID)
	if err != nil {
		s.logger.Error("error getting userprofile",
			zap.String("layer", logLayer),
			zap.String("function", logGetByProfileID),
			zap.Int64("profileID", profileID),
			zap.Error(err))
		return nil, fmt.Errorf("error getting userprofile %w", err)
	}
	if !isAliasValid(profile) {
		s.logger.Error(
			"Invalid Alias",
			zap.String("layer", logLayer),
			zap.String("function", logCreate),
		)
		return nil, fmt.Errorf("invalid alias: %v", profile.Alias)
	}
	s.logger.Info("user profile found",
		zap.String("layer", logLayer),
		zap.String("function", logGetByProfileID),
		zap.Int64("profileID", profileID),
		zap.Any("profile", profile),
	)

	return profile, nil
}
func (s *Service) Update(ctx context.Context, profile *userprofile.UserProfile) error {
	err := s.repo.Update(ctx, profile)
	if err != nil {
		s.logger.Error("error updating userprofile",
			zap.String("layer", logLayer),
			zap.String("function", logUpdate),
			zap.Error(err))
		return fmt.Errorf("error updating userprofile %w", err)
	}
	if !isAliasValid(profile) {
		s.logger.Error(
			"Invalid Alias",
			zap.String("layer", logLayer),
			zap.String("function", logCreate),
		)
		return fmt.Errorf("invalid alias: %v", profile.Alias)
	}
	s.logger.Info("user profile updated",
		zap.String("layer", logLayer),
		zap.String("function", logUpdate),
	)
	return nil
}

func isAliasValid(req *userprofile.UserProfile) bool {
	if !strings.Contains(req.Alias, "@") || req.Alias == "" {
		return false
	}
	return true
}

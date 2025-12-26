package facultyProfile

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/facultyProfile"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ facultyProfile.Service = (*Service)(nil)

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
func (s *Service) GetProfileByID(ctx context.Context, profileID uuid.UUID) (*facultyProfile.UserProfile, error) {
	profile, err := s.repo.GetProfileByID(ctx, profileID)
	if err != nil {
		s.logger.Error("error getting facultyProfile",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetProfileByID),
			zap.String("profileID", profileID.String()),
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
		zap.String("profileID", profileID.String()),
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

func (s *Service) GetProfilesByFilters(ctx context.Context, institutes []uuid.UUID, positions []uuid.UUID) ([]uuid.UUID, error) {
	profilesByInst, err := s.repo.GetProfileIDsByInstituteIDs(ctx, institutes)
	if err != nil {
		s.logger.Error("error getting facultyProfile",
			zap.String("layer", logctx.LogServiceLayer),
			zap.Any("institutes", institutes),
			zap.Any("positions", positions),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting facultyProfile %w", err)
	}
	profilesByInst = makeUnique(profilesByInst)
	s.logger.Debug("Check institutes by filters",
		zap.Any("institutesProfileIDs", profilesByInst),
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetProfileByID),
	)
	profilesByPosition, err := s.repo.GetProfileIDsByPositionIDs(ctx, positions)
	profilesByPosition = makeUnique(profilesByPosition)
	if err != nil {
		s.logger.Error("error getting facultyProfile",
			zap.String("layer", logctx.LogServiceLayer),
			zap.Any("positions", positions),
			zap.Any("institutes", institutes),
			zap.Error(err),
		)
	}
	s.logger.Warn("Check positions by filters",
		zap.Any("positionsProfileIDs", profilesByPosition),
		zap.Any("positions", positions),
		zap.Any("institutesProfileIDs", profilesByInst),
		zap.Any("institutes", institutes),
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetProfileByID),
	)
	union := getUnion(profilesByInst, profilesByPosition)
	return union, nil
}

func isAliasValid(req *facultyProfile.UserProfile) bool {
	if !strings.Contains(req.Alias, "@") || req.Alias == "" {
		return false
	}
	return true
}

func getUnion(arr1 []uuid.UUID, arr2 []uuid.UUID) []uuid.UUID {
	set := make(map[uuid.UUID]bool)
	for _, v := range arr1 {
		set[v] = true
	}
	ans := make([]uuid.UUID, 0)
	for _, v := range arr2 {
		if set[v] {
			ans = append(ans, v)
		}
	}
	return ans
}

func makeUnique(arr []uuid.UUID) []uuid.UUID {
	seen := make(map[uuid.UUID]bool)
	ans := make([]uuid.UUID, 0)
	for _, v := range arr {
		if !seen[v] {
			seen[v] = true
			ans = append(ans, v)
		}
	}
	return ans
}

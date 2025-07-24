package facultyProfile

import (
	"context"
	"fmt"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/facultyProfile"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
	"strings"
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

func (s *Service) GetProfilesByFilters(ctx context.Context, institutes []int64, positions []int64) ([]int64, error) {
	if len(institutes) == 0 {
		institutes = []int64{1, 2, 3, 4, 5}
	}
	if len(positions) == 0 {
		positions = []int64{1, 2, 3, 4, 5, 6, 7}
	}
	profilesByInst, err := s.repo.GetProfileIDsByInstituteIDs(ctx, institutes)
	if err != nil {
		s.logger.Error("error getting facultyProfile",
			zap.String("layer", logctx.LogServiceLayer),
			zap.Int64s("institutes", institutes),
			zap.Int64s("positions", positions),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting facultyProfile %w", err)
	}
	profilesByInst = makeUnique(profilesByInst)
	s.logger.Debug("Check institutes by filters",
		zap.Int64s("institutesProfileIDs", profilesByInst),
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetProfileByID),
	)
	profilesByPosition, err := s.repo.GetProfileIDsByPositionIDs(ctx, positions)
	profilesByPosition = makeUnique(profilesByPosition)
	if err != nil {
		s.logger.Error("error getting facultyProfile",
			zap.String("layer", logctx.LogServiceLayer),
			zap.Int64s("positions", positions),
			zap.Int64s("institutes", institutes),
			zap.Error(err),
		)
	}
	s.logger.Warn("Check positions by filters",
		zap.Int64s("positionsProfileIDs", profilesByPosition),
		zap.Int64s("positions", positions),
		zap.Int64s("institutesProfileIDs", profilesByInst),
		zap.Int64s("institutes", institutes),
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

func getUnion(arr1 []int64, arr2 []int64) []int64 {
	ans := make([]int64, 0)
	cnt1 := 0
	cnt2 := 0
	for cnt1 < len(arr1) && cnt2 < len(arr2) {
		if arr1[cnt1] < arr2[cnt2] {
			cnt1++
			continue
		}
		if arr1[cnt1] > arr2[cnt2] {
			cnt2++
			continue
		}
		ans = append(ans, arr1[cnt1])
		cnt1++
		cnt2++
	}
	return ans
}

func makeUnique(arr []int64) []int64 {
	ans := make([]int64, 0)
	if len(arr) == 0 {
		return ans
	}
	ans = append(ans, arr[0])
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			continue
		}
		ans = append(ans, arr[i])
	}
	return ans
}

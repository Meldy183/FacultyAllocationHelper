package institute

import (
	"context"
	"fmt"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/institute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ institute.Service = (*Service)(nil)

type Service struct {
	repo   institute.Repository
	logger *zap.Logger
}

func NewService(repo institute.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}
func (s *Service) GetInstituteIDByName(ctx context.Context, instituteName string) (*int64, error) {
	instituteID, err := s.repo.GetInstituteIDByName(ctx, instituteName)
	if err != nil {
		s.logger.Error("failed to retrieve institute by LabID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetInstituteByID),
			zap.String("instituteName", instituteName),
			zap.Error(err),
		)
		return nil, err
	}
	s.logger.Info("Successfully retrieved institute: ",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetInstituteByID),
		zap.String("instituteName:", instituteName),
	)
	return instituteID, nil
}
func (s *Service) GetInstituteByID(ctx context.Context, instituteID int64) (*institute.Institute, error) {
	if instituteID <= 0 {
		s.logger.Error("institute_id is invalid",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetInstituteByID))
		return nil, fmt.Errorf("invalid institute_id: %d", instituteID)
	}
	instituteByID, err := s.repo.GetInstituteByID(ctx, instituteID)
	if err != nil {
		s.logger.Error("failed to retrieve institute by LabID",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetInstituteByID),
			zap.Int64("institute_id", instituteID),
			zap.Error(err),
		)
		return nil, err
	}
	s.logger.Info("Successfully retrieved institute: ",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetInstituteByID),
		zap.Int64("institute_id:", instituteID),
	)
	return instituteByID, nil
}

func (s *Service) GetAllInstitutes(ctx context.Context) ([]*institute.Institute, error) {
	institutes, err := s.repo.GetAllInstitutes(ctx)
	if err != nil {
		s.logger.Error("failed to get all institutes",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetAllInstitutes),
			zap.Error(err),
		)
		return nil, err
	}
	s.logger.Info("Successfully got all",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetAllInstitutes),
		zap.Int64("institutes", int64(len(institutes))),
	)
	return institutes, nil
}

func ConvertInstitutesToString(institutes []*institute.Institute) *[]string {
	stringArray := make([]string, 0)
	for _, instituteStruct := range institutes {
		str := instituteStruct.Name
		stringArray = append(stringArray, str)
	}
	return &stringArray
}

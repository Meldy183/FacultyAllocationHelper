package language

import (
	"context"
	"fmt"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/language"
	"go.uber.org/zap"
)

var _ language.Repository = (*Service)(nil)

type Service struct {
	repo   language.Repository
	logger *zap.Logger
}

const (
	layer     = "Service"
	getAll    = "GetAll"
	getByCode = "GetByCode"
)

func NewService(repo language.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

func (s *Service) GetAll(ctx context.Context) ([]*language.Language, error) {
	languages, err := s.repo.GetAll(ctx)
	if err != nil {
		s.logger.Error("Failed to get all languages",
			zap.String("layer", layer),
			zap.String("function", getAll),
			zap.Error(err))
		return nil, fmt.Errorf("failed to get all languages. error: %w", err)
	}
	s.logger.Info("Successfully got all languages",
		zap.String("layer", layer),
		zap.String("function", getAll),
	)
	return languages, nil
}

func (s *Service) GetByCode(ctx context.Context, code string) (*language.Language, error) {
	lang, err := s.repo.GetByCode(ctx, code)
	if err != nil {
		s.logger.Error("Failed to get lang by code",
			zap.String("layer", layer),
			zap.String("function", getByCode),
			zap.String("code", code),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to get lang by code: %s. error: %w", code, err)
	}
	s.logger.Info("Successfully got lang by code",
		zap.String("layer", layer),
		zap.String("function", getByCode),
		zap.String("code", code),
	)
	return lang, nil
}

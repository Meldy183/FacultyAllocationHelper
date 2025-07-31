package language

import (
	"context"
	"fmt"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/language"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ language.Service = (*Service)(nil)

type Service struct {
	repo   language.Repository
	logger *zap.Logger
}

func NewService(repo language.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}
func (s *Service) GetAllLanguages(ctx context.Context) ([]string, error) {
	languages, err := s.repo.GetAllLanguages(ctx)
	if err != nil {
		s.logger.Error("Failed to get all languages",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetAllLanguages),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to get all languages. error: %w", err)
	}
	s.logger.Info("Successfully got all languages",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetAllLanguages),
	)
	return languages, nil
}
func (s *Service) GetCodeByLanguageName(ctx context.Context, name string) (*string, error) {
	languageCode, err := s.repo.GetCodeByLanguageName(ctx, name)
	if err != nil {
		s.logger.Error("Failed to get lang by code",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetCodeByLanguageName),
			zap.String("name", name),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to get lang by code: %s. error: %w", name, err)
	}
	s.logger.Info("Successfully got lang by code",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetCodeByLanguageName),
		zap.String("name", name),
	)
	return languageCode, nil
}
func (s *Service) GetLanguageByCode(ctx context.Context, code string) (*language.Language, error) {
	lang, err := s.repo.GetLanguageByCode(ctx, code)
	if err != nil {
		s.logger.Error("Failed to get lang by code",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetLanguageByCode),
			zap.String("code", code),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to get lang by code: %s. error: %w", code, err)
	}
	s.logger.Info("Successfully got lang by code",
		zap.String("layer", logctx.LogServiceLayer),
		zap.String("function", logctx.LogGetLanguageByCode),
		zap.String("code", code),
	)
	return lang, nil
}

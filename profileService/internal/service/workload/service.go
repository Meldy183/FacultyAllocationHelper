package workload

import (
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/workload"
	"go.uber.org/zap"
)

type Service struct {
	repo   *workload.Repository
	logger *zap.Logger
}

func NewService(repo *workload.Repository, logger *zap.Logger) *Service {
	return &Service{repo: repo, logger: logger}
}

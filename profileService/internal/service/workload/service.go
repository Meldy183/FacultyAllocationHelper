package workload

import (
	"context"
	"errors"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/transaction"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/workload"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

var _ workload.Service = (*Service)(nil)

type Service struct {
	uow    transaction.UnitOfWork // MOCK to avoid import error
	repo   workload.Repository
	logger *zap.Logger
}

func NewService(uow transaction.UnitOfWork, repo workload.Repository, logger *zap.Logger) *Service {
	return &Service{uow: uow, repo: repo, logger: logger}
}

func (s *Service) GetSemesterWorkloadByVersionID(
	ctx context.Context,
	profileVersionID int64,
	semesterID int64,
) (*workload.Workload, error) {
	if semesterID < 1 || semesterID > 3 {
		s.logger.Error(`semesterID must be between 1 and 3`,
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetSemesterWorkloadByVersionID),
		)
		return nil, errors.New(`semesterID must be between 1 and 3`)
	}
	semWorkload, err := s.repo.GetSemesterWorkloadByVersionID(ctx, profileVersionID, semesterID)
	if err != nil {
		s.logger.Error(`GetSemesterWorkloadByVersionID failed`,
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetSemesterWorkloadByVersionID),
			zap.Error(err),
		)
		return nil, errors.New(`GetSemesterWorkloadByVersionID failed`)
	}
	return semWorkload, nil
}
func (s *Service) AddSemesterWorkload(ctx context.Context, workload *workload.Workload) error {

	err := s.uow.Do(ctx, func(tx transaction.Transaction) error {
		err := s.repo.AddSemesterWorkload(ctx, tx, workload)
		if err != nil {
			s.logger.Error(`AddSemesterWorkload failed`,
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAddSemesterWorkload),
				zap.Error(err),
			)
			return errors.New(`AddSemesterWorkload failed`)
		}
		return err
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) UpdateSemesterWorkload(ctx context.Context, workload *workload.Workload) error {

	err := s.uow.Do(ctx, func(tx transaction.Transaction) error {
		err := s.repo.UpdateSemesterWorkload(ctx, tx, workload)
		if err != nil {
			s.logger.Error(`UpdateSemesterWorkload failed`,
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogUpdateSemesterWorkload),
				zap.Error(err),
			)
			return errors.New(`UpdateSemesterWorkload failed`)
		}
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

package service

import "go.uber.org/zap"

type Service struct {
	repo
	logger *zap.Logger
}

func NewService(logger *zap.Logger) *Service {

}

package userprofile

import (
	"github.com/go-chi/chi/v5"
	userprofile "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/userprofile"
	"go.uber.org/zap"
)

type Handler struct {
	service *userprofile.Service
	logger  *zap.Logger
}

func NewHandler(service *userprofile.Service, logger *zap.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger.Named("userprofile_handler"),
	}
}

func (handler *Handler) RegisterRoutes(r chi.Router) {
	r.Route("/")
}

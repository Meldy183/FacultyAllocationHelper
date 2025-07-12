package http

import (
	"github.com/go-chi/chi/v5"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/userprofile"
)

func NewRouter(h *userprofile.Handler) chi.Router {
	r := chi.NewRouter()
	userprofile.RegisterRoutes(r, h)
	return r
}

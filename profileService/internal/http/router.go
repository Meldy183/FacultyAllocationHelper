package http

import (
	"github.com/go-chi/chi/v5"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/facultyProfile"
)

func NewRouter(h *facultyProfile.Handler) chi.Router {
	r := chi.NewRouter()
	facultyProfile.RegisterRoutes(r, h)
	return r
}

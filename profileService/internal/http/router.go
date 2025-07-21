package http

import (
	"github.com/go-chi/chi/v5"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/courses"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/facultyProfile"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/filters"
)

func NewRouter(facultyHandler *facultyProfile.Handler, coursesHandler *courses.Handler, filtersHandler *filters.Handler) chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Route("/profile", func(r chi.Router) {
			facultyProfile.RegisterRoutes(r, facultyHandler)
		})
		r.Route("/academic", func(r chi.Router) {
			courses.RegisterRoutes(r, coursesHandler)
		})
		r.Route("/filter", func(r chi.Router) {
			filters.RegisterRoutes(r, filtersHandler)
		})
	})

	return r
}

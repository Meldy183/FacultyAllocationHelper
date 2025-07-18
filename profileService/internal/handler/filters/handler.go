package filters

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/institute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/position"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	logger           *zap.Logger
	servicePosition  *position.Service
	serviceInstitute *institute.Service
}

func (h *Handler) GetFacultyFilters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	positions, err := h.servicePosition.GetAllPositions(ctx)
	if err != nil {
		h.logger.Error("Error getting all positions",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetFacultyFilters),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting all positions")
		return
	}

	institutes, err := h.serviceInstitute.GetAllInstitutes(ctx)
	if err != nil {
		h.logger.Error("Error getting all institutes",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetFacultyFilters),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting all institutes")
		return
	}
	instituteObjects := make([]FilterObj, 0)
	for _, elem := range institutes {
		instituteObject := &FilterObj{
			ID:   elem.InstituteID,
			Name: elem.Name,
		}
		instituteObjects = append(instituteObjects, *instituteObject)
	}

	positionObjects := make([]FilterObj, 0)
	for _, elem := range positions {
		positionObject := &FilterObj{
			ID:   elem.PositionID,
			Name: elem.Name,
		}
		positionObjects = append(positionObjects, *positionObject)
	}
	facultyFiltersResponse := GetFacultyFiltersResponse{
		InstituteFilters: instituteObjects,
		PositionFilters:  positionObjects,
	}

	writeJSON(w, http.StatusOK, facultyFiltersResponse)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func RegisterRoutes(router chi.Router, h *Handler) {
	router.Route("/", func(r chi.Router) {
		r.Get("/filters", h.GetFacultyFilters)
	})
}

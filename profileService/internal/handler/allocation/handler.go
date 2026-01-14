package allocationHandler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/allocation"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

type Handler struct {
	allocationService allocation.Service
	logger            *zap.Logger
}

func NewHandler(allocationService allocation.Service, logger *zap.Logger) *Handler {
	return &Handler{
		allocationService: allocationService,
		logger:            logger,
	}
}
func (h *Handler) AllocateFaculty(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req AllocateFacultyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("error decoding AllocateFaculty request",
			zap.String("layer", "handler"),
			zap.String("function", "AllocateFaculty"),
			zap.Error(err))
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if req.Position_type != "TA" && req.GroupsAssigned != nil || req.Position_type == "TA" && req.GroupsAssigned == nil {
		h.logger.Error("invalid request body",
			zap.String("layer", "handler"),
			zap.String("function", "AllocateFaculty"))
		writeError(w, http.StatusBadRequest, "invalid request")
		return
	}
	err := h.allocationService.AllocateFaculty(ctx, req.CourseID, req.ProfileID, &req.Position_type, req.GroupsAssigned)
	if err != nil {
		h.logger.Error("error allocating faculty",
			zap.String("layer", "handler"),
			zap.String("function", "AllocateFaculty"),
			zap.Error(err))
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	resp := AllocateFacultyResponse{}
	h.logger.Info("success allocating profile",
		zap.String("layer", logctx.LogHandlerLayer),
		zap.String("function", logctx.LogAllocateFaculty),
	)
	writeJSON(w, http.StatusOK, resp)
}
func (h *Handler) DeallocateFaculty(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req DeallocateFacultyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("error decoding AllocateFaculty request",
			zap.String("layer", "handler"),
			zap.String("function", "AllocateFaculty"),
			zap.Error(err))
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if req.Position_type != "TA" && req.GroupsAssigned != nil || req.Position_type == "TA" && req.GroupsAssigned == nil {
		h.logger.Error("invalid request body",
			zap.String("layer", "handler"),
			zap.String("function", "AllocateFaculty"))
		writeError(w, http.StatusBadRequest, "invalid request")
		return
	}
	err := h.allocationService.DeallocateFaculty(ctx, req.CourseID, req.ProfileID, &req.Position_type, req.GroupsAssigned)
	if err != nil {
		h.logger.Error("error allocating faculty",
			zap.String("layer", "handler"),
			zap.String("function", "AllocateFaculty"),
			zap.Error(err))
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	resp := DeallocateFacultyRequest{}
	h.logger.Info("success deallocating profile",
		zap.String("layer", logctx.LogHandlerLayer),
		zap.String("function", logctx.LogDeallocateFaculty),
	)
	writeJSON(w, http.StatusOK, resp)
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
		r.Post("/allocate", h.AllocateFaculty)
		r.Delete("/allocate", h.DeallocateFaculty)
	})
}

package parse

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/parsing"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

type Handler struct {
	logger       *zap.Logger
	parseService parsing.Service
}

func NewHandler(logger *zap.Logger, parser parsing.Service) *Handler {
	return &Handler{logger: logger, parseService: parser}
}
func (h *Handler) ParseExcel(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	file, _, err := r.FormFile("file")
	if err != nil {
		h.logger.Error("Error opening file",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogParseExcel),
			zap.Error(err),
		)
		writeError(w, http.StatusBadRequest, "Error opening file")
		return
	}
	defer file.Close()
	err = h.parseService.Parse(ctx, &file)
	if err != nil {
		h.logger.Error("Error parsing file",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogParseExcel),
			zap.Error(err),
		)
		writeError(w, http.StatusBadRequest, "Error parsing file")
		return
	}
	resp := &ParseResponce{
		Message: "Success",
	}
	h.logger.Info("ParseExcelSuccess",
		zap.String("layer", logctx.LogHandlerLayer),
		zap.String("function", logctx.LogParseExcel),
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
		r.Post("/parse", h.ParseExcel)
	})
}

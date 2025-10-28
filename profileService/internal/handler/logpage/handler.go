package logpage

//
// import (
//	"encoding/json"
//	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/logpage"
//	"net/http"
//	"time"
//
//	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
//	"go.uber.org/zap"
//)
//
// type Handler struct {
//	serviceLogpage *logpage.Service
//	logger         *zap.Logger
//}
//
// func writeJSON(w http.ResponseWriter, status int, data interface{}) {
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(status)
//	_ = json.NewEncoder(w).Encode(data)
//}
//func writeError(w http.ResponseWriter, status int, message string) {
//	writeJSON(w, status, map[string]string{"error": message})
//}
//func NewHandler(ServiceLogpage *logpage.Service, Logger *zap.Logger) *Handler {
//	return &Handler{serviceLogpage: ServiceLogpage, logger: Logger}
//}
//func (h *Handler) GetLogpages(w http.ResponseWriter, r *http.Request) {
//	ctx := r.Context()
//	limit := r.URL.Query().Get("limit")
//	last_id := r.URL.Query().Get("last_id")
//	logpages, err := h.serviceLogpage.GetLogpages(ctx, last_id, limit)
//	if err != nil {
//		h.logger.Error("error converting to int the position",
//			zap.String("layer", logctx.LogHandlerLayer),
//			zap.String("function", logctx.LogGetLogPages),
//			zap.Error(err),
//		)
//		writeError(w, http.StatusInternalServerError, "error institute id")
//		return
//	}
//	var response []Logpage
//	for _, page := range logpages {
//		timestring := page.Timestamp.UTC().Format(time.RFC3339)
//		response = append(response, Logpage{
//			LogID:     page.LogID,
//			UserID:    page.UserID,
//			Action:    page.Action,
//			SubjectID: page.SubjectID,
//			CreatedAt: timestring,
//		})
//	}
//	writeJSON(w, http.StatusOK, response)
//}

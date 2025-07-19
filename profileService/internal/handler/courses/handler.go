package courses

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Handler struct {
	logger *zap.Logger
}

func NewHandler(logger *zap.Logger) *Handler {
	return &Handler{logger}
}

func (h *Handler) GetAllCoursesByFilters(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) AddNewCourse(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetCourse(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		h.logger.Error("error getting course id",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetCourseByID),
			zap.Error(err),
		)
		writeError(w, http.StatusBadRequest, "Invalid course id")
		return
	}

}

func UniteIDs(a []int64, b []int64) *[]int64 {
	union := make([]int64, 0)
	p1 := 0
	p2 := 0
	for p1 < len(a) && p2 < len(b) {
		if a[p1] == b[p2] {
			union = append(union, a[p1])
			p1++
			p2++
			continue
		}
		if a[p1] < b[p2] {
			p1++
		} else {
			p2++
		}
	}
	return &union
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
		r.Get("/getCourse/{id}", h.GetCourse)
	})
}

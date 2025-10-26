package workload

import (
	"context"
	"encoding/json"
	"net/http"

	workloadDomain "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/workload"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/workload"
	"go.uber.org/zap"
)

type Handler struct {
	logger          *zap.Logger
	serviceWorkload *workload.Service
}

func NewWorkloadHandler(
	serviceWorkload *workload.Service,
	logger *zap.Logger,
) *Handler {
	return &Handler{serviceWorkload: serviceWorkload, logger: logger}
}

func (h *Handler) WorkloadToClasses(
	sem1 *workloadDomain.Workload,
	sem2 *workloadDomain.Workload,
	sem3 *workloadDomain.Workload,
) *Stats {
	class1 := &Classes{
		Lec:  sem1.LecturesCount,
		Tut:  sem1.TutorialsCount,
		Lab:  sem1.LabsCount,
		Elec: sem1.ElectivesCount,
		Rate: sem1.Rate,
	}
	class2 := &Classes{
		Lec:  sem2.LecturesCount,
		Tut:  sem2.TutorialsCount,
		Lab:  sem2.LecturesCount,
		Elec: sem2.ElectivesCount,
		Rate: sem2.Rate,
	}
	class3 := &Classes{
		Lec:  sem3.LecturesCount,
		Tut:  sem3.TutorialsCount,
		Lab:  sem3.LecturesCount,
		Elec: sem3.ElectivesCount,
		Rate: sem3.Rate,
	}
	stats := &Stats{
		T1: *class1,
		T2: *class2,
		T3: *class3,
	}
	return stats
}

func (h *Handler) GetYearWorkload(
	w http.ResponseWriter,
	err error,
	ctx context.Context,
	versionID int64,
) (*workloadDomain.Workload, *workloadDomain.Workload, *workloadDomain.Workload, bool) {
	sem1, err := h.serviceWorkload.GetSemesterWorkloadByVersionID(ctx, versionID, 1)
	if err != nil {
		h.logger.Error(`GetSemesterWorkloadByVersionID failed`,
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetYearWorkloadByVersionID),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting semester workload by version id")
		return nil, nil, nil, true
	}
	sem2, err := h.serviceWorkload.GetSemesterWorkloadByVersionID(ctx, versionID, 2)
	if err != nil {
		h.logger.Error(`GetSemesterWorkloadByVersionID failed`,
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetYearWorkloadByVersionID),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting semester workload by version id")
		return nil, nil, nil, true
	}
	sem3, err := h.serviceWorkload.GetSemesterWorkloadByVersionID(ctx, versionID, 3)
	if err != nil {
		h.logger.Error(`GetSemesterWorkloadByVersionID failed`,
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogGetYearWorkloadByVersionID),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting semester workload by version id")
		return nil, nil, nil, true
	}
	return sem1, sem2, sem3, false
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

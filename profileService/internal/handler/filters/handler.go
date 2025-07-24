package filters

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/academicYear"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/institute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/position"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/program"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/responsibleInstitute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/semester"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	logger                      *zap.Logger
	positionService             position.Service
	instituteService            institute.Service
	academicYearService         academicYear.Service
	programService              program.Service
	semesterService             semester.Service
	responsibleInstituteService responsibleInstitute.Service
}

func NewHandler(
	positionService position.Service,
	instituteService institute.Service,
	academicYearService academicYear.Service,
	programService program.Service,
	semesterService semester.Service,
	responsibleInstituteService responsibleInstitute.Service,
	logger *zap.Logger,
) *Handler {
	return &Handler{
		logger:                      logger,
		positionService:             positionService,
		instituteService:            instituteService,
		academicYearService:         academicYearService,
		programService:              programService,
		semesterService:             semesterService,
		responsibleInstituteService: responsibleInstituteService,
	}
}

func (h *Handler) GetFacultyFilters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	positions, err := h.positionService.GetAllPositions(ctx)
	if err != nil {
		h.logger.Error("Error getting all positions",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetFacultyFilters),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting all positions")
		return
	}
	institutes, err := h.instituteService.GetAllInstitutes(ctx)
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
	resp2 := FiltersFaculty{
		Filters: facultyFiltersResponse,
	}
	writeJSON(w, http.StatusOK, resp2)
}

func (h *Handler) GetCoursesFilters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	years, err := h.academicYearService.GetAllAcademicYears(ctx)
	if err != nil {
		h.logger.Error("Error getting all years",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetAllAcademicYears),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting all years")
		return
	}
	semesters, err := h.semesterService.GetAllSemesters(ctx)
	if err != nil {
		h.logger.Error("Error getting all semesters",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetAllSemesters),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting all semesters")
		return
	}
	programs, err := h.programService.GetAllPrograms(ctx)
	if err != nil {
		h.logger.Error("Error getting all programs",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetAllPrograms),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting all programs")
		return
	}
	instst, err := h.instituteService.GetAllInstitutes(ctx)
	if err != nil {
		h.logger.Error("Error getting all institutes",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetAllInstitutes),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting all institutes")
		return
	}
	yearsOfStudy := make([]FilterObj, 0)
	for _, elem := range years {
		obj := FilterObj{
			ID:   elem.YearID,
			Name: elem.Name,
		}
		yearsOfStudy = append(yearsOfStudy, obj)
	}
	sem := make([]FilterObj, 0)
	for _, elem := range semesters {
		obj := FilterObj{
			ID:   elem.SemesterID,
			Name: elem.Name,
		}
		sem = append(sem, obj)
	}
	studyProgram := make([]FilterObj, 0)
	for _, elem := range programs {
		obj := FilterObj{
			ID:   elem.ProgramID,
			Name: elem.Name,
		}
		studyProgram = append(studyProgram, obj)
	}
	institutes := make([]FilterObj, 0)
	for _, elem := range instst {
		obj := FilterObj{
			ID:   elem.InstituteID,
			Name: elem.Name,
		}
		institutes = append(institutes, obj)
	}
	resp := GetCourseFiltersResponse{
		AllocationStatus: []bool{true, false},
		YearOfStudy:      yearsOfStudy,
		Semester:         sem,
		StudyProgram:     studyProgram,
		InstituteFilters: institutes,
	}
	resp2 := FiltersCourse{
		Filters: resp,
	}
	writeJSON(w, http.StatusOK, resp2)
	h.logger.Info(fmt.Sprintf("Filter Response %v", resp))
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
		r.Get("/profile", h.GetFacultyFilters)
		r.Get("/course", h.GetCoursesFilters)
	})
}

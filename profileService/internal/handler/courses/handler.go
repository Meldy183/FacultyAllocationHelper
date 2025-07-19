package courses

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/CompleteCourse"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/facultyProfile"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileVersion"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/staff"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/sharedContent"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/academicYear"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/institute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/semester"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Handler struct {
	logger              *zap.Logger
	fullCourseService   CompleteCourse.Service
	staffService        staff.Service
	academicYearService academicYear.Service
	semesterService     semester.Service
	instituteService    institute.Service
	profileService      profileVersion.Service
}

func NewHandler(logger *zap.Logger, fullCourseService CompleteCourse.Service,
	academicYearService academicYear.Service, semesterService semester.Service,
	instituteService institute.Service) *Handler {
	return &Handler{
		logger:              logger,
		fullCourseService:   fullCourseService,
		academicYearService: academicYearService,
		semesterService:     semesterService,
		instituteService:    instituteService,
	}
}

func (h *Handler) GetAllCoursesByFilters(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) AddNewCourse(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetCourse(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		h.logger.Error("error getting course id",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetCourseByID),
			zap.Error(err),
		)
		writeError(w, http.StatusBadRequest, "Invalid course id")
		return
	}
	fullCourse, err := h.fullCourseService.GetFullCourseInfoByID(ctx, id)
	if err != nil {
		h.logger.Error("error getting full course",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetCourseByID),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting full course")
		return
	}
	staffs, err := h.staffService.GetAllStaffByInstanceID(ctx, id)
	piStaff := h.staffService.GetPI(staffs)
	piVersionID := piStaff.ProfileVersionID
	piProfileVersion, err := h.profileService.GetVersionByProfileID(ctx, int64(piVersionID), fullCourse.Year)
	if err != nil {
		h.logger.Error("error getting version info",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetCourseByID),
			zap.Error(err),
			)
		writeError(w, http.StatusInternalServerError, "error getting version info")
		return
	}
	piFaculty := &sharedContent.Faculty{
		ProfileVersionID: int64(piStaff.ProfileVersionID),
		NameEng:          piNameEng,
	}
	tiStaff := h.staffService.GetTI(staffs)
	tiVersionID := tiStaff.ProfileVersionID
	tasStaff := h.staffService.GetTAs(staffs)
	if err != nil {
		h.logger.Error("error getting staff",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetCourseByID),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting staff")
		return
	}
	academicYearName, err := h.academicYearService.GetAcademicYearNameByID(ctx, fullCourse.InstanceID)
	semesterName, err := h.semesterService.GetSemesterNameByID(ctx, int64(fullCourse.SemesterID))
	instituteObj, err := h.instituteService.GetInstituteByID(ctx, fullCourse.ResponsibleInstituteID)
	isAllocDone := fullCourse.GroupsNeeded-*fullCourse.GroupsTaken == 0
	pi := &sharedContent.PI{
		AllocationStatus: (*string)(fullCourse.PIAllocationStatus),
		ProfileData:,
	}
	ti := &sharedContent.PI{}
	tas := make([]sharedContent.Faculty, 0)
	course := &sharedContent.Course{
		InstanceID:           &fullCourse.InstanceID,
		BriefName:            &fullCourse.Name,
		OfficialName:         fullCourse.OfficialName,
		AcademicYearName:     academicYearName,
		SemesterName:         semesterName,
		StudyPrograms:        fullCourse.StudyPrograms,
		InstituteName:        &instituteObj.Name,
		Tracks:               fullCourse.Tracks,
		IsAllocationFinished: &isAllocDone,
		Mode:                 (*string)(fullCourse.Mode),
		StudyYear:            &fullCourse.Year,
		Form:                 (*string)(fullCourse.Form),
		LectureHours:         fullCourse.LecHours,
		LabHours:             fullCourse.LabHours,
		GroupsNeeded:         &fullCourse.GroupsNeeded,
		GroupsTaken:          fullCourse.GroupsTaken,
		PI:                   *pi,
		TI:                   *ti,
		TAs:                  tas,
	}
	resp := &GetCourseResponse{Course: *course}
	h.logger.Info("GetCourse Success",
		zap.String("layer", logctx.LogHandlerLayer),
		zap.String("function", logctx.LogGetCourseByID),
	)
	writeJSON(w, http.StatusOK, resp)
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

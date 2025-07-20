package courses

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/CompleteCourse"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/courseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/facultyProfile"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileInstitute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileVersion"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/responsibleInstitute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/staff"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/sharedContent"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/academicYear"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/semester"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Handler struct {
	logger                      *zap.Logger
	fullCourseService           CompleteCourse.Service
	staffService                staff.Service
	academicYearService         academicYear.Service
	semesterService             semester.Service
	responsibleInstituteService responsibleInstitute.Service
	profileVersionService       profileVersion.Service
	profileService              facultyProfile.Service
	profileInstituteService     profileInstitute.Service
	courseInstanceService       courseInstance.Service
}

func NewHandler(logger *zap.Logger, fullCourseService CompleteCourse.Service,
	staffService staff.Service, profileInstituteService profileInstitute.Service,
	academicYearService academicYear.Service, semesterService semester.Service,
	responsibleInstituteService responsibleInstitute.Service, profileVersionService profileVersion.Service,
	profileService facultyProfile.Service, courseInstanceService courseInstance.Service) *Handler {
	return &Handler{
		logger:                      logger,
		fullCourseService:           fullCourseService,
		academicYearService:         academicYearService,
		semesterService:             semesterService,
		responsibleInstituteService: responsibleInstituteService,
		profileVersionService:       profileVersionService,
		profileService:              profileService,
		profileInstituteService:     profileInstituteService,
		staffService:                staffService,
		courseInstanceService:       courseInstanceService,
	}
}

func (h *Handler) GetAllCoursesByFilters(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	isAllocationFinished := r.URL.Query().Get("allocation_finished") == "true"
	year, err := strconv.ParseInt(r.URL.Query().Get("year"), 10, 64)
	if err != nil {
		h.logger.Error("Error parsing year",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetAllCourses),
			zap.Error(err),
		)
		writeError(w, http.StatusBadRequest, "Error parsing year")
		return
	}
	yearsOfStudyIDs, err := convertStrToInt(r.URL.Query()["academic_year"])
	if err != nil {
		h.logger.Error("Error parsing year_Studies",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetAllCourses),
			zap.Error(err),
		)
		writeError(w, http.StatusBadRequest, "Error parsing year_Studies")
		return
	}
	semesterIDs, err := convertStrToInt(r.URL.Query()["semester_ids"])
	if err != nil {
		h.logger.Error("Error parsing semester_id",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetAllCourses),
			zap.Error(err),
		)
		writeError(w, http.StatusBadRequest, "Error parsing semester_id")
		return
	}
	programIDs, err := convertStrToInt(r.URL.Query()["study_program_ids"])
	if err != nil {
		h.logger.Error("Error parsing study_program_ids",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetAllCourses),
			zap.Error(err),
		)
		writeError(w, http.StatusBadRequest, "Error parsing study_program_ids")
		return
	}
	responsibleInstituteIDs, err := convertStrToInt(r.URL.Query()["responsible_institute_ids"])
	if err != nil {
		h.logger.Error("Error parsing responsible_institute_ids",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetAllCourses),
			zap.Error(err),
		)
		writeError(w, http.StatusBadRequest, "Error parsing responsible_institute_ids")
		return
	}
	profile_version_id, err := strconv.ParseInt(r.URL.Query().Get("profile_version_id"), 10, 64)
	if err != nil {
		h.logger.Error("Error parsing profile_version_id",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetAllCourses),
			zap.Error(err),
		)
		writeError(w, http.StatusBadRequest, "Error parsing profile_version_id")
		return
	}
	isAllocationFinishedIDs, err := h.courseInstanceService.
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
	course, interrupted := h.CombineCourseCard(w, err, ctx, id)
	if interrupted {
		return
	}
	resp := &GetCourseResponse{Course: *course}
	h.logger.Info("GetCourse Success",
		zap.String("layer", logctx.LogHandlerLayer),
		zap.String("function", logctx.LogGetCourseByID),
	)
	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) CombineCourseCard(w http.ResponseWriter, err error, ctx context.Context, id int64) (*sharedContent.Course, bool) {
	fullCourse, err := h.fullCourseService.GetFullCourseInfoByID(ctx, id)
	if err != nil {
		h.logger.Error("error getting full course",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetCourseByID),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting full course")
		return nil, true
	}
	staffs, err := h.staffService.GetAllStaffByInstanceID(ctx, id)
	piStaff := h.staffService.GetPI(staffs)
	if err != nil {
		h.logger.Error("error getting version info",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetCourseByID),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting version info")
		return nil, true
	}
	piFaculty, err := h.staffToFaculty(ctx, piStaff)
	if err != nil {
		h.logger.Error("error getting faculty",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetCourseByID),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting faculty")
		return nil, true
	}
	tiStaff := h.staffService.GetTI(staffs)
	tiFaculty, err := h.staffToFaculty(ctx, tiStaff)
	if err != nil {
		h.logger.Error("error getting faculty",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetCourseByID),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting faculty")
		return nil, true
	}
	tasStaff := h.staffService.GetTAs(staffs)
	tas := make([]sharedContent.Faculty, 0)
	for _, elem := range tasStaff {
		facObj, err := h.staffToFaculty(ctx, elem)
		if err != nil {
			h.logger.Error("error getting faculty",
				zap.String("layer", logctx.LogHandlerLayer),
				zap.String("function", logctx.LogGetCourseByID),
				zap.Error(err),
			)
			writeError(w, http.StatusInternalServerError, "error getting faculty")
			return nil, true
		}
		tas = append(tas, *facObj)
	}
	academicYearName, err := h.academicYearService.GetAcademicYearNameByID(ctx, fullCourse.InstanceID)
	semesterName, err := h.semesterService.GetSemesterNameByID(ctx, int64(fullCourse.SemesterID))
	instituteObj, err := h.responsibleInstituteService.GetResponsibleInstituteNameByID(ctx, fullCourse.ResponsibleInstituteID)
	isAllocDone := fullCourse.GroupsNeeded-*fullCourse.GroupsTaken == 0
	pi := &sharedContent.PI{
		AllocationStatus: (*string)(fullCourse.PIAllocationStatus),
		ProfileData:      piFaculty,
	}
	ti := &sharedContent.PI{
		AllocationStatus: (*string)(fullCourse.PIAllocationStatus),
		ProfileData:      tiFaculty,
	}
	course := &sharedContent.Course{
		InstanceID:           &fullCourse.InstanceID,
		BriefName:            &fullCourse.Name,
		OfficialName:         fullCourse.OfficialName,
		AcademicYearName:     academicYearName,
		SemesterName:         semesterName,
		StudyPrograms:        fullCourse.StudyPrograms,
		InstituteName:        instituteObj,
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
	return course, false
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

func (h *Handler) staffToFaculty(ctx context.Context, s *staff.Staff) (*sharedContent.Faculty, error) {
	profileObj := h.getProfileByVersionID(ctx, s.ProfileVersionID)
	institutes, err := h.profileInstituteService.GetUserInstitutesByProfileID(ctx, profileObj.ProfileID)
	instNames := make([]string, 0)
	for _, elem := range institutes {
		instNames = append(instNames, elem.Name)
	}
	if err != nil {
		h.logger.Error("GetInstitutesByProfileID",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogStaffToFaculty),
			zap.Error(err),
		)
		return nil, fmt.Errorf("error getting institutes by id: %v", err)
	}
	fuck := &sharedContent.Faculty{
		ProfileVersionID: s.ProfileVersionID,
		NameEng:          &profileObj.EnglishName,
		Alias:            &profileObj.Alias,
		Email:            &profileObj.Email,
		PositionName:     s.PositionType,
		InstituteNames:   instNames,
		Classes:          nil, // TODO: implement me
		IsConfirmed:      s.IsConfirmed,
	}
	return fuck, nil
}

func (h *Handler) getProfileByVersionID(ctx context.Context, versionID int64) *facultyProfile.UserProfile {
	version, err := h.profileVersionService.GetVersionByVersionID(ctx, versionID)
	if err != nil {
		h.logger.Error("getProfileByVersionID",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetProfileByVersionID),
			zap.Error(err),
		)
		return nil
	}
	profileObj, err := h.profileService.GetProfileByID(ctx, version.ProfileID)
	if err != nil {
		h.logger.Error("getProfileByVersionID",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetProfileByVersionID),
			zap.Error(err),
		)
		return nil
	}
	return profileObj
}

func convertStrToInt(s []string) ([]int, error) {
	ints := make([]int, 0)
	for _, v := range s {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}

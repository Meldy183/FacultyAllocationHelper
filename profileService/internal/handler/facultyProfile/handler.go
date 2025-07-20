package facultyProfile

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/facultyProfile"
	institute2 "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/institute"
	position2 "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/position"
	profileCourseInstance2 "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileCourseInstance"
	userinstituteDomain "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileInstitute"
	profileLanguage2 "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileLanguage"
	profileVersionDomain "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileVersion"
	workloadDomain "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/workload"
	handlerWorkload "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/workload"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/institute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/workload"
	"go.uber.org/zap"
)

type Handler struct {
	serviceUP             facultyProfile.Service
	serviceUI             userinstituteDomain.Service
	serviceLang           profileLanguage2.Service
	serviceCourse         profileCourseInstance2.Service
	servicePosition       position2.Service
	serviceInstitute      institute2.Service
	serviceVersionProfile profileVersionDomain.Service
	serviceWorkload       workload.Service
	logger                *zap.Logger
}

func NewHandler(
	serviceUP facultyProfile.Service,
	serviceUI userinstituteDomain.Service,
	serviceLang profileLanguage2.Service,
	serviceCourse profileCourseInstance2.Service,
	servicePosition position2.Service,
	serviceInstitute institute2.Service,
	serviceVersionProfile profileVersionDomain.Service,
	workloadService workload.Service,
	logger *zap.Logger,
) *Handler {
	return &Handler{
		serviceUP:             serviceUP,
		serviceUI:             serviceUI,
		serviceLang:           serviceLang,
		serviceCourse:         serviceCourse,
		servicePosition:       servicePosition,
		serviceInstitute:      serviceInstitute,
		serviceVersionProfile: serviceVersionProfile,
		serviceWorkload:       workloadService,
		logger:                logger,
	}
}
func (h *Handler) AddProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req AddProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("error decoding json",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogAddProfile),
			zap.Error(err))
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if req.NameEnglish == "" {
		h.logger.Error("invalid name_english",
			zap.String("engName", req.NameEnglish),
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogAddProfile),
		)
		writeError(w, http.StatusBadRequest, "invalid NameEnglish")
		return
	}
	if req.PositionID <= 0 {
		h.logger.Error("invalid position",
			zap.Int("position_id", req.PositionID),
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogAddProfile),
		)
		writeError(w, http.StatusBadRequest, "invalid position")
		return
	}
	instituteIDs := req.InstituteIDs
	if len(instituteIDs) <= 0 {
		h.logger.Error("invalid institute",
			zap.Int("institute_id", req.PositionID),
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogAddProfile),
		)
		writeError(w, http.StatusBadRequest, "invalid institute")
		return
	}
	for _, elem := range req.InstituteIDs {
		if elem <= 0 {
			h.logger.Error("invalid instituteID",
				zap.Int("InstituteID", elem),
				zap.String("layer", logctx.LogHandlerLayer),
				zap.String("function", logctx.LogAddProfile),
			)
			writeError(w, http.StatusBadRequest, "invalid institute LabID")
			return
		}
	}

	if req.Year < 2015 {
		h.logger.Error("invalid year",
			zap.Int("Year", req.Year),
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogAddProfile),
		)
		writeError(w, http.StatusBadRequest, "invalid year")
		return
	}

	profile := &facultyProfile.UserProfile{
		EnglishName: req.NameEnglish,
		Email:       req.Email,
		Alias:       req.Alias,
	}
	if err := h.serviceUP.AddProfile(ctx, profile); err != nil {
		h.logger.Error("error creating facultyProfile",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogAddProfile),
			zap.Error(err))
		writeError(w, http.StatusInternalServerError, "error creating facultyProfile")
		return
	}
	institutesList := make([]string, 0)
	for _, elem := range req.InstituteIDs {
		userInstitute := &userinstituteDomain.UserInstitute{
			ProfileID:   profile.ProfileID,
			InstituteID: elem,
		}
		if err := h.serviceUI.AddUserInstitute(ctx, userInstitute); err != nil {
			h.logger.Error("error adding profileInstitute",
				zap.String("layer", logctx.LogHandlerLayer),
				zap.String("function", logctx.LogAddProfile),
				zap.Error(err))
			writeError(w, http.StatusInternalServerError, "error adding profileInstitute")
			return
		}
		inst, err := h.serviceInstitute.GetInstituteByID(ctx, int64(elem))
		if err != nil {
			h.logger.Error("error getting institute",
				zap.String("layer", logctx.LogHandlerLayer),
				zap.String("function", logctx.LogAddProfile),
				zap.Error(err),
			)
			writeError(w, http.StatusInternalServerError, "error getting institute")
			return
		}
		institutesList = append(institutesList, inst.Name)
	}
	version := &profileVersionDomain.ProfileVersion{
		ProfileID:  profile.ProfileID,
		Year:       req.Year,
		PositionID: req.PositionID,
	}
	err := h.serviceVersionProfile.AddProfileVersion(ctx, version)
	if err != nil {
		h.logger.Error("error adding profileVersion",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogAddProfile),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error adding profileVersion")
		return
	}
	positionName, err := h.servicePosition.GetPositionByID(ctx, version.PositionID)
	if err != nil {
		h.logger.Error("error getting position",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogAddProfile),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting position")
		return
	}
	if h.AddWorkloadAddingProfileVersion(w, version, err, ctx) {
		return
	}
	resp := &AddProfileResponse{
		ProfileVersionID: version.ProfileVersionId,
		NameEnglish:      req.NameEnglish,
		Year:             version.Year,
		PositionName:     *positionName,
		Email:            profile.Email,
		Alias:            profile.Alias,
		InstituteNames:   institutesList,
	}
	h.logger.Info("success adding profile",
		zap.String("layer", logctx.LogHandlerLayer),
		zap.String("function", logctx.LogAddProfile),
	)
	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) AddWorkloadAddingProfileVersion(w http.ResponseWriter, version *profileVersionDomain.ProfileVersion, err error, ctx context.Context) bool {
	workloadStats := workloadDomain.Workload{
		ProfileVersionID: version.ProfileVersionId,
		SemesterID:       1,
		LecturesCount:    0,
		TutorialsCount:   0,
		LabsCount:        0,
		ElectivesCount:   0,
		Rate:             0,
	}
	err = h.serviceWorkload.AddSemesterWorkload(ctx, &workloadStats)
	if err != nil {
		h.logger.Error("error adding workloadStats",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogAddProfile),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error adding workloadStats")
		return true
	}
	workloadStats.SemesterID = 2
	err = h.serviceWorkload.AddSemesterWorkload(ctx, &workloadStats)
	if err != nil {
		h.logger.Error("error adding workloadStats",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogAddProfile),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error adding workloadStats")
		return true
	}
	workloadStats.SemesterID = 3
	err = h.serviceWorkload.AddSemesterWorkload(ctx, &workloadStats)
	if err != nil {
		h.logger.Error("error adding workloadStats",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogAddProfile),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error adding workloadStats")
		return true
	}
	return false
}

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idParam := chi.URLParam(r, "id")
	versionID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		h.logger.Error("error parsing id param",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogAddProfile),
			zap.String("id", idParam),
			zap.Error(err),
		)
	}
	version, err := h.serviceVersionProfile.GetVersionByVersionID(ctx, versionID)
	if err != nil {
		h.logger.Error("error getting version",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogAddProfile),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting version")
		return
	}
	if err != nil || versionID <= 0 {
		h.logger.Error("invalid profileID",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetProfileByID),
			zap.Int64("id", versionID),
		)
		writeError(w, http.StatusBadRequest, "invalid profileID")
		return
	}
	profile, err := h.serviceUP.GetProfileByID(ctx, version.ProfileID)
	if err != nil {
		h.logger.Error("error getting facultyProfile",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetProfileByID),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting facultyProfile")
		return
	}
	instObjects, err := h.serviceUI.GetUserInstitutesByProfileID(ctx, version.ProfileID)
	institutes := institute.ConvertInstitutesToString(instObjects)
	if err != nil {
		h.logger.Error("error getting institute",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetProfileByID),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting facultyProfile")
		return
	}
	languages, err := h.serviceLang.GetUserLanguages(ctx, version.ProfileID)
	if err != nil {
		h.logger.Error("error getting languages",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetProfileByID),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting languages")
		return
	}
	var langEntries []Lang
	for _, l := range languages {
		langEntries = append(langEntries, Lang{
			Language: l.LanguageName,
		})
	}

	positionName, err := h.servicePosition.GetPositionByID(ctx, version.PositionID)
	if err != nil {
		h.logger.Error("error getting position name by id",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetProfileByID),
			zap.Int("id", version.PositionID),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting position by id")
		return
	}
	workloadHandler := handlerWorkload.NewWorkloadHandler(&h.serviceWorkload, h.logger)

	sem1, sem2, sem3, notDone := workloadHandler.GetYearWorkload(w, err, ctx, versionID)
	if notDone {
		return
	}
	stats := workloadHandler.WorkloadToClasses(sem1, sem2, sem3)
	resp := GetProfileResponse{
		ProfileVersionID: version.ProfileVersionId,
		Year:             version.Year,
		NameEnglish:      profile.EnglishName,
		NameRussian:      profile.RussianName,
		Alias:            profile.Alias,
		Email:            profile.Email,
		PositionName:     *positionName,
		InstituteNames:   *institutes,
		StudentType:      version.StudentType,
		Degree:           version.Degree,
		Fsro:             version.Fsro,
		LanguageCodes:    &langEntries,
		EmploymentType:   version.EmploymentType,
		HiringStatus:     profile.Status,
		Mode:             version.Mode,
		MaxLoad:          version.MaxLoad,
		FrontalHours:     version.FrontalHours,
		ExtraActivity:    version.ExtraActivities,
		WorkloadStats:    stats,
	}
	h.logger.Info("Successfully fetched facultyProfile",
		zap.String("layer", logctx.LogHandlerLayer),
		zap.String("function", logctx.LogGetProfileByID),
	)
	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) GetAllFaculties(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	instQuery := r.URL.Query()["institute_ids"]
	year, err := strconv.Atoi(r.URL.Query().Get("year"))
	if err != nil {
		h.logger.Error("error parsing year",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogAddProfile),
			zap.Error(err),
		)
		writeError(w, http.StatusBadRequest, "error parsing year")
		return
	}
	var institutes []int
	for _, elem := range instQuery {
		id, err := strconv.Atoi(elem)
		if err != nil {
			h.logger.Error(
				"Error converting query to int",
				zap.String("layer", logctx.LogHandlerLayer),
				zap.String("function", logctx.LogGetAllFaculties),
				zap.Error(err),
			)
			writeError(w, http.StatusInternalServerError, "error institute id")
			return
		}
		institutes = append(institutes, id)
	}
	posQuery := r.URL.Query()["positions"]
	var positions []int
	for _, elem := range posQuery {
		pos, err := strconv.Atoi(elem)
		if err != nil {
			h.logger.Error("error converting to int the position",
				zap.String("layer", logctx.LogHandlerLayer),
				zap.String("function", logctx.LogGetAllFaculties),
				zap.Error(err),
			)
			writeError(w, http.StatusInternalServerError, "error position id")
			return
		}
		positions = append(positions, pos)
	}
	h.logger.Warn("success getting all faculties",
		zap.String("layer", logctx.LogHandlerLayer),
		zap.String("function", logctx.LogGetAllFaculties),
		zap.Any("institutes", institutes),
		zap.Any("positions", positions),
	)
	profileIds, err := h.serviceUP.GetProfilesByFilters(ctx, institutes, positions)
	if err != nil {
		h.logger.Error("Error getting facultyProfile ids",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetAllFaculties),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting profiles")
		return
	}
	h.logger.Debug("Successfully fetched facultyProfiles",
		zap.String("layer", logctx.LogHandlerLayer),
		zap.String("function", logctx.LogGetAllFaculties),
		zap.Any("ids", profileIds),
	)
	resp := make([]ShortProfile, 0)
	for _, id := range profileIds {
		profile, err := h.serviceUP.GetProfileByID(ctx, id)
		if err != nil {
			h.logger.Error("Error getting facultyProfile by id",
				zap.String("layer", logctx.LogHandlerLayer),
				zap.String("function", logctx.LogGetAllFaculties),
				zap.Int64("LabID", id),
				zap.Error(err),
			)
			writeError(w, http.StatusInternalServerError, "error getting facultyProfile")
			return
		}
		version, err := h.serviceVersionProfile.GetVersionByProfileID(ctx, id, year)
		if err != nil {
			h.logger.Error("Error getting facultyVersionProfile by id",
				zap.String("layer", logctx.LogHandlerLayer),
				zap.String("function", logctx.LogGetAllFaculties),
				zap.Int64("LabID", id),
				zap.Error(err),
			)
			writeError(w, http.StatusInternalServerError, "error getting facultyVersionProfile")
			return
		}
		h.logger.Warn("success getting facultyProfile by id",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetAllFaculties),
			zap.Any("version", version),
		)
		pos, err := h.servicePosition.GetPositionByID(ctx, version.PositionID)
		if err != nil {
			h.logger.Error("Error getting position by id",
				zap.String("layer", logctx.LogHandlerLayer),
				zap.String("function", logctx.LogGetAllFaculties),
				zap.Int64("LabID", id),
				zap.Error(err),
			)
			writeError(w, http.StatusInternalServerError, "error getting position by id")
			return
		}
		institutesStruct, err := h.serviceUI.GetUserInstitutesByProfileID(ctx, profile.ProfileID)
		if err != nil {
			h.logger.Error("Error getting user institute by id",
				zap.String("layer", logctx.LogHandlerLayer),
				zap.String("function", logctx.LogGetAllFaculties),
				zap.Int64("LabID", id),
				zap.Error(err),
			)
			writeError(w, http.StatusInternalServerError, "error getting user institute by id")
			return

		}
		instNames := make([]string, 0)
		for _, inst := range institutesStruct {
			instNames = append(instNames, inst.Name)
		}
		resp = append(resp, ShortProfile{
			ProfileVersionID: version.ProfileVersionId,
			NameEnglish:      profile.EnglishName,
			Alias:            profile.Alias,
			Email:            profile.Email,
			Position:         *pos,
			Institutes:       instNames,
		})
	}
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
		r.Post("/addProfile", h.AddProfile)
		r.Get("/getProfile/{id}", h.GetProfile)
		r.Get("/getAllProfiles", h.GetAllFaculties)
	})
}

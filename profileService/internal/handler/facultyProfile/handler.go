package facultyProfile

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	userprofileDomain "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/facultyProfile"
	userinstituteDomain "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileInstitute"
	profileVersionDomain "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileVersion"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/facultyProfile"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/institute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/position"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/profileCourseInstance"
	userinstitute "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/profileInstitute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/profileLanguage"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/profileVersion"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Handler struct {
	serviceUP             *facultyProfile.Service
	serviceUI             *userinstitute.Service
	serviceLang           *profileLanguage.Service
	serviceCourse         *profileCourseInstance.Service
	servicePosition       *position.Service
	serviceInstitute      *institute.Service
	serviceVersionProfile *profileVersion.Service
	logger                *zap.Logger
}

func NewHandler(serviceUP *facultyProfile.Service,
	serviceUI *userinstitute.Service,
	serviceLang *profileLanguage.Service,
	serviceCourse *profileCourseInstance.Service,
	servicePosition *position.Service,
	serviceInstitute *institute.Service,
	logger *zap.Logger,
) *Handler {
	return &Handler{
		serviceUP:        serviceUP,
		serviceUI:        serviceUI,
		serviceLang:      serviceLang,
		serviceCourse:    serviceCourse,
		servicePosition:  servicePosition,
		serviceInstitute: serviceInstitute,
		logger:           logger.Named("userprofile_handler"),
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
	for _, elem := range req.InstituteID {
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

	profile := &userprofileDomain.UserProfile{
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
	for _, elem := range req.InstituteID {
		userInstitute := &userinstituteDomain.UserInstitute{
			ProfileID:        profile.ProfileID,
			InstituteID:      elem,
			IsRepresentative: req.IsRepresentative,
		}
		if err := h.serviceUI.AddUserInstitute(ctx, userInstitute); err != nil {
			h.logger.Error("error adding profileInstitute",
				zap.String("layer", logctx.LogHandlerLayer),
				zap.String("function", logctx.LogAddProfile),
				zap.Error(err))
			writeError(w, http.StatusInternalServerError, "error adding profileInstitute")
			return
		}
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
	}
	PositionName, err := h.servicePosition.GetPositionByID(ctx, version.PositionID)
	if err != nil {
		h.logger.Error("error getting position",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogAddProfile),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting position")
	}
	profileSub := &Profile{
		NameEnglish: req.NameEnglish,
		PositionID:  PositionName,
	}
	resp := &AddProfileResponse{
		ProfileVersionID: version.ProfileVersionId,
	}
}

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idParam := chi.URLParam(r, "id")
	_profileID, err := strconv.ParseUint(idParam, 10, 64)
	profileID := int64(_profileID)
	if err != nil || profileID <= 0 {
		h.logger.Error("invalid profileID",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetProfileByID),
			zap.Int64("id", profileID),
		)
		writeError(w, http.StatusBadRequest, "invalid profileID")
		return
	}
	profile, err := h.serviceUP.GetProfileByID(ctx, profileID)
	if err != nil {
		h.logger.Error("error getting facultyProfile",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetProfileByID),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting facultyProfile")
		return
	}
	inst, err := h.serviceUI.GetUserInstituteByID(ctx, profileID)
	if err != nil {
		h.logger.Error("error getting institute",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetProfileByID),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting facultyProfile")
		return
	}
	languages, err := h.serviceLang.GetUserLanguages(ctx, profileID)
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
	coursesID, err := h.serviceCourse.GetInstancesByProfileID(ctx, profileID)
	if err != nil {
		h.logger.Error("error getting course ids",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetProfileByID),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting courses")
		return
	}
	courseEntries := make([]Course, 0)
	for _, courseID := range coursesID {
		h.logger.Info("getting course entry",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetProfileByID),
			zap.Int64("id", courseID),
		)
		courseEntries = append(courseEntries, Course{
			CourseInstanceID: courseID,
		})
	}
	positionName, err := h.servicePosition.GetPositionByID(ctx, profile.PositionID)
	if err != nil {
		h.logger.Error("error getting position name by id",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetProfileByID),
			zap.Int("id", profile.PositionID),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting position by id")
		return
	}
	resp := GetProfileResponse{
		ProfileID:      profileID,
		NameEnglish:    profile.EnglishName,
		NameRussian:    profile.RussianName,
		Alias:          profile.Alias,
		Email:          profile.Email,
		Position:       *positionName,
		Institute:      inst.Name,
		StudentType:    profile.StudentType,
		Degree:         profile.Degree,
		Languages:      &langEntries,
		Courses:        &courseEntries,
		EmploymentType: profile.EmploymentType,
		Mode:           (*string)(profile.Mode),
		MaxLoad:        profile.MaxLoad,
	}
	h.logger.Info("Successfully fetched facultyProfile",
		zap.String("layer", logctx.LogHandlerLayer),
		zap.String("function", logctx.LogGetProfileByID),
	)
	writeJSON(w, http.StatusOK, resp)
}

func (h *Handler) GetAllFaculties(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	instQuery := r.URL.Query()["institute_id"]
	var insts []int
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
		insts = append(insts, id)
	}
	posQuery := r.URL.Query()["position"]
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
	profileIds, err := h.serviceUP.GetProfilesByFilters(ctx, insts, positions)
	if err != nil {
		h.logger.Error("Error getting facultyProfile ids",
			zap.String("layer", logctx.LogHandlerLayer),
			zap.String("function", logctx.LogGetAllFaculties),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting profiles")
		return
	}
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
			return
		}
		version, err := h.serviceVersionProfile.GetVersionByProfileID(ctx, id)
		if err != nil {
			h.logger.Error("Error getting facultyVersionProfile by id",
				zap.String("layer", logctx.LogHandlerLayer),
				zap.String("function", logctx.LogGetAllFaculties),
				zap.Int64("LabID", id),
				zap.Error(err),
			)
			return
		}
		pos, err := h.servicePosition.GetPositionByID(ctx, version.PositionID)
		if err != nil {
			h.logger.Error("Error getting position by id",
				zap.String("layer", logctx.LogHandlerLayer),
				zap.String("function", logctx.LogGetAllFaculties),
				zap.Int64("LabID", id),
				zap.Error(err),
			)
		}
		inst, err := h.serviceUI.GetUserInstituteByID(ctx, profile.ProfileID)
		if err != nil {
			h.logger.Error("Error getting user institute by id",
				zap.String("layer", logctx.LogHandlerLayer),
				zap.String("function", logctx.LogGetAllFaculties),
				zap.Int64("LabID", id),
				zap.Error(err),
			)
		}
		resp = append(resp, ShortProfile{
			ProfileID:   profile.ProfileID,
			NameEnglish: profile.EnglishName,
			Alias:       profile.Alias,
			Email:       profile.Email,
			Position:    *pos,
			Institut:    inst.Name,
		})
	}
	writeJSON(w, http.StatusOK, resp)
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
	instituteObjects := make([]InstituteObj, 0)
	for _, elem := range institutes {
		iobj := InstituteObj{
			ID:   elem.InstituteID,
			Name: elem.Name,
		}
		instituteObjects = append(instituteObjects, iobj)
	}

	positionObjects := make([]PositionObj, 0)
	for _, elem := range positions {
		pobj := PositionObj{
			ID:   elem.PositionID,
			Name: elem.Name,
		}
		positionObjects = append(positionObjects, pobj)
	}
	responce := GetFacultyFiltersResponse{
		InstituteFilters: instituteObjects,
		PositionFilters:  positionObjects,
	}

	writeJSON(w, http.StatusOK, responce)
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
		r.Get("/filters", h.GetFacultyFilters)
	})
}

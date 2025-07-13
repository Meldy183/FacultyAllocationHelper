package userprofile

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	userinstituteDomain "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/userinstitute"
	userprofileDomain "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/userprofile"
	_ "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/institute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/position"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/usercourseinstance"
	userinstitute "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/userinstitute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/userlanguage"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/userprofile"
	"go.uber.org/zap"
)

type Handler struct {
	serviceUP       *userprofile.Service
	serviceUI       *userinstitute.Service
	serviceLang     *userlanguage.Service
	serviceCourse   *usercourseinstance.Service
	servicePosition *position.Service
	logger          *zap.Logger
}

const (
	logLayer           = "Handler"
	logAddProfile      = "AddProfile"
	logGetProfile      = "GetProfile"
	logGetAllFaculties = "GetAllFaculties"
)

func NewHandler(serviceUP *userprofile.Service,
	serviceUI *userinstitute.Service,
	serviceLang *userlanguage.Service,
	serviceCourse *usercourseinstance.Service,
	servicePosition *position.Service,
	logger *zap.Logger,
) *Handler {
	return &Handler{
		serviceUP:       serviceUP,
		serviceUI:       serviceUI,
		serviceLang:     serviceLang,
		serviceCourse:   serviceCourse,
		servicePosition: servicePosition,
		logger:          logger.Named("userprofile_handler"),
	}
}
func (h *Handler) AddProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req AddProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("error decoding json",
			zap.String("layer", logLayer),
			zap.String("function", logAddProfile),
			zap.Error(err))
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.NameEnglish == "" {
		h.logger.Error("invalid name_english",
			zap.String("engName", req.NameEnglish),
			zap.String("function", logAddProfile),
			zap.String("layer", logLayer),
		)
		writeError(w, http.StatusBadRequest, "invalid NameEnglish")
		return
	}
	if req.PositionID <= 0 {
		h.logger.Error("invalid position",
			zap.Int("position_id", req.PositionID),
			zap.String("function", logAddProfile),
			zap.String("layer", logLayer),
		)
		writeError(w, http.StatusBadRequest, "invalid position")
		return
	}
	if req.InstituteID <= 0 {
		h.logger.Error("invalid institute ID",
			zap.Int("ID", req.InstituteID),
			zap.String("function", logAddProfile),
			zap.String("layer", logLayer),
		)
		writeError(w, http.StatusBadRequest, "invalid institute ID")
		return
	}
	profile := &userprofileDomain.UserProfile{
		EnglishName: req.NameEnglish,
		Email:       req.Email,
		PositionID:  req.PositionID,
		Alias:       req.Alias,
	}
	if err := h.serviceUP.Create(ctx, profile); err != nil {
		h.logger.Error("error creating userprofile",
			zap.String("layer", logLayer),
			zap.String("function", logAddProfile),
			zap.Error(err))
		writeError(w, http.StatusInternalServerError, "error creating userprofile")
		return
	}
	userInstitute := &userinstituteDomain.UserInstitute{
		ProfileID:        profile.ProfileID,
		InstituteID:      req.InstituteID,
		IsRepresentative: req.IsRepresentative,
	}
	if err := h.serviceUI.AddUserInstitute(ctx, userInstitute); err != nil {
		h.logger.Error("error adding userinstitute",
			zap.String("layer", logLayer),
			zap.String("function", logAddProfile),
			zap.Error(err))
		writeError(w, http.StatusInternalServerError, "error adding userinstitute")
		return
	}

	h.logger.Info("Successfully added userinstitute",
		zap.String("layer", logLayer),
		zap.String("function", logAddProfile),
	)
	writeJSON(w, http.StatusCreated, profile)
}

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idParam := chi.URLParam(r, "id")
	_profileID, err := strconv.ParseUint(idParam, 10, 64)
	profileID := int64(_profileID)
	if err != nil || profileID <= 0 {
		h.logger.Error("invalid profileID",
			zap.String("function", logGetProfile),
			zap.String("layer", logLayer),
			zap.Int64("id", profileID),
		)
		writeError(w, http.StatusBadRequest, "invalid profileID")
		return
	}
	profile, err := h.serviceUP.GetByProfileID(ctx, profileID)
	if err != nil {
		h.logger.Error("error getting userprofile",
			zap.String("layer", logLayer),
			zap.String("function", logGetProfile),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting userprofile")
		return
	}
	inst, err := h.serviceUI.GetUserInstituteByID(ctx, profileID)
	if err != nil {
		h.logger.Error("error getting institute",
			zap.String("layer", logLayer),
			zap.String("function", logGetProfile),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting userprofile")
		return
	}
	languages, err := h.serviceLang.GetUserLanguages(ctx, profileID)
	if err != nil {
		h.logger.Error("error getting languages",
			zap.String("layer", logLayer),
			zap.String("function", logGetProfile),
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
			zap.String("layer", logLayer),
			zap.String("function", logGetProfile),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting courses")
		return
	}
	courseEntries := make([]Course, 0)
	for _, courseID := range coursesID {
		h.logger.Info("getting course entry",
			zap.String("layer", logLayer),
			zap.String("function", logGetProfile),
			zap.Int64("id", courseID),
		)
		courseEntries = append(courseEntries, Course{
			CourseInstanceID: courseID,
		})
	}
	positionName, err := h.servicePosition.GetByID(ctx, profile.PositionID)
	if err != nil {
		h.logger.Error("error getting position name by id",
			zap.String("layer", logLayer),
			zap.String("function", logGetProfile),
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
	h.logger.Info("Successfully fetched profile",
		zap.String("layer", logLayer),
		zap.String("function", logGetProfile),
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
				zap.String("layer", logLayer),
				zap.String("function", logGetAllFaculties),
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
				zap.String("layer", logLayer),
				zap.String("function", logGetAllFaculties),
				zap.Error(err),
			)
			writeError(w, http.StatusInternalServerError, "error position id")
			return
		}
		positions = append(positions, pos)
	}
	profileIds, err := h.serviceUP.GetProfilesByFilter(ctx, insts, positions)
	if err != nil {
		h.logger.Error("Error getting profile ids",
			zap.String("layer", logLayer),
			zap.String("function", logGetAllFaculties),
			zap.Error(err),
		)
		writeError(w, http.StatusInternalServerError, "error getting profiles")
		return
	}
	resp := &GetAllFacultiesResponse{
		Profiles: profileIds,
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
		r.Post("/addUser", h.AddProfile)
		r.Get("/getUser/{id}", h.GetProfile)
		r.Get("/getAllUsers", h.GetAllFaculties)
	})
}

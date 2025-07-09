package userprofile

import (
	"encoding/json"
	userinstituteDomain "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/userinstitute"
	"strings"

	userprofileDomain "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/userprofile"
	_ "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/institute"
	userinstitute "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/userinstitute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/userprofile"
	"go.uber.org/zap"
	"net/http"
)

type Handler struct {
	serviceUP *userprofile.Service
	serviceUI *userinstitute.Service
	logger    *zap.Logger
}

const (
	logLayer      = "Handler"
	lofAddProfile = "AddProfile"
)

func NewHandler(serviceUP *userprofile.Service, serviceUI *userinstitute.Service, logger *zap.Logger) *Handler {
	return &Handler{
		serviceUP: serviceUP,
		logger:    logger.Named("userprofile_handler"),
		serviceUI: serviceUI,
	}
}
func (h *Handler) AddProfile(w http.ResponseWriter, r *http.Request) {
	var req AddProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("error decoding json",
			zap.String("layer", logLayer),
			zap.String("function", lofAddProfile),
			zap.Error(err))
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if !strings.Contains(req.Alias, "@") || req.Alias == "" {
		h.logger.Error("invalid alias",
			zap.String("alias", req.Alias),
			zap.String("function", lofAddProfile),
			zap.String("layer", logLayer),
		)
		writeError(w, http.StatusBadRequest, "invalid alias")
		return
	}
	if req.NameEnglish == "" {
		h.logger.Error("invalid name_english",
			zap.String("engName", req.NameEnglish),
			zap.String("function", lofAddProfile),
			zap.String("layer", logLayer),
		)
		writeError(w, http.StatusBadRequest, "invalid NameEnglish")
		return
	}
	if req.Position == "" {
		h.logger.Error("invalid position",
			zap.String("position", req.Position),
			zap.String("function", lofAddProfile),
			zap.String("layer", logLayer),
		)
		writeError(w, http.StatusBadRequest, "invalid position")
		return
	}
	if req.InstituteID <= 0 {
		h.logger.Error("invalid institute ID",
			zap.Int("ID", req.InstituteID),
			zap.String("function", lofAddProfile),
			zap.String("layer", logLayer),
		)
		writeError(w, http.StatusBadRequest, "invalid institute ID")
		return
	}
	profile := &userprofileDomain.UserProfile{
		EnglishName: req.NameEnglish,
		Email:       req.Email,
		Position:    userprofileDomain.Position(req.Position),
		Alias:       req.Alias,
	}
	if err := h.serviceUP.Create(r.Context(), profile); err != nil {
		h.logger.Error("error creating userprofile",
			zap.String("layer", logLayer),
			zap.String("function", lofAddProfile),
			zap.Error(err))
		writeError(w, http.StatusInternalServerError, "error creating userprofile")
		return
	}
	userInstitute := &userinstituteDomain.UserInstitute{
		ProfileID:        profile.ProfileID,
		InstituteID:      req.InstituteID,
		IsRepresentative: req.IsRepresentative,
	}
	if err := h.serviceUI.AddUserInstitute(r.Context(), userInstitute); err != nil {
		h.logger.Error("error adding userinstitute",
			zap.String("layer", logLayer),
			zap.String("function", lofAddProfile),
			zap.Error(err))
		writeError(w, http.StatusInternalServerError, "error adding userinstitute")
		return
	}

}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

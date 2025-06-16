package handlers

import (
	"encoding/json"
	"net/http"

	app "gitlab.pg.innopolis.university/f.markin/fah/auth/internal/application/service"
)

type Handlers struct {
	AuthService *app.AuthService
}

func NewHandlers(authService *app.AuthService) *Handlers {
	return &Handlers{
		AuthService: authService,
	}
}
func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	logresp, err := h.AuthService.Login(r.Context(), loginData.Email, loginData.Password)
	if err != nil {
		http.Error(w, "Failed to login", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(logresp)
	if err != nil {
		http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResp)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
func (h *Handlers) Logout(w http.ResponseWriter, r *http.Request) {
	err := h.AuthService.CookieService.ClearAuthCookies(w)
	if err != nil {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}
	resp := &app.LoginResponse{
		Message: "Success"}
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResp)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) RefreshToken(w http.ResponseWriter, r *http.Request) {
	tokenCookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, "Refresh token not found", http.StatusUnauthorized)
		return
	}

	accesstoken, logresp, err := h.AuthService.Refresh(r.Context(), tokenCookie.Value)
	if err != nil {
		http.Error(w, "Failed to refresh token", http.StatusInternalServerError)
		return
	}
	h.AuthService.CookieService.SetAccessTokenCookie(w, accesstoken)
	jsonResp, err := json.Marshal(logresp)
	if err != nil {
		http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResp)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
func (h *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	var registerData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		RoleID   int    `json:"role_id"`
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&registerData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	logresp, err := h.AuthService.Register(r.Context(), registerData.Email, registerData.Password, registerData.RoleID)
	if err != nil {
		http.Error(w, "Failed to register", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(logresp)
	if err != nil {
		http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResp)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

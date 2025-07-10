package handlers

import (
	"encoding/json"
	"log"
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
	logresp, access, refresh, err := h.AuthService.Login(r.Context(), loginData.Email, loginData.Password)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		jsonResp, err := json.Marshal(logresp)
		if err != nil {
			http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(jsonResp)
		if err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
		return
	}
	if err := h.AuthService.CookieService.SetAccessTokenCookie(w, access); err != nil {
		log.Println("setting to cookie failed")
		return
	}
	log.Println("setting to cookie")
	if err := h.AuthService.CookieService.SetRefreshTokenCookie(w, refresh); err != nil {
		log.Println("setting to cookie failed")
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
	tokenCookie, err := r.Cookie(h.AuthService.CookieService.Cookie.RefreshToken.Name)
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
		log.Printf("err parsing %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	logresp, access, refresh, err := h.AuthService.Register(r.Context(), registerData.Email, registerData.Password, registerData.RoleID)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		jsonResp, err := json.Marshal(logresp)
		if err != nil {
			http.Error(w, "Failed to create JSON response", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(jsonResp)
		if err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
		return
	}
	if err := h.AuthService.CookieService.SetAccessTokenCookie(w, access); err != nil {
		return
	}
	if err := h.AuthService.CookieService.SetRefreshTokenCookie(w, refresh); err != nil {
		return
	}
	w.Header().Add("Content-Type", "application/json")
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

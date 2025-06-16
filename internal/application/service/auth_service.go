package app

import (
	"context"
	"net/http"

	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/config"
	auth "gitlab.pg.innopolis.university/f.markin/fah/auth/internal/domain/auth/service"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/cookies"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/jwt"
)

type LoginResponse struct {
	Message string `json:"message"`
}
type AuthService struct {
	DomainService *auth.Service
	JwtService    *jwt.JWTserv
	CookieService *cookies.CookiesService
	Config        *config.Config
}

func NewAuthService(domainService *auth.Service, jwtService *jwt.JWTserv, cookieService *cookies.CookiesService, config config.Config) *AuthService {
	return &AuthService{
		DomainService: domainService,
		JwtService:    jwtService,
		CookieService: cookieService,
		Config:        &config,
	}
}
func (s *AuthService) Register(ctx context.Context, email string, password string, roleId int) (*LoginResponse, error) {
	user, err := s.DomainService.CreateUser(ctx, *s.Config, email, password, roleId)
	if err != nil {
		return &LoginResponse{Message: "Fail"}, err
	}
	accessToken, err := s.JwtService.GenerateAccessToken(ctx, user.Email, user.RoleID, user.ID)
	if err != nil {
		return &LoginResponse{Message: "Fail"}, err
	}
	refreshToken, err := s.JwtService.GenerateRefreshToken(ctx, user.Email, user.RoleID)
	if err != nil {
		return &LoginResponse{Message: "Fail"}, err
	}
	if err := s.CookieService.SetAccessTokenCookie(ctx.Value("response").(http.ResponseWriter), accessToken); err != nil {
		return &LoginResponse{Message: "Fail"}, err
	}
	if err := s.CookieService.SetRefreshTokenCookie(ctx.Value("response").(http.ResponseWriter), refreshToken); err != nil {
		return &LoginResponse{Message: "Fail"}, err
	}
	return &LoginResponse{Message: "Success"}, nil
}
func (s *AuthService) Login(ctx context.Context, email string, password string) (*LoginResponse, error) {
	user, err := s.DomainService.LoginUser(ctx, email, password)
	if err != nil {
		return &LoginResponse{Message: "Fail"}, err
	}
	if user == nil {
		return &LoginResponse{Message: "User not found"}, nil
	}
	accessToken, err := s.JwtService.GenerateAccessToken(ctx, user.Email, user.RoleID, user.ID)
	if err != nil {
		return &LoginResponse{Message: "Fail"}, err
	}
	refreshToken, err := s.JwtService.GenerateRefreshToken(ctx, user.Email, user.RoleID)
	if err != nil {
		return &LoginResponse{Message: "Fail"}, err
	}
	if err := s.CookieService.SetAccessTokenCookie(ctx.Value("response").(http.ResponseWriter), accessToken); err != nil {
		return &LoginResponse{Message: "Fail"}, err
	}
	if err := s.CookieService.SetRefreshTokenCookie(ctx.Value("response").(http.ResponseWriter), refreshToken); err != nil {
		return &LoginResponse{Message: "Fail"}, err
	}
	return &LoginResponse{Message: "Success"}, nil
}
func (s *AuthService) Refresh(ctx context.Context, token string) (string, *LoginResponse, error) {
	claims, err := s.JwtService.ValidateRefreshToken(ctx, token)
	if err != nil {
		return "", &LoginResponse{Message: "Fail"}, err
	}
	accessToken, err := s.JwtService.GenerateAccessToken(ctx, claims.Email, claims.Role, claims.UserID)
	if err != nil {
		return "", &LoginResponse{Message: "Fail"}, err
	}

	return accessToken, &LoginResponse{Message: "Success"}, nil
}

package app

import (
	"context"
	"errors"
	"log"

	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/config"
	auth "gitlab.pg.innopolis.university/f.markin/fah/auth/internal/domain/auth/service"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/cookies"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/jwt"
	errs "gitlab.pg.innopolis.university/f.markin/fah/auth/shared/errors"
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
func (s *AuthService) Register(ctx context.Context, email string, password string, roleId int) (*LoginResponse, string, string, error) {
	log.Println("Creating user")
	user, err := s.DomainService.CreateUser(ctx, *s.Config, email, password, roleId)
	if err != nil {
		log.Println("Creating user failed")
		if errors.Is(err, errs.ErrInvalidMail) {
			return &LoginResponse{Message: "Invalid mail"}, "", "", err
		} else if errors.Is(err, errs.ErrPassTooLong) {
			return &LoginResponse{Message: "Password is too long"}, "", "", err
		} else if errors.Is(err, errs.ErrPassTooShort) {
			return &LoginResponse{Message: "Password is too short"}, "", "", err
		} else if errors.Is(err, errs.ErrUserExists) {
			return &LoginResponse{Message: "User is already exist"}, "", "", err
		}
	}
	log.Println("generating access token")
	accessToken, err := s.JwtService.GenerateAccessToken(ctx, user.Email, user.RoleID, user.ID)
	if err != nil {
		log.Println("generating access token failed")
		return &LoginResponse{Message: "Fail"}, "", "", err
	}
	log.Println("generating refresh token")
	refreshToken, err := s.JwtService.GenerateRefreshToken(ctx, user.Email, user.RoleID)
	if err != nil {
		log.Println("generating refresh token failed")
		return &LoginResponse{Message: "Fail"}, "", "", err
	}
	log.Println("setting to cookie")
	return &LoginResponse{Message: "Success"}, accessToken, refreshToken, nil
}
func (s *AuthService) Login(ctx context.Context, email string, password string) (*LoginResponse, string, string, error) {
	user, err := s.DomainService.LoginUser(ctx, email, password)
	if err != nil {
		if errors.Is(err, errs.ErrWrongLogOrPass) {
			return &LoginResponse{Message: "wrong login or password"}, "", "", err
		}
	}
	if user == nil {
		return &LoginResponse{Message: "wrong login or password"}, "", "", errs.ErrWrongLogOrPass
	}
	accessToken, err := s.JwtService.GenerateAccessToken(ctx, user.Email, user.RoleID, user.ID)
	if err != nil {
		return &LoginResponse{Message: "Fail"}, "", "", err
	}
	refreshToken, err := s.JwtService.GenerateRefreshToken(ctx, user.Email, user.RoleID)
	if err != nil {
		return &LoginResponse{Message: "Fail"}, "", "", err
	}
	return &LoginResponse{Message: "Success"}, accessToken, refreshToken, nil

}

func (s *AuthService) Refresh(ctx context.Context, token string) (string, *LoginResponse, error) {
	claims, err := s.JwtService.ValidateRefreshToken(ctx, token)
	if err != nil {
		log.Printf("Err validating: %v", err)
		return "", &LoginResponse{Message: "Fail"}, err
	}
	accessToken, err := s.JwtService.GenerateAccessToken(ctx, claims.Email, claims.Role, claims.UserID)
	if err != nil {
		log.Printf("Err generating: %v", err)
		return "", &LoginResponse{Message: "Fail"}, err
	}

	return accessToken, &LoginResponse{Message: "Success"}, nil
}

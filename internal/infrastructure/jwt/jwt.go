package jwt

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/config"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/database/postgres"
)

type JWTserv struct {
	repo postgres.Repository
	JWT  config.JWT
}

func NewJWTService(repo postgres.Repository, jwtConfig config.JWT) *JWTserv {
	return &JWTserv{
		repo: repo,
		JWT:  jwtConfig,
	}
}

type Claims struct {
	UserID string `json:"user_id,omitempty"`
	Email  string `json:"email"`
	Role   int    `json:"role"`
	jwt.RegisteredClaims
}

func generateUniqueTokenID() string {
	return uuid.New().String()
}

// уточнить, правильно ли создавать так access token? разве он не генерируется по уже имеющемуся refresh token? (тип используя refresh создаем access)
func (s *JWTserv) GenerateAccessToken(ctx context.Context, email string, roleID int, userID string) (string, error) {
	if email == "" || roleID <= 0 || userID == "" {
		return "", jwt.ErrInvalidKey
	}
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	Claims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.RoleID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "fah-auth",
			Subject:   user.Email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.JWT.AccessToken.Expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			// роль пользователя
			ID: generateUniqueTokenID()},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(s.JWT.AccessToken.Algorithm), Claims)
	tokenString, err := token.SignedString([]byte(s.JWT.AccessToken.Secret))
	if err != nil {
		log.Printf("err creating: %v", err)
		return "", err
	}
	return tokenString, nil
}
func (s *JWTserv) GenerateRefreshToken(ctx context.Context, email string, role_id int) (string, error) {
	if email == "" {
		return "", jwt.ErrInvalidKey
	}
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	Claims := &Claims{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.RoleID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.JWT.RefreshToken.Expiry)), // Different expiry!
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "your-app",
			Subject:   user.Email,
			ID:        generateUniqueTokenID(),
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(s.JWT.RefreshToken.Algorithm), Claims)
	tokenString, err := token.SignedString([]byte(s.JWT.RefreshToken.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func (s *JWTserv) ValidateAccessToken(ctx context.Context, token string) (*Claims, error) {
	if token == "" {
		return nil, errors.New("token is empty")
	}
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != s.JWT.AccessToken.Algorithm {
			return nil, jwt.ErrInvalidKey
		}
		return []byte(s.JWT.AccessToken.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := parsedToken.Claims.(*Claims); ok && parsedToken.Valid && claims.Email != "" && claims.Role > 0 {
		return claims, nil
	}
	return nil, jwt.ErrInvalidKey
}
func (s *JWTserv) ValidateRefreshToken(ctx context.Context, token string) (*Claims, error) {
	if token == "" {
		return nil, errors.New("token is empty")
	}
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != s.JWT.RefreshToken.Algorithm {
			return nil, jwt.ErrInvalidKey
		}
		return []byte(s.JWT.RefreshToken.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := parsedToken.Claims.(*Claims); ok && parsedToken.Valid && claims.Email != "" && claims.Role > 0 {
		return claims, nil
	}
	return nil, jwt.ErrInvalidKey
}
func (s *JWTserv) ExtractClaims(ctx context.Context, token string) (*Claims, error) {
	if token == "" {

		return nil, errors.New("token is empty")
	}
	parsedToken, _, err := jwt.NewParser().ParseUnverified(token, &Claims{})
	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(*Claims); ok {
		return claims, nil
	}
	return nil, jwt.ErrInvalidKey
}
func (s *JWTserv) IsTokenExpired(ctx context.Context, token string) (bool, error) {
	if token == "" {
		return false, errors.New("token is empty")
	}
	Claims, err := s.ExtractClaims(ctx, token)
	if err != nil {
		return true, err
	}
	if Claims.ExpiresAt != nil && Claims.ExpiresAt.Time.Before(time.Now()) {
		return true, nil // Token IS expired
	}
	return false, nil
}
func (s *JWTserv) GetTokenExpiry(ctx context.Context, token string) (time.Time, error) {
	if token == "" {
		// выглядит как кринж, посмотреть, что есть в нативной либе jwt
		return time.Time{}, errors.New("token is empty")
	}
	Claims, err := s.ExtractClaims(ctx, token)
	if err != nil {
		return time.Time{}, err
	}
	if Claims.ExpiresAt == nil {
		// выглядит как кринж, посмотреть, что есть в нативной либе jwt
		return time.Time{}, errors.New("token does not have an expiry time")
	}
	return Claims.ExpiresAt.Time, nil
}

package auth

import (
	"context"
	"errors"
	"log"
	"strings"

	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/config"
	"gitlab.pg.innopolis.university/f.markin/fah/auth/internal/infrastructure/database/postgres"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo postgres.Repository
}

func NewService(repo postgres.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
func (s *Service) CreateUser(ctx context.Context, config config.Config, email string, password string, roleId int) (*postgres.User, error) {
	log.Println("validing mail")
	if !ValidateEmail(email) {
		log.Println("ivalid")
		return nil, errors.New("not inno mail")
	}
	log.Println("validing pass")
	if len(password) < 8 {
		log.Println("ivalid")
		return nil, errors.New("too small password")
	}
	log.Println("hashing pass")
	passhash, err := bcrypt.GenerateFromPassword([]byte(password), config.Sequrity.Bcrypt.Cost)
	if err != nil {
		log.Println("failed to hashing pass")
		return nil, err
	}
	log.Println("creating user")
	user := postgres.NewUser(email, roleId, passhash)
	log.Println("creating user in db")
	if err := s.repo.CreateUser(ctx, *user); err != nil {
		log.Println("creating user in db failed")
		return nil, err
	}
	return user, nil
}
func (s *Service) GetUserByEmail(ctx context.Context, email string) (*postgres.User, error) {
	if !ValidateEmail(email) {
		return nil, errors.New("Not inno mail")
	}
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *Service) LoginUser(ctx context.Context, email string, password string) (*postgres.User, error) {
	if !ValidateEmail(email) {
		return nil, errors.New("not inno mail")
	}
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password)); err != nil {
		return nil, bcrypt.ErrMismatchedHashAndPassword
	}
	return user, nil
}

func ValidateEmail(email string) bool {
	ind := strings.Index(email, "@")
	if email[ind+1:] == "innopolis.university" {
		return true
	}
	return false
}

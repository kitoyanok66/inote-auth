package auth

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/kitoyanok66/inote-auth/internal/auth/domain"
	"github.com/kitoyanok66/inote-auth/internal/auth/dto"
	"github.com/kitoyanok66/inote-auth/internal/auth/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, email, username, password string) (*dto.UserDTO, string, error)
	Login(ctx context.Context, email, password string) (*dto.UserDTO, string, error)
	GetUser(ctx context.Context, id string) (*dto.UserDTO, error)
}

type authService struct {
	repo AuthRepository
	jwt  jwt.JWTManager
}

func NewAuthService(repo AuthRepository, jwt jwt.JWTManager) AuthService {
	return &authService{repo: repo, jwt: jwt}
}

func (s *authService) Register(ctx context.Context, email, username, password string) (*dto.UserDTO, string, error) {
	_, err := s.repo.GetByEmail(ctx, email)
	if err == nil {
		return nil, "", errors.New("user already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}

	user := &domain.User{
		ID:        uuid.New().String(),
		Email:     email,
		Username:  username,
		Password:  string(hashed),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, "", err
	}

	token, err := s.jwt.Generate(user.ID, user.Email)
	if err != nil {
		return nil, "", err
	}

	return ToDTO(user), token, nil
}

func (s *authService) Login(ctx context.Context, email, password string) (*dto.UserDTO, string, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	token, err := s.jwt.Generate(user.ID, user.Email)
	if err != nil {
		return nil, "", err
	}

	return ToDTO(user), token, nil
}

func (s *authService) GetUser(ctx context.Context, id string) (*dto.UserDTO, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return ToDTO(user), nil
}

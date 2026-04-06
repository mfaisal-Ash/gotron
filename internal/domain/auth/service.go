package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo      Repository
	JWTSecret string
}

func NewService(repo Repository, secret string, tokenHours int) *service {
	return &service{
		repo:      repo,
		JWTSecret: secret,
	}
}

func (s *Service) Register(req RegisterRequest) (*AuthResponse, error) {
	if req.FullName == "" || req.Email == "" || req.Password == "" {
		return nil, errors.New("full_name, email, and password are required")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:           fmt.Sprintf("usr_%d", time.Now().UnixNano()),
		FullName:     req.FullName,
		Email:        req.Email,
		PasswordHash: string(hash),
		Role:         "customer",
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	token, expiresIn, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		User: UserResponse{
			ID:       user.ID,
			FullName: user.FullName,
			Email:    user.Email,
			Role:     user.Role,
		},
		AccessToken: token,
		ExpiresIn:   expiresIn,
	}, nil
}

func (s *service) Login(req LoginRequest) (*AuthResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email and password are required")
	}

	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, expiresIn, err := s.generateToken(user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		User: UserResponse{
			ID:       user.ID,
			FullName: user.FullName,
			Email:    user.Email,
			Role:     user.Role,
		},
		AccessToken: token,
		ExpiresIn:   expiresIn,
	}, nil
}

func (s *service) Me(userID string) (*UserResponse, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &UserResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
	}, nil
}

func (s *service) generateToken(user *User) (string, int64, error) {
	expiresIn := time.Now().Add(24 * time.Hour).Unix()
	claims := jwt.MapClaims{
		"user_id":   user.ID,
		"full_name": user.FullName,
		"email":     user.Email,
		"role":      user.Role,
		"exp":       expiresIn,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokString, err := token.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return "", 0, err
	}

	return tokString, expiresIn, nil
}

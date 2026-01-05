package usecase

import (
	"errors"

	"github.com/herlianali/goCommerce/internal/domain/entity"
	"github.com/herlianali/goCommerce/internal/domain/repository"
	"github.com/herlianali/goCommerce/internal/infrastructure/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewAuthUsecase(userRepo repository.UserRepository, jwtSecret string) *AuthUsecase {
	return &AuthUsecase{userRepo: userRepo, jwtSecret: jwtSecret}
}

// Register user baru
func (a *AuthUsecase) Register(user *entity.User) error {
	_, err := a.userRepo.FindByEmail(user.Email)
	if err == nil {
		return errors.New("email already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashed)

	return a.userRepo.Create(user)
}

// Login user
func (a *AuthUsecase) Login(email, password string) (string, error) {
	user, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("Invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("Invalid credentials")
	}

	token, err := jwt.GenerateToken(user.ID, user.Role, a.jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

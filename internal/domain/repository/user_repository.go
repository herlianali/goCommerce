package repository

import "github.com/herlianali/goCommerce/internal/domain/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

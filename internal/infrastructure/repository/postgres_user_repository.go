package repository

import (
	"database/sql"
	"errors"

	"github.com/herlianali/goCommerce/internal/domain/entity"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(user *entity.User) error {
	query := `INSERT INTO user (name, email, password, role, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id`
	err := r.db.QueryRow(query, user.Name, user.Email, user.Password, user.Role).Scan(&user.ID)
	return err
}

func (r *PostgresUserRepository) FindByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	query := `SELECT id, name, email, password, role, created_at FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("User not Found!")
	}

	return user, err
}

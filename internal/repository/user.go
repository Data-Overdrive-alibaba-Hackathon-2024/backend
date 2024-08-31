package repository

import (
	"database/sql"
	"github.com/data-overdrive-alibaba-hackathon-2024/internal/model"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type userRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

type UserRepository interface {
	CreateUser(input model.CreateUserInput) error
	GetUserByEmail(email string) (model.User, error)
}

func NewUserRepository(db *sql.DB, logger *zap.Logger) UserRepository {
	return &userRepository{
		db:     db,
		logger: logger,
	}
}

func (r *userRepository) CreateUser(input model.CreateUserInput) error {
	_, err := r.db.Exec(`
		INSERT INTO users (id, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
	`, uuid.NewString(), input.Email, input.Password)
	if err != nil {
		r.logger.Error("failed to create user", zap.Error(err))
		return err
	}
	return nil
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.QueryRow(`
		SELECT id, email, password
		FROM users
		WHERE email = $1
	`, email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		r.logger.Error("failed to get user by email", zap.Error(err))
		return user, err
	}
	return user, nil
}

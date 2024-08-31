package service

import (
	"github.com/data-overdrive-alibaba-hackathon-2024/internal/model"
	"github.com/data-overdrive-alibaba-hackathon-2024/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

type userService struct {
	userRepository repository.UserRepository
	logger         *zap.Logger
}

type UserService interface {
	CreateUser(input model.CreateUserInput) error
	Login(email, password string) (string, error)
}

func NewUserService(userRepository repository.UserRepository, logger *zap.Logger) UserService {
	return &userService{
		userRepository: userRepository,
		logger:         logger,
	}
}

func (s *userService) CreateUser(input model.CreateUserInput) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	input.Password = string(hashedPassword)
	return s.userRepository.CreateUser(input)
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRETKEY")))
	if err != nil {
		log.Println("error: " + err.Error())
		return "", err
	}

	return tokenString, nil
}

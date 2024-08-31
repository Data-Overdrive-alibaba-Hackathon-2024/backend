package handler

import (
	"github.com/data-overdrive-alibaba-hackathon-2024/internal/model"
	"github.com/data-overdrive-alibaba-hackathon-2024/internal/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type userHandler struct {
	userService service.UserService
	logger      *zap.Logger
}

type UserHandler interface {
	CreateUser(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

func NewUserHandler(userService service.UserService, logger *zap.Logger) UserHandler {
	return &userHandler{
		userService: userService,
		logger:      logger,
	}
}

func (h *userHandler) CreateUser(c *fiber.Ctx) error {
	var input model.CreateUserRequest
	if err := c.BodyParser(&input); err != nil {
		h.logger.Error("failed to parse request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "failed to parse request body",
		})
	}

	if err := h.userService.CreateUser(model.CreateUserInput{
		Email:    input.Email,
		Password: input.Password,
	}); err != nil {
		h.logger.Error("failed to create user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "user created",
	})
}

func (h *userHandler) Login(c *fiber.Ctx) error {
	var input model.LoginRequest
	if err := c.BodyParser(&input); err != nil {
		h.logger.Error("failed to parse request body", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "failed to parse request body",
		})
	}

	token, err := h.userService.Login(input.Email, input.Password)
	if err != nil {
		h.logger.Error("failed to login", zap.Error(err))
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "failed",
			"message": "failed to login",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "login success",
		"token":   token,
	})
}

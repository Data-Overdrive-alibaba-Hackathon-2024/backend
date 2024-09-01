package handler

import (
	"github.com/data-overdrive-alibaba-hackathon-2024/internal/model"
	"github.com/data-overdrive-alibaba-hackathon-2024/internal/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type userHandler struct {
	userService     service.UserService
	logger          *zap.Logger
	questionHandler QuestionHandler
	questionService service.QuestionService
}

type UserHandler interface {
	CreateUser(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
}

func NewUserHandler(userService service.UserService, logger *zap.Logger, questionHandler QuestionHandler, questionService service.QuestionService) UserHandler {
	return &userHandler{
		userService:     userService,
		logger:          logger,
		questionHandler: questionHandler,
		questionService: questionService,
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

	user, err := h.userService.GetUserByEmail(input.Email)
	if err != nil {
		h.logger.Error("failed to get user by email", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "failed to get user by email",
		})
	}

	question, err := h.questionHandler.RequestAI(1)
	if err != nil {
		h.logger.Error("failed to generate question", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "failed to generate question",
		})
	}

	if err := h.questionService.InsertQuestion(model.InsertQuestionInput{
		UserId:        user.Id,
		Level:         1,
		Question:      question.Question,
		Options:       question.Options,
		CorrectAnswer: question.CorrectAnswer,
	}); err != nil {
		h.logger.Error("failed to insert question", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "failed to insert question",
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

func (h *userHandler) GetUser(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)

	user, err := h.userService.GetUserById(userId)
	if err != nil {
		h.logger.Error("failed to get user", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "failed to get user",
		})
	}
	user.Password = ""

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   user,
	})
}

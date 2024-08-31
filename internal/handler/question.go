package handler

import (
	"github.com/data-overdrive-alibaba-hackathon-2024/internal/model"
	"github.com/data-overdrive-alibaba-hackathon-2024/internal/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type questionHandler struct {
	questionService service.QuestionService
	logger          *zap.Logger
}

type QuestionHandler interface {
	GenerateQuestion(c *fiber.Ctx) error
	RequestAI(level int) (model.GenerateQuestionAIResponse, error)
	GetQuestion(c *fiber.Ctx) error
	UpdateQuestionDone(c *fiber.Ctx) error
}

func NewQuestionHandler(questionService service.QuestionService, logger *zap.Logger) QuestionHandler {
	return &questionHandler{
		questionService: questionService,
		logger:          logger,
	}
}

func (h *questionHandler) GenerateQuestion(c *fiber.Ctx) error {
	var req model.GenerateQuestionRequest

	if err := c.BodyParser(&req); err != nil {
		h.logger.Error("failed to parse request", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "invalid request",
		})
	}

	questionText, err := h.RequestAI(req.Level)
	if err != nil {
		h.logger.Error("failed to generate question", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "internal server error " + err.Error(),
		})
	}

	if err := h.questionService.InsertQuestion(model.InsertQuestionInput{
		UserId:        req.UserId,
		Level:         req.Level,
		Question:      questionText.Question,
		Options:       questionText.Options,
		CorrectAnswer: questionText.CorrectAnswer,
	}); err != nil {
		h.logger.Error("failed to insert question", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "internal server error " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "question generated",
	})
}

func (h *questionHandler) GetQuestion(c *fiber.Ctx) error {
	userId := c.Params("user_id")
	level := c.QueryInt("lv")

	question, err := h.questionService.GetQuestion(model.GetQuestionInput{
		UserId: userId,
		Level:  level,
	})
	if err != nil {
		h.logger.Error("failed to get question", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "internal server error " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   question,
	})
}

func (h *questionHandler) UpdateQuestionDone(c *fiber.Ctx) error {
	questionId := c.Params("question_id")

	if err := h.questionService.UpdateQuestionDone(questionId); err != nil {
		h.logger.Error("failed to update question done", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "internal server error " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "question updated",
	})
}

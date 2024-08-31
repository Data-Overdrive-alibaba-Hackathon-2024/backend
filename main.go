package main

import (
	"github.com/data-overdrive-alibaba-hackathon-2024/config"
	"github.com/data-overdrive-alibaba-hackathon-2024/internal/handler"
	"github.com/data-overdrive-alibaba-hackathon-2024/internal/repository"
	"github.com/data-overdrive-alibaba-hackathon-2024/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("failed to load env", zap.Error(err))
	}

	db := config.NewDBPool()
	defer db.Close()

	questionRepository := repository.NewQuestionRepository(db, logger)
	questionService := service.NewQuestionService(questionRepository, logger)
	questionHandler := handler.NewQuestionHandler(questionService, logger)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/generate/questions", questionHandler.GenerateQuestion)

	app.Listen(":3000")
}

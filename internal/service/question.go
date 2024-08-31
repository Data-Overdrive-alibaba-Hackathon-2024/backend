package service

import (
	"github.com/data-overdrive-alibaba-hackathon-2024/internal/model"
	"github.com/data-overdrive-alibaba-hackathon-2024/internal/repository"
	"go.uber.org/zap"
)

type questionService struct {
	questionRepository repository.QuestionRepository
	logger             *zap.Logger
}

type QuestionService interface {
	InsertQuestion(input model.InsertQuestionInput) error
}

func NewQuestionService(questionRepository repository.QuestionRepository, logger *zap.Logger) QuestionService {
	return &questionService{
		questionRepository: questionRepository,
		logger:             logger,
	}
}

func (s *questionService) InsertQuestion(input model.InsertQuestionInput) error {
	return s.questionRepository.InsertQuestion(input)
}

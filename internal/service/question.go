package service

import (
	"github.com/data-overdrive-alibaba-hackathon-2024/internal/model"
	"github.com/data-overdrive-alibaba-hackathon-2024/internal/repository"
	"go.uber.org/zap"
)

type questionService struct {
	questionRepository repository.QuestionRepository
	userRepository     repository.UserRepository
	logger             *zap.Logger
}

type QuestionService interface {
	InsertQuestion(input model.InsertQuestionInput) error
	GetQuestion(input model.GetQuestionInput) (model.GetQuestionOutput, error)
	UpdateQuestionDone(questionId string) error
	ResetQuestionAndLevel(userId string) error
}

func NewQuestionService(questionRepository repository.QuestionRepository, userRepository repository.UserRepository, logger *zap.Logger) QuestionService {
	return &questionService{
		questionRepository: questionRepository,
		userRepository:     userRepository,
		logger:             logger,
	}
}

func (s *questionService) InsertQuestion(input model.InsertQuestionInput) error {
	return s.questionRepository.InsertQuestion(input)
}

func (s *questionService) GetQuestion(input model.GetQuestionInput) (model.GetQuestionOutput, error) {
	return s.questionRepository.GetQuestion(input)
}

func (s *questionService) UpdateQuestionDone(questionId string) error {
	return s.questionRepository.UpdateQuestionDone(questionId)
}

func (s *questionService) ResetQuestionAndLevel(userId string) error {
	if err := s.questionRepository.DeleteAllQuestionByUserId(userId); err != nil {
		s.logger.Error("failed to delete all question by user id", zap.Error(err))
		return err
	}

	if err := s.userRepository.UpdateUserLevel(userId, 1); err != nil {
		s.logger.Error("failed to update user level", zap.Error(err))
		return err
	}

	return nil
}

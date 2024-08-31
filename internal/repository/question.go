package repository

import (
	"database/sql"
	"github.com/data-overdrive-alibaba-hackathon-2024/internal/model"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

type questionRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

type QuestionRepository interface {
	InsertQuestion(input model.InsertQuestionInput) error
}

func NewQuestionRepository(db *sql.DB, logger *zap.Logger) QuestionRepository {
	return &questionRepository{db: db, logger: logger}
}

func (r *questionRepository) InsertQuestion(input model.InsertQuestionInput) error {
	_, err := r.db.Exec(`
		INSERT INTO questions (id, user_id, question, level, option_1, option_2, option_3, option_4, correct_answer, done, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
		uuid.NewString(), input.UserId, input.Question, input.Level, input.Options.Option1, input.Options.Option2, input.Options.Option3, input.Options.Option4,
		input.CorrectAnswer, false, time.Now(), time.Now())
	if err != nil {
		r.logger.Error("failed to insert question", zap.Error(err))
		return err
	}

	return nil
}

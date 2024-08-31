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
	GetQuestion(input model.GetQuestionInput) (model.GetQuestionOutput, error)
	UpdateQuestionDone(questiionId string) error
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

func (r *questionRepository) GetQuestion(input model.GetQuestionInput) (model.GetQuestionOutput, error) {
	var output model.GetQuestionOutput

	row := r.db.QueryRow(`
		SELECT id, question, level, option_1, option_2, option_3, option_4, correct_answer, done
		FROM questions
		WHERE user_id = $1 AND level = $2
	`, input.UserId, input.Level)
	if err := row.Scan(&output.Id, &output.Question, &output.Level, &output.Option1, &output.Option2, &output.Option3,
		&output.Option4, &output.CorrectAnswer, &output.Done); err != nil {
		r.logger.Error("failed to get question", zap.Error(err))
		return model.GetQuestionOutput{}, err
	}

	return output, nil
}

func (r *questionRepository) UpdateQuestionDone(questiionId string) error {
	_, err := r.db.Exec(`
		UPDATE questions
		SET done = true
		WHERE id = $1
	`, questiionId)
	if err != nil {
		r.logger.Error("failed to update question done", zap.Error(err))
		return err
	}

	return nil
}

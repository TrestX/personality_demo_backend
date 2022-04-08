package quiz_db

import "personality_demo_backend/pkg/entity"

type QuizService interface {
	AddQuiz(*entity.QuizEntity) (string, error)
	GetQuiz(quizType string) ([]entity.QuizEntity, error)
}

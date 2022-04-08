package quiz_type_db

import "personality_demo_backend/pkg/entity"

type QuizTypeService interface {
	AddQuizType(*entity.QuizTypeEntity) (string, error)
	GetQuizType() ([]entity.QuizTypeEntity, error)
}

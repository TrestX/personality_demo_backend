package quiz_type_db

import (
	"personality_demo_backend/pkg/entity"
	"personality_demo_backend/pkg/repository"
	"personality_demo_backend/pkg/repository/quiz_type"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo  = quiz_type.NewQuizTypeRepository("quiz_type")
	repo2 = repository.NewRepository("quiz_type")
)

type quizTypeService struct{}

func NewQuizTypeService(repository quiz_type.QuizTypeRepository) QuizTypeService {
	repo = repository
	return &quizTypeService{}
}

// AddQuizType implements QuizTypeService
func (*quizTypeService) AddQuizType(quiz *entity.QuizTypeEntity) (string, error) {
	quiz.ID = primitive.NewObjectID()

	return repo2.InsertOne(quiz)
}

// GetQuizType implements QuizTypeService
func (*quizTypeService) GetQuizType() ([]entity.QuizTypeEntity, error) {
	return repo.Find(bson.M{}, bson.M{})
}

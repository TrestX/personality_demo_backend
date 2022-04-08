package quiz_db

import (
	"personality_demo_backend/pkg/entity"
	"personality_demo_backend/pkg/repository"
	"personality_demo_backend/pkg/repository/quiz"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo  = quiz.NewQuizRepository("quiz")
	repo2 = repository.NewRepository("quiz")
)

type quizService struct{}

// AddQuiz implements QuizService
func (*quizService) AddQuiz(quiz *entity.QuizEntity) (string, error) {
	quiz.ID = primitive.NewObjectID()
	return repo2.InsertOne(quiz)
}

// GetQuiz implements QuizService
func (*quizService) GetQuiz(quizType string) ([]entity.QuizEntity, error) {
	return repo.Find(bson.M{"quiz_type": quizType}, bson.M{})
}
func GetQuizInternally(quizType string) ([]entity.QuizEntity, error) {
	return repo.Find(bson.M{"quiz_type": quizType}, bson.M{})
}

func NewQuizService(repository quiz.QuizRepository) QuizService {
	repo = repository
	return &quizService{}
}

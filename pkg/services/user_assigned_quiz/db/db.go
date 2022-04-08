package user_assigned_db

import (
	"personality_demo_backend/pkg/entity"
	"personality_demo_backend/pkg/repository"
	user_assigned_quiz_repo "personality_demo_backend/pkg/repository/user_assigned_quiz"
	quiz_db "personality_demo_backend/pkg/services/quiz/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo  = user_assigned_quiz_repo.NewRepository("user_assigned")
	repo2 = repository.NewRepository("user_assigned")
)

type userAssignedService struct{}

// AssignUser implements UserAssignedService
func (uAS *userAssignedService) AssignUser(userData *entity.UserAssignedEntity) (string, error) {
	da, err := uAS.GetAssignedQuesion(userData.QuizTypeID, userData.UserID)
	if err == nil {
		return da.ID.Hex(), nil
	}
	quizes, err := quiz_db.GetQuizInternally(userData.QuizTypeID)
	if err != nil {
		return "", err
	}
	if len(quizes) > 5 {
		userData.Quiz = quizes[:5]
	} else {
		userData.Quiz = quizes
	}
	userData.ID = primitive.NewObjectID()
	return repo2.InsertOne(userData)
}

// GetAssignedQuesion implements UserAssignedService
func (*userAssignedService) GetAssignedQuesion(quizType, userID string) (entity.UserAssignedEntity, error) {
	return repo.FindOne(bson.M{"quiz_type_id": quizType, "user_id": userID}, bson.M{})
}
func (uAS *userAssignedService) UpdateQuizValue(quesID, answer string, quizNumber int) (string, error) {
	id, _ := primitive.ObjectIDFromHex(quesID)
	filter := bson.M{"_id": id}
	data, err := repo.FindOne(filter, bson.M{})
	if err != nil {
		return "", err
	}
	questions := data.Quiz
	for i := range questions {
		if i == quizNumber {
			questions[i].SelectedAnswer = answer
		}
	}
	return repo2.UpdateOne(filter, bson.M{"$set": bson.M{"quiz_id": questions}})
}

func NewQuizService(repository user_assigned_quiz_repo.UserAssignedQuizTypeRepository) UserAssignedService {
	repo = repository
	return &userAssignedService{}
}

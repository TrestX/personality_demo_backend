package user_progress_db

import (
	"personality_demo_backend/pkg/entity"
	"personality_demo_backend/pkg/repository"
	"personality_demo_backend/pkg/repository/user_progress"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	repo  = user_progress.NewRepository("user_progress")
	repo2 = repository.NewRepository("user_progress")
)

type userProgressService struct{}

// GetProgress implements UserProgressService
func (*userProgressService) GetProgress(userID string) (entity.UserProgressEntity, error) {
	return repo.FindOne(bson.M{"user_id": userID}, bson.M{})
}

// UpdateUserProgress implements UserProgressService
func (*userProgressService) UpdateUserProgress(userProgress *entity.UserProgressEntity) (string, error) {
	response, err := repo.FindOne(bson.M{"user_id": userProgress.UserID}, bson.M{})
	if err != nil {
		userProgress.ID = primitive.NewObjectID()
		return repo2.InsertOne(userProgress)
	}
	completed := response.Completed
	if userProgress.Progress[0].Done == 5 {
		completed = append(completed, userProgress.Progress[0].QuizType)
	} else {
		progress := response.Progress
		if len(progress) > 0 {
			updated := false
			for i := range progress {
				if progress[i].QuizType == userProgress.Progress[0].QuizType {
					progress[i].Done = userProgress.Progress[0].Done
					updated = true
				}
			}
			if !updated {
				progress = append(progress, userProgress.Progress[0])
			}
		} else {
			progress = append(progress, userProgress.Progress[0])
		}
		return repo2.UpdateOne(bson.M{"user_id": userProgress.UserID}, bson.M{"$set": bson.M{"progress": progress}})
	}

	return repo2.UpdateOne(bson.M{"user_id": userProgress.UserID}, bson.M{"$set": bson.M{"completed": completed}})
}

func NewQuizService(repository user_progress.UserProgressRepository) UserProgressService {
	repo = repository
	return &userProgressService{}
}

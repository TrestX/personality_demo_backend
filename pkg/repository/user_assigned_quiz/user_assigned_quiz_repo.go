package user_assigned_quiz_repo

import (
	"personality_demo_backend/pkg/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type UserAssignedQuizTypeRepository interface {
	FindOne(filter, projection bson.M) (entity.UserAssignedEntity, error)
	Find(filter, projection bson.M) ([]entity.UserAssignedEntity, error)
}

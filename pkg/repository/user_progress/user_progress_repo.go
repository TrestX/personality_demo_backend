package user_progress

import (
	"personality_demo_backend/pkg/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type UserProgressRepository interface {
	FindOne(filter, projection bson.M) (entity.UserProgressEntity, error)
	Find(filter, projection bson.M) ([]entity.UserProgressEntity, error)
}

package quiz_type

import (
	"personality_demo_backend/pkg/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type QuizTypeRepository interface {
	FindOne(filter, projection bson.M) (entity.QuizTypeEntity, error)
	Find(filter, projection bson.M) ([]entity.QuizTypeEntity, error)
}

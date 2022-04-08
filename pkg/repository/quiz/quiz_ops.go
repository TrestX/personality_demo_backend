package quiz

import (
	"personality_demo_backend/pkg/entity"

	"go.mongodb.org/mongo-driver/bson"
)

type QuizRepository interface {
	FindOne(filter, projection bson.M) (entity.QuizEntity, error)
	Find(filter, projection bson.M) ([]entity.QuizEntity, error)
}

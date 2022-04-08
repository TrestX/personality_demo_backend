package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserAssignedEntity struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	Quiz       []QuizEntity       `bson:"quiz_id" json:"quiz_id"`
	UserID     string             `bson:"user_id" json:"user_id"`
	QuizTypeID string             `bson:"quiz_type_id" json:"quiz_type_id"`
}

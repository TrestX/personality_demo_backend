package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserProgressEntity struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	UserID    string             `bson:"user_id" json:"user_id"`
	Completed []string           `bson:"completed" json:"completed"`
	Progress  []ProgressStruct   `bson:"progress" json:"progress"`
}

type ProgressStruct struct {
	QuizType string `bson:"quiz_type" json:"quiz_type"`
	Done     int    `bson:"done" json:"done"`
}

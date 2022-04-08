package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type QuizTypeEntity struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	TypeName string             `bson:"type_name" json:"type_name"`
}

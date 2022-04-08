package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type QuizEntity struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	QuizType       string             `bson:"quiz_type" json:"quiz_type"`
	Question       string             `bson:"question" json:"question"`
	SelectedAnswer string             `bson:"selected_answer" json:"selected_answer"`
	Options        []string           `bson:"options" json:"options"`
}

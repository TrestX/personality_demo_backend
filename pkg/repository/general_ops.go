package repository

import (
	"go.mongodb.org/mongo-driver/bson"
)

type GeneralRepository interface {
	InsertOne(document interface{}) (string, error)
	UpdateOne(filter, update bson.M) (string, error)
	DeleteOne(filter bson.M) error
}

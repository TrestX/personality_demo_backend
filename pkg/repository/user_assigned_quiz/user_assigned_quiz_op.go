package user_assigned_quiz_repo

import (
	"context"
	"personality_demo_backend/pkg/entity"

	"github.com/aekam27/trestCommon"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type repo struct {
	CollectionName string
}

//NewFirestoreRepository creates a new repo
func NewRepository(collectionName string) UserAssignedQuizTypeRepository {
	return &repo{
		CollectionName: collectionName,
	}
}

//used by get profile ,login and email verification
func (r *repo) FindOne(filter, projection bson.M) (entity.UserAssignedEntity, error) {
	var profile entity.UserAssignedEntity
	err := trestCommon.FindOne(filter, projection, r.CollectionName).Decode(&profile)
	if err != nil {
		trestCommon.ECLog3(
			"Find profile",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return profile, err
	}
	return profile, err
}

//not used may use in future for gettin list of profiles
func (r *repo) Find(filter, projection bson.M) ([]entity.UserAssignedEntity, error) {
	var profiles []entity.UserAssignedEntity
	cursor, err := trestCommon.FindSort(filter, projection, bson.M{"_id": -1}, 1000000, 0, r.CollectionName)
	if err != nil {
		trestCommon.ECLog3(
			"Find profiles",
			err,
			logrus.Fields{
				"filter":          filter,
				"collection name": r.CollectionName,
			})
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.TODO()) {
		var profile entity.UserAssignedEntity
		if err = cursor.Decode(&profile); err != nil {
			trestCommon.ECLog3(
				"Find profiles",
				err,
				logrus.Fields{
					"filter":          filter,
					"collection name": r.CollectionName,
					"error at":        cursor.RemainingBatchLength(),
				})
			return profiles, nil
		}
		profiles = append(profiles, profile)
	}
	return profiles, nil
}

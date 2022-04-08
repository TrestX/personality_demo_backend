package userassignedquiz_service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"personality_demo_backend/pkg/entity"
	user_assigned_quiz_repo "personality_demo_backend/pkg/repository/user_assigned_quiz"
	user_assigned_db "personality_demo_backend/pkg/services/user_assigned_quiz/db"
	"strconv"
	"time"

	"github.com/aekam27/trestCommon"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	service = user_assigned_db.NewQuizService(user_assigned_quiz_repo.NewRepository("user_assigned"))
)

func AssignUser(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting post", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var assignUser *entity.UserAssignedEntity
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &assignUser)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	data, err := service.AssignUser(assignUser)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set post"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set post"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("post updated", logrus.Fields{
		"duration": duration,
	})
}

func GetAssignedQuiestion(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting post", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	quizType := mux.Vars(r)["quizType"]
	userID := mux.Vars(r)["userID"]
	data, err := service.GetAssignedQuesion(quizType, userID)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set post"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set post"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("post updated", logrus.Fields{
		"duration": duration,
	})
}

func UpdateAssignedQuiestion(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting post", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	quizID := mux.Vars(r)["quizID"]
	answer := mux.Vars(r)["answer"]
	quizNumber, err := strconv.Atoi(mux.Vars(r)["quizNumber"])
	if err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "error occurred", "data": ""})
		return
	}
	data, err := service.UpdateQuizValue(quizID, answer, quizNumber)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to set post"))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Unable to set post"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "data": data})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("post updated", logrus.Fields{
		"duration": duration,
	})
}

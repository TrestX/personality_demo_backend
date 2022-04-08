package quiz_service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"personality_demo_backend/pkg/entity"
	"personality_demo_backend/pkg/repository/quiz"
	quiz_db "personality_demo_backend/pkg/services/quiz/db"
	"time"

	"github.com/aekam27/trestCommon"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	service = quiz_db.NewQuizService(quiz.NewQuizRepository("quiz"))
)

func AddQuiz(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("setting post", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var quiz *entity.QuizEntity
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &quiz)
	if err != nil {
		trestCommon.ECLog1(errors.Wrapf(err, "unable to unmarshal body"))
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(bson.M{"status": false, "error": "Something Went wrong"})
		return
	}
	data, err := service.AddQuiz(quiz)
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

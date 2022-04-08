package router

import (
	"encoding/json"
	"time"

	"net/http"
	quiz_service "personality_demo_backend/pkg/services/quiz"
	quiztype_service "personality_demo_backend/pkg/services/quiz_type"
	userassignedquiz_service "personality_demo_backend/pkg/services/user_assigned_quiz"
	userprogress_service "personality_demo_backend/pkg/services/user_progress"

	"github.com/aekam27/trestCommon"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// Route type description
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var routes = Routes{
	Route{
		"Add Quiz",
		"POST",
		"/quiz",
		quiz_service.AddQuiz,
	},
	Route{
		"Add Quiz Type",
		"POST",
		"/quiztype",
		quiztype_service.AddQuizType,
	},
	Route{
		"GET Quiz Type",
		"GET",
		"/quiztype",
		quiztype_service.GetQuizType,
	},
	Route{
		"AssignQuestion",
		"POST",
		"/assignuser",
		userassignedquiz_service.AssignUser,
	},
	Route{
		"Get Assigned Question",
		"GET",
		"/assignuser/{quizType}/{userID}",
		userassignedquiz_service.GetAssignedQuiestion,
	},
	Route{
		"Update Assigned Question",
		"PATCH",
		"/assignuser/{quizID}/{answer}/{quizNumber}",
		userassignedquiz_service.UpdateAssignedQuiestion,
	},
	Route{
		"Update User Progress",
		"POST",
		"/userprogress",
		userprogress_service.UpdateUserProgress,
	},
	Route{
		"Get Assigned Question",
		"GET",
		"/userprogress/{userID}",
		userprogress_service.GetUserProgress,
	},
	Route{
		"HEALTH",
		"GET",
		"/",
		Health,
	},
}

func Health(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	trestCommon.DLogMap("social media login", logrus.Fields{
		"start_time": startTime})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"status": true, "error": "", "health": "ok"})
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	trestCommon.DLogMap("login successfully", logrus.Fields{"duration": duration})
}

package user_assigned_db

import "personality_demo_backend/pkg/entity"

type UserAssignedService interface {
	AssignUser(*entity.UserAssignedEntity) (string, error)
	GetAssignedQuesion(quizType, userID string) (entity.UserAssignedEntity, error)
	UpdateQuizValue(id, answer string, quizNumber int) (string, error)
}

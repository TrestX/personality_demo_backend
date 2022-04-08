package user_progress_db

import "personality_demo_backend/pkg/entity"

type UserProgressService interface {
	UpdateUserProgress(*entity.UserProgressEntity) (string, error)
	GetProgress(userID string) (entity.UserProgressEntity, error)
}

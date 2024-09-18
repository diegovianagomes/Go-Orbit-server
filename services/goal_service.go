package services

import (
	"context"
	"go-orbit-server/models"
	"go-orbit-server/repositories"
	"time"

	"github.com/google/uuid"
)

type GoalService interface {
	CreateGoal(ctx context.Context, title string, desiredWeeklyFrequecy int) (models.Goal, error)
	GetGoals(ctx context.Context) ([]models.Goal, error)
}

type GoalServiceImpl struct {
	goalRepo repositories.GoalRepository
}

func NewGoalService(goalRepo repositories.GoalRepository) *GoalServiceImpl {
	return &GoalServiceImpl{goalRepo: goalRepo}
}

func (s *GoalServiceImpl) CreateGoal(ctx context.Context, title string, desiredWeeklyFrequency int) (models.Goal, error) {
	goal := models.Goal{
		ID:                     uuid.New().String(),
		Title:                  title,
		DesiredWeeklyFrequency: desiredWeeklyFrequency,

		CreatedAt: time.Now(),
	}

	err := s.goalRepo.CreateGoal(ctx, goal)
	return goal, err
}
func (s *GoalServiceImpl) GetGoals(ctx context.Context) ([]models.Goal, error) {
	return s.goalRepo.GetGoals(ctx)
}

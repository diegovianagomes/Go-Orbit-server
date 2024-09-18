package models

import (
	"time"
)

// Defining data structure for table Goal
type Goal struct {
	ID                     string
	Title                  string
	DesiredWeeklyFrequency int
	CreatedAt              time.Time
}

// Defining data structure for table Goal-Completion
type GoalCompletion struct {
	ID        string
	GoalID    string
	CreatedAt time.Time
}

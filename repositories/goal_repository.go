package repositories

import (
	"context"
	"go-orbit-server/models"

	"github.com/jackc/pgx/v4"
)

// Defining main methods Create Goal and Get Goals
type GoalRepository interface {
	CreateGoal(ctx context.Context, goal models.Goal) error
	GetGoals(ctx context.Context) ([]models.Goal, error)
}

// Implements interfaces interacting with PostgreSQL using pgx
type PostgreSQLGoalRepository struct {
	db *pgx.Conn
}

func NewGoalRepository(db *pgx.Conn) *PostgreSQLGoalRepository {
	return &PostgreSQLGoalRepository{db: db}
}

// Goal Insertion into the DB
func (r *PostgreSQLGoalRepository) CreatedGoal(ctx context.Context, goal models.Goal) error {
	query := `INSERT INTO goals (id, title, desired_weekly_frequency, created_at) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(
		ctx,
		query,
		goal.ID,
		goal.Title,
		goal.DesiredWeeklyFrequency,
		goal.CreatedAt,
	)
	return err
}

// Retrieving all goals from the DB
func (r *PostgreSQLGoalRepository) GetGoals(ctx context.Context) ([]models.Goal, error) {
	query := `SELECT id, title, desired_weekly_frequency, created_at FROM goals`
	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var goals []models.Goal
	for rows.Next() {
		var goal models.Goal
		if err := rows.Scan(
			&goal.ID,
			&goal.Title,
			&goal.DesiredWeeklyFrequency,
			&goal.CreatedAt,
		); err != nil {
			return nil, err
		}
		goals = append(goals, goal)
	}
	return goals, nil

}

// Retrieving goal by Id from the DB
func (r *PostgreSQLGoalRepository) GetGoalByID(ctx context.Context, id string) (models.Goal, error) {
	query := `SELECT id, title, desired_weekly_frequency, created_at FROM goals WHERE id=$1`
	var goal models.Goal
	err := r.db.QueryRow(ctx, query, id).Scan(
		&goal.ID,
		&goal.ID,
		&goal.DesiredWeeklyFrequency,
		&goal.CreatedAt,
	)

	if err != nil {
		return models.Goal{}, err
	}
	return goal, nil
}

// Updating goals details in DB
func (r *PostgreSQLGoalRepository) UpdateGoal(ctx context.Context, goal models.Goal) error {
	query := `UPDATE goals SET title=$1, desired_weekly_frequency=$2 WHERE id=$3`
	_, err := r.db.Exec(
		ctx,
		query,
		goal.Title,
		goal.DesiredWeeklyFrequency,
		goal.ID,
	)
	return err
}

// Deleting goals in the DB
func (r *PostgreSQLGoalRepository) DeleteGoal(ctx context.Context, id string) error {
	query := `DELETE FROM goals WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

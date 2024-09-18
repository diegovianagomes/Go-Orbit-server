package repositories

import (
	"context"
	"go-orbit-server/models"

	"github.com/jackc/pgx/v4"
)

// Creates a interface for the repository
type GoalCompletionRepository interface {
	CreateGoalCompletion(ctx context.Context, completion models.GoalCompletion) error
	GetCompletionByGoalID(ctx context.Context, goalID string) ([]models.GoalCompletion, error)
	DeleteGoalCompletion(ctx context.Context, id string) error
}

// Implments the interface
type PostgreSQLGoalCompletionRepository struct {
	db *pgx.Conn
}

// CreateGoal implements GoalRepository.
func (r *PostgreSQLGoalCompletionRepository) CreateGoal(ctx context.Context, goal models.Goal) error {
	return nil
}

// GetGoals implements GoalRepository.
func (r *PostgreSQLGoalCompletionRepository) GetGoals(ctx context.Context) ([]models.Goal, error) {
	return nil, nil
}

func NewGoalCompletionRepository(db *pgx.Conn) *PostgreSQLGoalCompletionRepository {
	return &PostgreSQLGoalCompletionRepository{db: db}
}

// Insertion of a new goal into the DB
func (r *PostgreSQLGoalCompletionRepository) CreateGoalCompletion(ctx context.Context, completion models.GoalCompletion) error {
	query := `INSERT INTO goal_completions (id, goal_id, created_at) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(
		ctx,
		query,
		completion.ID,
		completion.GoalID,
		completion.CreatedAt,
	)
	return err
}

// Retrives all completions for an specific goal
func (r *PostgreSQLGoalCompletionRepository) GetCompletionByGoalID(ctx context.Context, goalID string) ([]models.GoalCompletion, error) {
	query := `SELECT id, goal_id, created_at FROM goal_completions WHERE goal_id=$1`
	rows, err := r.db.Query(ctx, query, goalID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var completions []models.GoalCompletion
	for rows.Next() {
		var completion models.GoalCompletion
		if err := rows.Scan(&completion.ID, &completion.GoalID, &completion.CreatedAt); err != nil {
			return nil, err
		}
		completions = append(completions, completion)
	}
	return completions, nil
}

// Remove a goal from the DB
func (r *PostgreSQLGoalCompletionRepository) DeleteGoalCompletion(ctx context.Context, id string) error {
	query := `DELETE FROM goal_completions WHERE id=$1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

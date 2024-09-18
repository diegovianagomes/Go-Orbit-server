package routes

import (
	"go-orbit-server/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GoalHandler struct {
	goalService services.GoalService
}

func NewGoalHandler(GoalService services.GoalService) *GoalHandler {
	return &GoalHandler{goalService: GoalService}
}

// Handling the creation of a goal
func (h *GoalHandler) CreateGoalHandler(c echo.Context) error {
	type request struct {
		Title                  string
		DesiredWeeklyFrequency int
	}
	req := new(request)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	goal, err := h.goalService.CreateGoal(
		c.Request().Context(),
		req.Title,
		req.DesiredWeeklyFrequency,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, goal)
}

// Retrieves all goals
func (h *GoalHandler) GetGoalsHandler(c echo.Context) error {
	goals, err := h.goalService.GetGoals(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, goals)
}

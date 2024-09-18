package main

import (
	"go-orbit-server/db"
	"go-orbit-server/repositories"
	"go-orbit-server/routes"
	"go-orbit-server/services"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env File")
	}
	err = db.Connect()
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	defer db.Close()

	// Initiates echo
	e := echo.New()

	// Repo and Services
	goalRepo := repositories.NewGoalCompletionRepository(db.Conn)
	goalService := services.NewGoalService(goalRepo)
	goalHandler := routes.NewGoalHandler(goalService)

	//Routes
	e.POST("/goals", goalHandler.CreateGoalHandler)
	e.GET("/goals", goalHandler.GetGoalsHandler)

	//Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))

}

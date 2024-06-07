package main

import (
	"github.com/MagnusV9/tietoevry-pong-mmr/api/config"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/controllers"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/repos"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/services"
)

// @Title TIME-SYS-API
// @Version 0.1
func main() {
	config.Connect()

	// Initialize repos
	employeeRepo := repos.NewEmployeeRepo(config.DB)
	gameRepo := repos.NewGameRepo(config.DB)

	// Initialize services
	authService := services.NewAuthService(employeeRepo)
	employeeService := services.NewEmployeeService(employeeRepo, authService)
	gameService := services.NewGameService(gameRepo)

	// Initialize controllers
	authController := controllers.NewAuthController(authService)
	employeeController := controllers.NewEmployeeController(employeeService)
	gameController := controllers.NewGameController(gameService, employeeService)

	app := config.NewRouter(
		authController,
		employeeController,
		gameController,
	)

	app.Listen(":8080")
}

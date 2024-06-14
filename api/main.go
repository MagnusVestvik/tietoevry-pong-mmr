package main

import (
	"github.com/MagnusV9/tietoevry-pong-mmr/api/config"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/controllers"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/repos"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/services"
)

// @Title tietoevry-pong-mmr
// @Version 0.1
func main() {
	config.Connect()

	// Initialize repos
	employeeRepo := repos.NewEmployeeRepo(config.DB)
	gameRepo := repos.NewGameRepo(config.DB)
	tournamentRepo := repos.NewTournamentRepo(config.DB)

	// Initialize services
	authService := services.NewAuthService(employeeRepo)
	employeeService := services.NewEmployeeService(employeeRepo, authService)
	gameService := services.NewGameService(gameRepo, employeeRepo)
	tournamentService := services.NewTournamentService(tournamentRepo)

	// Initialize controllers
	authController := controllers.NewAuthController(authService)
	employeeController := controllers.NewEmployeeController(employeeService)
	gameController := controllers.NewGameController(gameService, employeeService)
	tournamentController := controllers.NewTournamentController(tournamentService)

	app := config.NewRouter(
		authController,
		employeeController,
		gameController,
		tournamentController,
	)

	app.Listen(":8080")
}

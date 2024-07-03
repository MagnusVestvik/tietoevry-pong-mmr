package services

import (
	weberrors "github.com/MagnusV9/tietoevry-pong-mmr/api/errors"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/models"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/repos"
	"github.com/google/uuid"
)

type GameService struct {
	gameRepo     *repos.GameRepo
	employeeRepo *repos.EmployeeRepo
}

func NewGameService(gameRepo *repos.GameRepo, employeeRepo *repos.EmployeeRepo) *GameService {
	return &GameService{
		gameRepo:     gameRepo,
		employeeRepo: employeeRepo,
	}
}

func (as *GameService) GetGames(jwt *models.JWT) (*[]models.Game, error) {
	return as.gameRepo.GetGames()
}

func (as *GameService) GetGame(jwt *models.JWT, id uuid.UUID) (*models.Game, error) {
	return as.gameRepo.GetGame(id)
}

func (as *GameService) GetGamesByEmployeeID(jwt *models.JWT, employeeID uuid.UUID) (*[]models.Game, error) {
	return as.gameRepo.GetGamesByEmployeeID(employeeID)
}

func (as *GameService) CreateGame(jwt *models.JWT, game *models.Game) error {
	e1, err := as.employeeRepo.GetEmployee(game.Employees[0].ID)
	if err != nil {
		return weberrors.NewError(400, "invalid employee")
	}
	game.Employees[0] = *e1

	e2, err := as.employeeRepo.GetEmployee(game.Employees[1].ID)
	if err != nil {
		return weberrors.NewError(500, "invalid employee")
	}
	game.Employees[1] = *e2

	return as.gameRepo.CreateGame(game)
}

func (as *GameService) DeleteGame(jwt *models.JWT, id uuid.UUID) error {
	return as.gameRepo.DeleteGame(id)
}

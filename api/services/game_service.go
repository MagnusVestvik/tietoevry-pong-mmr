package services

import (
	"github.com/MagnusV9/tietoevry-pong-mmr/api/models"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/repos"
	"github.com/google/uuid"
)

type GameService struct {
	gameRepo *repos.GameRepo
}

func NewGameService(gameRepo *repos.GameRepo) *GameService {
	return &GameService{gameRepo: gameRepo}
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
	return as.gameRepo.CreateGame(game)
}

func (as *GameService) DeleteGame(jwt *models.JWT, id uuid.UUID) error {
	return as.gameRepo.DeleteGame(id)
}

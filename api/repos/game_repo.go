package repos

import (
	weberrors "github.com/MagnusV9/tietoevry-pong-mmr/api/errors"
	weberrormapper "github.com/MagnusV9/tietoevry-pong-mmr/api/errors/mappers"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GameRepo struct {
	db *gorm.DB
}

func NewGameRepo(db *gorm.DB) *GameRepo {
	return &GameRepo{db: db}
}

func (ar *GameRepo) GetGames(bid uuid.UUID) (*[]models.Game, error) {
	var games []models.Game
	result := ar.db.Where("business_id = ?", bid).Find(&games)

	if err := weberrormapper.MapGormError("games", result.Error); err != nil {
		return nil, err
	}

	return &games, nil
}

func (ar *GameRepo) GetGame(bid uuid.UUID, id uuid.UUID) (*models.Game, error) {
	var game models.Game
	result := ar.db.Where("business_id = ? AND id = ?", bid, id).First(&game)

	if err := weberrormapper.MapGormError("game", result.Error); err != nil {
		return nil, err
	}

	return &game, nil
}

func (ar *GameRepo) GetGamesByEmployeeID(bid uuid.UUID, employeeID uuid.UUID) (*[]models.Game, error) {
	var games []models.Game
	result := ar.db.Where("business_id = ? AND employee_id = ?", bid, employeeID).Find(&games)
	if err := weberrormapper.MapGormError("games", result.Error); err != nil {
		return nil, err
	}

	return &games, nil
}

func (ar *GameRepo) CreateGame(bid uuid.UUID, game *models.Game) error {
	game.ID = uuid.New()
	result := ar.db.Create(game)
	if err := weberrormapper.MapGormError("game", result.Error); err != nil {
		return err
	}

	return nil
}

func (ar *GameRepo) DeleteGame(bid uuid.UUID, id uuid.UUID) error {
	result := ar.db.Where("business_id = ? AND id = ?", bid, id).Delete(&models.Game{})

	if err := weberrormapper.MapGormError("game", result.Error); err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return weberrors.NewError(404, "game not found")
	}

	return nil
}

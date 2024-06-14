package repos

import (
	weberrors "github.com/MagnusV9/tietoevry-pong-mmr/api/errors"
	weberrormapper "github.com/MagnusV9/tietoevry-pong-mmr/api/errors/mappers"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TournamentRepo struct {
	db *gorm.DB
}

func NewTournamentRepo(db *gorm.DB) *TournamentRepo {
	return &TournamentRepo{db: db}
}

func (ar *TournamentRepo) GetTournaments() (*[]models.Tournament, error) {
	var tournaments []models.Tournament
	result := ar.db.Find(&tournaments)

	if err := weberrormapper.MapGormError("tournaments", result.Error); err != nil {
		return nil, err
	}

	return &tournaments, nil
}

func (ar *TournamentRepo) GetTournament(id uuid.UUID) (*models.Tournament, error) {
	var tournament models.Tournament
	result := ar.db.Where("id = ?", id).First(&tournament)

	if err := weberrormapper.MapGormError("tournament", result.Error); err != nil {
		return nil, err
	}

	return &tournament, nil
}

func (ar *TournamentRepo) GetTournamentsByEmployeeID(employeeID uuid.UUID) (*[]models.Tournament, error) {
	var tournaments []models.Tournament
	result := ar.db.Where("employee_id = ?", employeeID).Find(&tournaments)
	if err := weberrormapper.MapGormError("tournaments", result.Error); err != nil {
		return nil, err
	}

	return &tournaments, nil
}

func (ar *TournamentRepo) CreateTournament(tournament *models.Tournament) error {
	tournament.ID = uuid.New()
	result := ar.db.Create(tournament)
	if err := weberrormapper.MapGormError("tournament", result.Error); err != nil {
		return err
	}

	return nil
}

func (ar *TournamentRepo) UpdateTournament(id uuid.UUID, tournament *models.Tournament) error {
	result := ar.db.Save(tournament)

	if err := weberrormapper.MapGormError("tournament", result.Error); err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return weberrors.NewError(404, "tournament not found")
	}

	return nil
}

func (ar *TournamentRepo) DeleteTournament(id uuid.UUID) error {
	result := ar.db.Where("id = ?", id).Delete(&models.Tournament{})

	if err := weberrormapper.MapGormError("tournament", result.Error); err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return weberrors.NewError(404, "tournament not found")
	}

	return nil
}

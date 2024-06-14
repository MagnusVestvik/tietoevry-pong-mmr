package services

import (
	"github.com/MagnusV9/tietoevry-pong-mmr/api/models"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/repos"
	"github.com/google/uuid"
)

type TournamentService struct {
	tournamentRepo *repos.TournamentRepo
}

func NewTournamentService(tournamentRepo *repos.TournamentRepo) *TournamentService {
	return &TournamentService{tournamentRepo: tournamentRepo}
}

func (as *TournamentService) GetTournaments(jwt *models.JWT) (*[]models.Tournament, error) {
	return as.tournamentRepo.GetTournaments()
}

func (as *TournamentService) GetTournament(jwt *models.JWT, id uuid.UUID) (*models.Tournament, error) {
	return as.tournamentRepo.GetTournament(id)
}

func (as *TournamentService) CreateTournament(jwt *models.JWT, tournament *models.Tournament) error {
	return as.tournamentRepo.CreateTournament(tournament)
}

func (as *TournamentService) UpdateTournament(jwt *models.JWT, id uuid.UUID, tournament *models.Tournament) error {
	return as.tournamentRepo.UpdateTournament(id, tournament)
}

func (as *TournamentService) DeleteTournament(jwt *models.JWT, id uuid.UUID) error {
	return as.tournamentRepo.DeleteTournament(id)
}

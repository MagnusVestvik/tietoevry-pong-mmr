package controllers

import (
	weberrors "github.com/MagnusV9/tietoevry-pong-mmr/api/errors"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/models"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TournamentController struct {
	tournamentService *services.TournamentService
	employeeService   *services.EmployeeService
}

func NewTournamentController(tournamentService *services.TournamentService, employeeService *services.EmployeeService) *TournamentController {
	return &TournamentController{
		tournamentService: tournamentService,
		employeeService:   employeeService,
	}
}

// @Summary Get all tournaments
// @Description Get all tournaments
// @Tags Tournament
// @Produce json
// @Failure 401
// @Failure 404
// @Failure 500
// @Success 200 {array} Tournament
// @Router /api/tournaments [get]
func (tc *TournamentController) GetTournaments(c *fiber.Ctx) error {
	jwt := mapReqToJWT(c)
	tournaments, err := tc.tournamentService.GetTournaments(jwt)
	if err != nil {
		return err
	}
	return c.JSON(tournaments)
}

// @Summary Get a tournament by ID
// @Description Get a tournament by its ID
// @Tags Tournament
// @Produce json
// @Param id path int true "Tournament ID"
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Success 200 {object} Tournament
// @Router /api/tournaments/{id} [get]
func (tc *TournamentController) GetTournament(c *fiber.Ctx) error {
	jwt := mapReqToJWT(c)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return weberrors.NewError(400, "invalid id")
	}

	tournament, err := tc.tournamentService.GetTournament(jwt, id)
	if err != nil {
		return err
	}

	return c.JSON(tournament)
}

// @Summary Create a new tournament
// @Description Create a new tournament
// @Tags Tournament
// @Produce json
// @Param tournament body Tournament true "Tournament object"
// @Failure 400
// @Failure 401
// @Failure 409
// @Failure 500
// @Success 201 {object} Tournament
// @Router /api/tournaments [post]
func (tc *TournamentController) CreateTournament(c *fiber.Ctx) error {
	var tournament models.Tournament
	jwt := mapReqToJWT(c)
	if err := c.BodyParser(&tournament); err != nil {
		return weberrors.NewError(400, "invalid request body")
	}

	if err := tc.tournamentService.CreateTournament(jwt, &tournament); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(tournament)
}

// @Summary Update a tournament
// @Description Update a tournament
// @Tags Tournament
// @Produce json
// @Param id path int true "Tournament ID"
// @Param tournament body Tournament true "Tournament object"
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Success 200 {object} Tournament
// @Router /api/tournaments/{id} [put]
func (tc *TournamentController) UpdateTournament(c *fiber.Ctx) error {
	jwt := mapReqToJWT(c)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return weberrors.NewError(400, "invalid id")
	}

	var tournament models.Tournament
	if err := c.BodyParser(&tournament); err != nil {
		return weberrors.NewError(400, "invalid request body")
	}

	if err := tc.tournamentService.UpdateTournament(jwt, id, &tournament); err != nil {
		return err
	}

	return c.JSON(tournament)
}

// @Summary Delete a tournament
// @Description Delete a tournament
// @Tags Tournament
// @Produce json
// @Param id path int true "Tournament ID"
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Success 204
// @Router /api/tournaments/{id} [delete]
func (tc *TournamentController) DeleteTournament(c *fiber.Ctx) error {
	jwt := mapReqToJWT(c)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return weberrors.NewError(400, "invalid id")
	}

	if err := tc.tournamentService.DeleteTournament(jwt, id); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	weberrors "github.com/MagnusV9/tietoevry-pong-mmr/api/errors"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/models"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/services"
)

type GameController struct {
	gameService *services.GameService
}

func NewGameController(gameService *services.GameService) *GameController {
	return &GameController{gameService: gameService}
}

// @Summary Get all games
// @Description Get all games
// @Tags Game
// @Produce json
// @Failure 401
// @Failure 404
// @Failure 500
// @Success 200 {array} Game
// @Router /api/games [get]
func (ac *GameController) GetGames(c *fiber.Ctx) error {
	jwt := mapReqToJWT(c)
	games, err := ac.gameService.GetGames(jwt)
	if err != nil {
		return err
	}
	return c.JSON(games)
}

// @Summary Get an game by ID
// @Description Get an game by its ID
// @Tags Game
// @Produce json
// @Param id path int true "Game ID"
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Success 200 {object} Game
// @Router /api/games/{id} [get]
func (ac *GameController) GetGame(c *fiber.Ctx) error {
	jwt := mapReqToJWT(c)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return weberrors.NewError(400, "invalid id")
	}

	game, err := ac.gameService.GetGame(jwt, id)
	if err != nil {
		return err
	}

	return c.JSON(game)
}

// @Summary Create a new game
// @Description Create a new game
// @Tags Game
// @Produce json
// @Param game body Game true "Game object"
// @Failure 400
// @Failure 401
// @Failure 409
// @Failure 500
// @Success 201 {object} Game
// @Router /api/games [post]
func (ac *GameController) CreateGame(c *fiber.Ctx) error {
	var game models.Game
	jwt := mapReqToJWT(c)
	if err := c.BodyParser(&game); err != nil {
		return weberrors.NewError(400, err.Error())
	}

	if err := ac.gameService.CreateGame(jwt, &game); err != nil {
		return err
	}

	return c.Status(201).JSON(game)
}

// @Summary Delete an game by ID
// @Description Delete an game by its ID
// @Tags Game
// @Produce json
// @Param id path int true "Game ID"
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Success 204 "No Content"
// @Router /api/games/{id} [delete]
func (ac *GameController) DeleteGame(c *fiber.Ctx) error {
	jwt := mapReqToJWT(c)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return weberrors.NewError(400, "invalid id")
	}

	if err := ac.gameService.DeleteGame(jwt, id); err != nil {
		return err
	}

	return c.SendStatus(204)
}

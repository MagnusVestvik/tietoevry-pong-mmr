package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	weberrors "github.com/MagnusV9/tietoevry-pong-mmr/api/errors"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/models"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/services"
)

// Shared auth controller-logic for mapping jwt claim to object
func mapReqToJWT(c *fiber.Ctx) *models.JWT {
	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)

	bidStr, ok := claims["bid"].(string)
	if !ok {
		return nil
	}

	role, ok := claims["role"].(string)
	if !ok {
		return nil
	}

	bid, err := uuid.Parse(bidStr)
	if err != nil {
		return nil
	}

	jwt := models.JWT{Bid: bid, Role: role}
	return &jwt
}

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Login handles the authentication process
// @Summary Login
// @Description Authenticates a user and returns a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param loginRequest body LoginRequest true "Login Request"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 401
// @Failure 500
// @Router /api/login [post]
func (ac *AuthController) Login(c *fiber.Ctx) error {
	var loginRequest models.LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		return weberrors.NewError(400, err.Error())
	}

	token, err := ac.authService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"token": token})
}

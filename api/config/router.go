package config

import (
	"errors"
	"os"

	"github.com/MagnusV9/tietoevry-pong-mmr/api/controllers"
	_ "github.com/MagnusV9/tietoevry-pong-mmr/api/docs"
	weberrors "github.com/MagnusV9/tietoevry-pong-mmr/api/errors"
	json "github.com/goccy/go-json"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func NewRouter(
	authController *controllers.AuthController,
	employeeController *controllers.EmployeeController,
	gameController *controllers.GameController,
	tournamentController *controllers.TournamentController,
) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,

		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			// Retrieve the custom status code if it's a *WebError
			var e2 *weberrors.WebError
			if errors.As(err, &e2) {
				code = e2.StatusCode
			}
			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

			return c.Status(code).JSON(fiber.Map{"error": err.Error()})
		},
	})
	app.Use(logger.New())
	app.Use(recover.New())
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	api := app.Group("/api")
	// Public Routes
	api.Post("/login", authController.Login)
	api.Post("/employees/", employeeController.CreateEmployee)

	// JWT Middleware
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid or expired jwt")
		},
	}))

	// Restricted Routes
	employees := api.Group("/employees")
	employees.Get("", employeeController.GetEmployees)
	// employees.Post("", employeeController.CreateEmployee)
	employees.Get("/:id", employeeController.GetEmployee)
	employees.Put("/:id", employeeController.UpdateEmployee)
	employees.Delete("/:id", employeeController.DeleteEmployee)

	games := api.Group("/games")
	games.Get("", gameController.GetGames)
	games.Post("", gameController.CreateGame)
	games.Get("/:id", gameController.GetGame)
	games.Delete("/:id", gameController.DeleteGame)

	tournaments := api.Group("/tournaments")
	tournaments.Get("", tournamentController.GetTournaments)
	tournaments.Post("", tournamentController.CreateTournament)
	tournaments.Get("/:id", tournamentController.GetTournament)
	tournaments.Delete("/:id", tournamentController.DeleteTournament)

	return app
}

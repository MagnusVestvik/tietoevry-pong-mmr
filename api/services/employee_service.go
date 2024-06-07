package services

import (
	"fmt"
	"regexp"

	weberrors "github.com/MagnusV9/tietoevry-pong-mmr/api/errors"
	"github.com/google/uuid"

	"github.com/MagnusV9/tietoevry-pong-mmr/api/models"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/repos"
)

type EmployeeService struct {
	employeeRepo *repos.EmployeeRepo
	authService  *AuthService
}

func NewEmployeeService(employeeRepo *repos.EmployeeRepo, authService *AuthService) *EmployeeService {
	return &EmployeeService{
		employeeRepo: employeeRepo,
		authService:  authService,
	}
}

func (es *EmployeeService) GetEmployees(jwt *models.JWT) (*[]models.Employee, error) {
	return es.employeeRepo.GetEmployees()
}

func (es *EmployeeService) GetEmployee(jwt *models.JWT, id uuid.UUID) (*models.Employee, error) {
	return es.employeeRepo.GetEmployee(id)
}

func (es *EmployeeService) CreateEmployee(employee *models.Employee) error {
	// Validate submitted required fields
	if err := es.validateEmployee(employee); err != nil {
		return weberrors.NewError(400, err.Error())
	}

	// Hash password and create employee
	hash, err := es.authService.HashPassword(employee.Password)
	if err != nil {
		return weberrors.NewError(500, "failed to hash password")
	}
	employee.Password = hash

	// Set starter elo
	employee.Elo = 1000

	return es.employeeRepo.CreateEmployee(employee)
}

func (es *EmployeeService) UpdateEmployee(jwt *models.JWT, id uuid.UUID, updatedEmployee *models.UpdateEmployee) error {
	existingEmployee, err := es.employeeRepo.GetEmployee(id)
	if err != nil {
		return err
	}

	if updatedEmployee.Name != nil {
		existingEmployee.Name = *updatedEmployee.Name
	}
	if updatedEmployee.Email != nil {
		existingEmployee.Email = *updatedEmployee.Email
	}
	if updatedEmployee.Password != nil {
		existingEmployee.Password = *updatedEmployee.Password
	}

	// Validate submitted required fields
	if err := es.validateEmployee(existingEmployee); err != nil {
		return weberrors.NewError(400, err.Error())
	}

	// Hash password if its submitted
	if updatedEmployee.Password != nil {
		hash, err := es.authService.HashPassword(*updatedEmployee.Password)
		if err != nil {
			return weberrors.NewError(500, "failed to hash password")
		}
		existingEmployee.Password = hash
	}

	return es.employeeRepo.UpdateEmployee(existingEmployee)
}

func (es *EmployeeService) UpdateEmployeeElo(jwt *models.JWT, game *models.Game) error {
	e1, _ := es.GetEmployee(jwt, game.Employee1ID)
	e2, _ := es.GetEmployee(jwt, game.Employee2ID)

	winner, loser := e1, e2
	if game.Employee2Score > game.Employee1Score {
		winner, loser = e2, e1
	}

	// calculate and update elo
	eloChange, _ := es.CalculateEloChange(winner.Elo, loser.Elo)
	winner.Elo += eloChange
	loser.Elo -= eloChange

	es.employeeRepo.UpdateEmployee(winner)
	es.employeeRepo.UpdateEmployee(loser)

	return nil
}

func (es *EmployeeService) DeleteEmployee(jwt *models.JWT, id uuid.UUID) error {
	return es.employeeRepo.DeleteEmployee(id)
}

//
// HELPER METHODS
//

func (es *EmployeeService) validateEmployee(employee *models.Employee) error {
	if employee == nil {
		return fmt.Errorf("employee is required")
	}

	if len(employee.Name) == 0 {
		return fmt.Errorf("name is required")
	}

	if err := validateEmail(&employee.Email); err != nil {
		return err
	}

	if len(employee.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	return nil
}

func validateEmail(email *string) error {
	if email == nil || *email == "" {
		return fmt.Errorf("email is required")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(*email) {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

func validatePhoneNumber(phone *string) error {
	if phone == nil || *phone == "" {
		return nil
	}

	phoneRegex := regexp.MustCompile(`^([+][0-9]{1,3}|[0-9]{3,5})?\s?[0-9]{7,10}$`)
	if !phoneRegex.MatchString(*phone) {
		return fmt.Errorf("invalid phone format")
	}
	return nil
}

func absDiffInt(x int, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

// ELO
// https://usatt.simplycompete.com/info/ratings
func (es *EmployeeService) CalculateEloChange(winnerElo int, loserElo int) (int, error) {
	eloDiff := absDiffInt(winnerElo, loserElo)
	var eloChange int
	var err error

	if winnerElo >= loserElo {
		eloChange, err = es.GetExpectedResult(eloDiff)
	} else {
		eloChange, err = es.GetUpsetResult(eloDiff)
	}

	if err != nil {
		return 0, err
	}

	return eloChange, nil
}

func (es *EmployeeService) GetUpsetResult(pointSpread int) (int, error) {
	elo := 0
	switch {
	case pointSpread >= 0 && pointSpread <= 12:
		elo = 8
	case pointSpread <= 37:
		elo = 10
	case pointSpread <= 62:
		elo = 13
	case pointSpread <= 87:
		elo = 16
	case pointSpread <= 112:
		elo = 20
	case pointSpread <= 137:
		elo = 25
	case pointSpread <= 162:
		elo = 30
	case pointSpread <= 187:
		elo = 35
	case pointSpread <= 212:
		elo = 40
	case pointSpread <= 237:
		elo = 45
	case pointSpread >= 238:
		elo = 50
	default:
		return 0, weberrors.NewError(400, "invalid point spread")
	}

	return elo, nil
}

func (es *EmployeeService) GetExpectedResult(pointSpread int) (int, error) {
	elo := 0
	switch {
	case pointSpread >= 0 && pointSpread <= 12:
		elo = 8
	case pointSpread <= 37:
		elo = 7
	case pointSpread <= 62:
		elo = 6
	case pointSpread <= 87:
		elo = 5
	case pointSpread <= 112:
		elo = 4
	case pointSpread <= 137:
		elo = 3
	case pointSpread <= 162:
		elo = 2
	case pointSpread <= 187:
		elo = 2
	case pointSpread <= 212:
		elo = 1
	case pointSpread <= 237:
		elo = 1
	case pointSpread >= 238:
		elo = 0
	default:
		return 0, weberrors.NewError(400, "invalid point spread")
	}

	return elo, nil
}

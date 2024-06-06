package services

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
	weberrors "github.com/MagnusV9/tietoevry-pong-mmr/api/errors"

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
	return es.employeeRepo.GetEmployees(jwt.Bid)
}

func (es *EmployeeService) GetEmployee(jwt *models.JWT, id uuid.UUID) (*models.Employee, error) {
	return es.employeeRepo.GetEmployee(jwt.Bid, id)
}

func (es *EmployeeService) GetEmployeesByBusinessLocation(jwt *models.JWT, location string) (*[]models.Employee, error) {
	return es.employeeRepo.GetEmployeesByBusinessLocation(jwt.Bid, location)
}

func (es *EmployeeService) CreateEmployee(jwt *models.JWT, employee *models.Employee) error {
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

	return es.employeeRepo.CreateEmployee(jwt.Bid, employee)
}

func (es *EmployeeService) UpdateEmployee(jwt *models.JWT, id uuid.UUID, updatedEmployee *models.UpdateEmployee) error {
	existingEmployee, err := es.employeeRepo.GetEmployee(jwt.Bid, id)
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

	return es.employeeRepo.UpdateEmployee(jwt.Bid, existingEmployee)
}

func (es *EmployeeService) DeleteEmployee(jwt *models.JWT, id uuid.UUID) error {
	return es.employeeRepo.DeleteEmployee(jwt.Bid, id)
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

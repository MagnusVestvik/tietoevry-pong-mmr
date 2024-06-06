package repos

import (
	weberrors "github.com/MagnusV9/tietoevry-pong-mmr/api/errors"
	weberrormapper "github.com/MagnusV9/tietoevry-pong-mmr/api/errors/mappers"
	"github.com/google/uuid"

	"github.com/MagnusV9/tietoevry-pong-mmr/api/models"
	"gorm.io/gorm"
)

type EmployeeRepo struct {
	db *gorm.DB
}

func NewEmployeeRepo(db *gorm.DB) *EmployeeRepo {
	return &EmployeeRepo{db: db}
}

func (er *EmployeeRepo) GetEmployees() (*[]models.Employee, error) {
	var employees []models.Employee
	result := er.db.Find(&employees)

	if err := weberrormapper.MapGormError("employees", result.Error); err != nil {
		return nil, err
	}

	return &employees, nil
}

func (er *EmployeeRepo) GetEmployee(id uuid.UUID) (*models.Employee, error) {
	var employee models.Employee
	result := er.db.Where("id = ?", id).First(&employee)

	if err := weberrormapper.MapGormError("employee", result.Error); err != nil {
		return nil, err
	}

	return &employee, nil
}

func (er *EmployeeRepo) CreateEmployee(employee *models.Employee) error {
	employee.ID = uuid.New()
	result := er.db.Create(employee)

	if err := weberrormapper.MapGormError("employee", result.Error); err != nil {
		return err
	}

	return nil
}

func (er *EmployeeRepo) UpdateEmployee(employee *models.Employee) error {
	result := er.db.Save(employee)

	if err := weberrormapper.MapGormError("employee", result.Error); err != nil {
		return err
	}

	return nil
}

func (er *EmployeeRepo) DeleteEmployee(id uuid.UUID) error {
	result := er.db.Where("id = ?", id).Delete(&models.Employee{})

	if err := weberrormapper.MapGormError("employee", result.Error); err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		return weberrors.NewError(404, "employee not found")
	}

	return nil
}

// AUTH
func (er *EmployeeRepo) GetEmployeeByEmail(email string) (*models.Employee, error) {
	var employee models.Employee
	result := er.db.Where("email = ?", email).First(&employee)

	if err := weberrormapper.MapGormError("employee", result.Error); err != nil {
		return nil, err
	}

	return &employee, nil
}

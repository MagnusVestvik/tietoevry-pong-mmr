package repos

import (
	"github.com/google/uuid"
	weberrors "github.com/MagnusV9/tietoevry-pong-mmr/api/errors"
	weberrormapper "github.com/MagnusV9/tietoevry-pong-mmr/api/errors/mappers"

	"github.com/MagnusV9/tietoevry-pong-mmr/api/models"
	"gorm.io/gorm"
)

type EmployeeRepo struct {
	db *gorm.DB
}

func NewEmployeeRepo(db *gorm.DB) *EmployeeRepo {
	return &EmployeeRepo{db: db}
}

func (er *EmployeeRepo) GetEmployees(bid uuid.UUID) (*[]models.Employee, error) {
	var employees []models.Employee
	result := er.db.Where("business_id = ?", bid).Find(&employees)

	if err := weberrormapper.MapGormError("employees", result.Error); err != nil {
		return nil, err
	}

	return &employees, nil
}

func (er *EmployeeRepo) GetEmployee(bid uuid.UUID, id uuid.UUID) (*models.Employee, error) {
	var employee models.Employee
	result := er.db.Where("business_id = ? AND id = ?", bid, id).First(&employee)

	if err := weberrormapper.MapGormError("employee", result.Error); err != nil {
		return nil, err
	}

	return &employee, nil
}

func (er *EmployeeRepo) GetEmployeesByBusinessLocation(bid uuid.UUID, location string) (*[]models.Employee, error) {
	var employees []models.Employee
	result := er.db.
		Joins("JOIN businesses ON employees.business_id = businesses.id").
		Where("businesses.location = ?", location).
		Find(&employees)

	if err := weberrormapper.MapGormError("employees", result.Error); err != nil {
		return nil, err
	}

	return &employees, nil
}

func (er *EmployeeRepo) CreateEmployee(bid uuid.UUID, employee *models.Employee) error {
	employee.ID = uuid.New()
	result := er.db.Create(employee)

	if err := weberrormapper.MapGormError("employee", result.Error); err != nil {
		return err
	}

	return nil
}

func (er *EmployeeRepo) UpdateEmployee(bid uuid.UUID, employee *models.Employee) error {
	result := er.db.Save(employee)

	if err := weberrormapper.MapGormError("employee", result.Error); err != nil {
		return err
	}

	return nil
}

func (er *EmployeeRepo) DeleteEmployee(bid uuid.UUID, id uuid.UUID) error {
	result := er.db.Where("business_id = ? AND id = ?", bid, id).Delete(&models.Employee{})

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

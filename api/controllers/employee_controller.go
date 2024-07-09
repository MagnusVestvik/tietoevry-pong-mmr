package controllers

import (
	weberrors "github.com/MagnusV9/tietoevry-pong-mmr/api/errors"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/models"
	"github.com/MagnusV9/tietoevry-pong-mmr/api/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type EmployeeController struct {
	employeeService *services.EmployeeService
}

func NewEmployeeController(employeeService *services.EmployeeService) *EmployeeController {
	return &EmployeeController{employeeService: employeeService}
}

// @Summary Get all employees
// @Description Get all employees
// @Tags Employee
// @Produce json
// @Failure 401
// @Failure 404
// @Failure 500
// @Success 200 {array} Employee
// @Router /api/employees [get]
func (ec *EmployeeController) GetEmployees(c *fiber.Ctx) error {
	jwt := mapReqToJWT(c)
	employees, err := ec.employeeService.GetEmployees(jwt)
	if err != nil {
		return err
	}
	return c.JSON(employees)
}

// @Summary Get an employee by ID
// @Description Get an employee by its ID
// @Tags Employee
// @Produce json
// @Param id path int true "Employee ID"
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Success 200 {object} Employee
// @Router /api/employees/{id} [get]
func (ec *EmployeeController) GetEmployee(c *fiber.Ctx) error {
	jwt := mapReqToJWT(c)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return weberrors.NewError(400, "invalid id")
	}

	employee, err := ec.employeeService.GetEmployee(jwt, id)
	if err != nil {
		return err
	}

	return c.JSON(employee)
}

// @Summary Create a new employee
// @Description Create a new employee
// @Tags Employee
// @Produce json
// @Param employee body Employee true "Employee object"
// @Failure 400
// @Failure 401
// @Failure 409
// @Failure 500
// @Success 201 {object} Employee
// @Router /api/employees [post]
func (ec *EmployeeController) CreateEmployee(c *fiber.Ctx) error {
	var employee models.Employee
	if err := c.BodyParser(&employee); err != nil {
		return weberrors.NewError(400, err.Error())
	}

	if err := ec.employeeService.CreateEmployee(&employee); err != nil {
		return err
	}

	return c.Status(201).JSON(employee)
}

// @Summary Update an existing employee
// @Description Update an existing employee by its ID
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body UpdateEmployee true "UpdateEmployee object"
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 409
// @Failure 500
// @Success 204 "No Content"
// @Router /api/employees/{id} [put]
func (ec *EmployeeController) UpdateEmployee(c *fiber.Ctx) error {
	jwt := mapReqToJWT(c)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return weberrors.NewError(400, "invalid id")
	}

	var employee models.UpdateEmployee
	if err := c.BodyParser(&employee); err != nil {
		return weberrors.NewError(400, err.Error())
	}

	if err := ec.employeeService.UpdateEmployee(jwt, id, &employee); err != nil {
		return err
	}

	return c.SendStatus(204)
}

// @Summary Delete an employee by ID
// @Description Delete an employee by its ID
// @Tags Employee
// @Produce json
// @Param id path int true "Employee ID"
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Success 204 "No Content"
// @Router /api/employees/{id} [delete]
func (ec *EmployeeController) DeleteEmployee(c *fiber.Ctx) error {
	jwt := mapReqToJWT(c)
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return weberrors.NewError(400, "invalid id")
	}

	if err := ec.employeeService.DeleteEmployee(jwt, id); err != nil {
		return err
	}

	return c.SendStatus(204)
}
